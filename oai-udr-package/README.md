# oai-udr package

## Description

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/nf-deploy-packages/oai-udr@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-udr`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init oai-udr
kpt live apply oai-udr --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
