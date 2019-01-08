package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/autonomy/talos/internal/app/init/internal/platform"
	"github.com/autonomy/talos/internal/app/init/internal/rootfs"
	"github.com/autonomy/talos/internal/app/init/internal/rootfs/mount"
	"github.com/autonomy/talos/internal/app/init/pkg/system"
	"github.com/autonomy/talos/internal/app/init/pkg/system/services"
	"github.com/autonomy/talos/internal/pkg/constants"
	"github.com/autonomy/talos/internal/pkg/userdata"
	"github.com/pkg/errors"

	"golang.org/x/sys/unix"
)

var (
	switchRoot *bool
)

func init() {
	switchRoot = flag.Bool("switch-root", false, "perform a switch_root")
	flag.Parse()
}

func kmsg(prefix string) (*os.File, error) {
	out, err := os.OpenFile("/dev/kmsg", os.O_RDWR|unix.O_CLOEXEC|unix.O_NONBLOCK|unix.O_NOCTTY, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open /dev/kmsg")
	}
	log.SetOutput(out)
	log.SetPrefix(prefix + " ")
	log.SetFlags(0)

	return out, nil
}

// nolint: gocyclo
func initram() (err error) {
	var initializer *mount.Initializer
	if initializer, err = mount.NewInitializer(constants.NewRoot); err != nil {
		return err
	}
	// Mount the special devices.
	if err = initializer.InitSpecial(); err != nil {
		return err
	}
	// Setup logging to /dev/kmsg.
	_, err = kmsg("[talos] [initramfs]")
	if err != nil {
		return err
	}
	// Discover the platform.
	log.Println("discovering the platform")
	var p platform.Platform
	if p, err = platform.NewPlatform(); err != nil {
		return err
	}
	log.Printf("platform is: %s", p.Name())
	// Retrieve the user data.
	log.Printf("retrieving the user data")
	var data userdata.UserData
	if data, err = p.UserData(); err != nil {
		return err
	}
	// Perform rootfs/datafs installation if needed.
	if err = p.Install(data); err != nil {
		return err
	}
	// Mount the owned partitions.
	log.Printf("mounting the partitions")
	if err = initializer.InitOwned(); err != nil {
		return err
	}
	// Perform any tasks required by a particular platform.
	log.Printf("performing platform specific tasks")
	if err = p.Prepare(data); err != nil {
		return err
	}
	// Prepare the necessary files in the rootfs.
	log.Println("preparing the root filesystem")
	if err = rootfs.Prepare(constants.NewRoot, data); err != nil {
		return err
	}
	// Perform the equivalent of switch_root.
	log.Println("entering the new root")
	if err = initializer.Switch(); err != nil {
		return err
	}

	return nil
}

func root() error {
	// Setup logging to /dev/kmsg.
	if _, err := kmsg("[talos]"); err != nil {
		return fmt.Errorf("failed to setup logging to /dev/kmsg: %v", err)
	}
	// Read the user data.
	log.Printf("reading the user data: %s\n", constants.UserDataPath)
	data, err := userdata.Open(constants.UserDataPath)
	if err != nil {
		return err
	}

	// Write any user specified files to disk.
	log.Println("writing the files specified in the user data to disk")
	if err := data.WriteFiles(); err != nil {
		return err
	}

	// Set the requested environment variables.
	log.Println("setting environment variables")
	for key, val := range data.Env {
		if err := os.Setenv(key, val); err != nil {
			log.Printf("WARNING failed to set enivronment variable: %v", err)
		}
	}

	// Get a handle to the system services API.
	systemservices := system.Services(data)

	// Start the services common to all nodes.
	log.Println("starting node services")
	systemservices.Start(
		&services.Containerd{},
		&services.Udevd{},
		&services.OSD{},
		&services.Blockd{},
		&services.Kubelet{},
		&services.Kubeadm{},
	)

	// Start the services common to all master nodes.
	if data.IsMaster() {
		log.Println("starting master services")
		systemservices.Start(
			&services.Trustd{},
			&services.Proxyd{},
		)
	}

	return nil
}

func recovery() {
	if r := recover(); r != nil {
		log.Printf("recovered from: %+v\n", r)
		for i := 10; i >= 0; i-- {
			log.Printf("rebooting in %d seconds\n", i)
			time.Sleep(1 * time.Second)
		}

		// nolint: errcheck
		unix.Reboot(int(unix.LINUX_REBOOT_CMD_RESTART))
	}

	select {}
}

func main() {
	defer recovery()

	// TODO(andrewrynhard): Remove this and be explicit.
	if err := os.Setenv("PATH", constants.PATH); err != nil {
		panic(errors.New("error setting PATH"))
	}

	if *switchRoot {
		if err := root(); err != nil {
			panic(errors.Wrap(err, "boot failed"))
		}

		// Hang forever.
		select {}
	}

	if err := initram(); err != nil {
		panic(errors.Wrap(err, "early boot failed"))
	}

	// We should never reach this point if things are working as intended.
	panic(errors.New("unkown error"))
}