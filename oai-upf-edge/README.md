# oai-upf package

## Description

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/nf-deploy-packages/oai-upf@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-upf`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init oai-upf-edge
kpt live apply oai-upf-edge --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
