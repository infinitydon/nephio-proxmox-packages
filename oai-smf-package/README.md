# oai-smf package

## Description

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/nf-deploy-packages/oai-smf@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-smf`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init oai-smf
kpt live apply oai-smf --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
