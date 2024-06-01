# oai-ranc-operators

The package contains oai ranc operators

- CU-CP

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/oai-ranc-operators@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-ranc-operators`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package

The default namespace is oai-ranc-operators

```
kpt live init oai-ranc-operators
kpt live apply oai-ranc-operators --reconcile-timeout=2m --output=table
```

Details: https://kpt.dev/reference/cli/live/
