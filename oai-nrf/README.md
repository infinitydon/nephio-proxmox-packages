# oai-nrf package

## Description

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/nf-deploy-packages/oai-nrf@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-nrf`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init oai-nrf
kpt live apply oai-nrf --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
