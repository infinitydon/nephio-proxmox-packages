# oai-amf package

## Description

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/nf-deploy-packages/oai-amf@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-amf`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init oai-amf
kpt live apply oai-amf --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
