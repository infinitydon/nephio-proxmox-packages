# oai-ausf package

## Description

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/nf-deploy-packages/oai-ausf@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-ausf`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init oai-ausf
kpt live apply oai-ausf --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
