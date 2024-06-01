# oai-ranu-operators

The package contains oai ranu operators

- CU-UP
- DU

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/oai-ranu-operators@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-ranu-operators`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package

The default namespace is oai-ranu-operators

```
kpt live init oai-ranu-operators
kpt live apply oai-ranu-operators --reconcile-timeout=2m --output=table
```

Details: https://kpt.dev/reference/cli/live/
