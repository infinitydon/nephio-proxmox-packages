# oai-up-operators

The package contains operator for OAI core network user plane network function

- UPF

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/oai-up-operators@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-up-operators`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package

The default namespace is oai-up-operators

```
kpt live init oai-up-operators
kpt live apply oai-up-operators --reconcile-timeout=2m --output=table
```

Details: https://kpt.dev/reference/cli/live/
