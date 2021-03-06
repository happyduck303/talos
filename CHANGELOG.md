
<a name="v0.7.0-alpha.8"></a>
## [v0.7.0-alpha.8](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.7...v0.7.0-alpha.8) (2020-10-29)

### Chore

* output more logs from the installer
* update CI scripts
* move to newer release of rtnetlink with fn args
* reduce numer of steps/parallelism of Drone build
* fix the check-dirty command to abort on untracked files

### Docs

* use grid instead of flexbox
* improve the config reference documentation
* improve search bar
* address small nits
* add robots.txt and fix sitemap.xml
* fix config reference types links
* move to gridsome

### Feat

* bump CoreDNS to 1.7.0
* bump Linux to 5.8.16, enable mpt3sas driver
* bump CoreDNS to 1.7.0
* encode comments as part of talosctl generated configs
* extend etcd health check on upgrade
* wipe disks faster in the installer

### Fix

* bump type for `DiskSize` to be 64-bit
* properly initialize manifest in user disks creation
* remove default time server in time command
* retry connection refused errors while bootstrapping a cluster
* re-implement upgrade (install) with preserve
* revert "feat: bump CoreDNS to 1.7.0"
* stop CRI pods on upgrade with preserve
* stop etcd on any path on upgrade
* ignore transient errors in upgrade Kubernetes code
* stop ignoring `EINVAL` on mount
* implement preserving contents of partition on install
* correctly calculate output width in colored health reporter

### Test

* bump Talos version for upgrade tests, bump Cilium version


<a name="v0.7.0-alpha.7"></a>
## [v0.7.0-alpha.7](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.6...v0.7.0-alpha.7) (2020-10-20)

### Chore

* bump module dependencies in go.mod
* bump Go to 1.15.3

### Docs

* link container images to our repository
* fix latest tag
* add link to latest docs

### Feat

* upgrade Kubernetes to 1.19.3
* support MTU and route changes for DHCP
* bump packages for Linux 5.8.15 and containerd 1.4.1
* support metric values for DHCP
* bump packages version for the kernel with BBR TCP congestion algo
* handle unsupported commands being called for docker
* support disk usage command in talosctl

### Fix

* update handling of ntp disable
* address nil pointer panic
* improve logging and errors for extra manifests by URL

### Refactor

* bring more control to install.Manifest execution

### Release

* **v0.7.0-alpha.7:** prepare release

### Test

* clean up integration test code, fix flakes
* add unit-test for the installer manifest


<a name="v0.7.0-alpha.6"></a>
## [v0.7.0-alpha.6](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.5...v0.7.0-alpha.6) (2020-10-09)

### Release

* **v0.7.0-alpha.6:** prepare release

### Test

* potential fix for talosctl cluster destroy being stuck


<a name="v0.7.0-alpha.5"></a>
## [v0.7.0-alpha.5](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.4...v0.7.0-alpha.5) (2020-10-08)

### Chore

* add Context as param to some methods of `Platform` interface
* bump pkgs version
* publish list of images to release notes

### Feat

* bring in install-cni & pod-checkpointer from extras packages
* implement talos.shutdown=[halt|poweroff] kernel argument

### Fix

* random failures in cluster health checks
* apply --removable option always to get standard UEFI filename
* nil pointer panic in talosctl dashboard

### Release

* **v0.7.0-alpha.5:** prepare release


<a name="v0.7.0-alpha.4"></a>
## [v0.7.0-alpha.4](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.3...v0.7.0-alpha.4) (2020-10-06)

### Chore

* attempt to fix image pushing for GitHub
* update qemu hack script to use ISO
* fix 'push' targets

### Docs

* small fixes for the config docs and air-gapped
* add guide on setting up air-gapped environment with `images`
* add note on settings endpoints on MacOS
* remove second meeting from README
* fix cluster name in docker docs

### Feat

* add etcd API
* allow disabling NoSchedule on master nodes
* colorize output of cluster health checks
* pull kubeconfig from the cluster on successful `cluster create`
* use kubeconfig merge in `talosctl kubeconfig` by default
* support --registry-insecure-skip-verify for `cluster create`
* show cluster state when `talosctl cluster create` finishes
* support custom filename for talosctl kubeconfig
* add support for disabling time
* add ApplyConfiguration API
* validate cluster DNS name
* build Talos images/artifacts for amd64/arm64

