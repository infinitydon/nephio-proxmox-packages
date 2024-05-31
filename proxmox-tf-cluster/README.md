# proxmox-tf-cluster

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] proxmox-tf-cluster`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree proxmox-tf-cluster`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init proxmox-tf-cluster
kpt live apply proxmox-tf-cluster --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
