# oai-cp-operators

The package contains only control plane network function operator for OAI 5G Core. 

- AMF
- SMF
- NRF
- SMF
- UDR
- UDM
- AUSF

## Usage

### Fetch the package
`kpt pkg get https://gitlab.eurecom.fr/nephio/oai-packages/oai-cp-operators@main`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree oai-cp-operators`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package

The default namespace is oai-cp-operators

```
kpt live init oai-cp-operators
kpt live apply oai-cp-operators --reconcile-timeout=2m --output=table
```

Details: https://kpt.dev/reference/cli/live/