### Fix

* make CLI context exit immediately on second ^C
* registry auth config building
* provide unique username in generate kubeconfig
* make Flannel CNI image follow `$PKGS` version
* retry container image import

### Refactor

* extract blockdevice library

### Release

* **v0.7.0-alpha.4:** prepare release


<a name="v0.7.0-alpha.3"></a>
## [v0.7.0-alpha.3](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.2...v0.7.0-alpha.3) (2020-09-25)

### Chore

* fix edge push
* fix docker login
* fix docker login
* migrate to ghcr.io
* push edge releases on successful nightly integration

### Docs

* add note around link-local addressing
* add ghcr.io to the registry cache docs
* add v0.7 docs

### Feat

* bump default resource limits for `talosctl cluster create`
* add default install image
* add images command

### Fix

* update one more places which had stale reference for constants
* update the docs to fix the lint-markdown
* use images package in integration tests
* move installer image variables out of machinery
* enable --removable options for GRUB
* retry image pulling, stop on 404, no duplicate pulls

### Release

* **v0.7.0-alpha.3:** prepare release


<a name="v0.7.0-alpha.2"></a>
## [v0.7.0-alpha.2](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.1...v0.7.0-alpha.2) (2020-09-18)

### Chore

* update ntp time headers
* upgrade Go to 1.15.1
* remove extra COPY from rootfs

### Docs

* add recommneded settings in overview
* update upgrade guide with `talosctl upgrade-k8s`
* update 0.6 links

### Feat

* ugrade Linux kernel to 5.8.10
* allow for link local networking
* use architecture-specific image for core k8s components
* update Flannel to 0.12, support for arm64
* upgrade kubernetes to 1.19.1
* implement command `talosctl upgrade-k8s`
* use latest packages
* upgrade runc to v1.0.0-rc92
* upgrade containerd to v1.4.0
* remove ISO support

### Fix

* address node package update
* validate cluster endpoint
* improve error message on empty config
* gracefully handle invalid interfaces in bond
* set environment variable for etcd on arm64
* don't enforce k8s version in `talosctl cluster create` by default
* tell grub to use console output
* update vmware image and platform
* don't abort reboot sequence on bootloader meta failure
* default endpoint to 127.0.0.1 for Docker/OS X
* remove udevd debug flag
* update permissions for directories and files created via extraFiles
* allow static pod files

### Refactor

* garbage collect unused constants
* deduplicate packages version in Dockerfile

### Release

* **v0.7.0-alpha.2:** prepare release

### Test

* implement API for QEMU VM provisioner
* re-enable Cilium e2e upgrade test
* verify kubernetes control plane upgrade in provision tests
* add e2e test to the provision (upgrade) tests


<a name="v0.7.0-alpha.1"></a>
## [v0.7.0-alpha.1](https://github.com/talos-systems/talos/compare/v0.7.0-alpha.0...v0.7.0-alpha.1) (2020-09-02)

### Chore

* update k8s modules to 1.19 final version
* upgrade Go to 1.14.8
* drop vmlinux from assets
* add a method to merge Talos client config
* bump next version to v0.6.0-beta.2
* update machinery version in go.mod
* update node.js dependencies

### Docs

* graduate v0.6 docs
* add Kubernetes upgrade guide
* add reset doc
* add QEMU provisioner documentation

### Feat

* add grub bootloader
* upgrade etcd to 3.4.12
* provide option to run Talos under UEFI in QEMU
* update linux to 5.8.5
* update kubernetes to v1.19.0
* make boostrap via API default choice in talosctl cluster create
* upgrade Linux to v5.7.15

### Fix

* change apid container image name to expected value
* add syslinux to create ISO
* pass config via stdin
* handle bootkube recover correctly, support recovery from etcd

### Refactor

* move udevadm trigger/settle to udevd healthcheck
* extract packages loadbalancer and retry
* extract cluster bootstrapper via API as common component

### Release

* **v0.7.0-alpha.1:** prepare release
* **v0.7.0-alpha.1:** prepare release

### Test

* determine reboots using boot id
* add support for PXE nodes in qemu provision library

### BREAKING CHANGE


Single node upgrades will fail in this change. This
will also break the A/B fallback setup since this version introduces
an entirely new partition scheme, that any fallback will not know about.
We plan on addressing these issues in a follow up change.

