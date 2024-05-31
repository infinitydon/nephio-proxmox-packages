# oai-nr-ue

KPT package for oai-nr-ue

## Usage

### Fetch the package

`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] oai-nr-ue`

Details: https://kpt.dev/reference/cli/pkg/get/

### View package content

`kpt pkg tree oai-nr-ue`

Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package

```
kpt fn render oai-nr-ue
kpt live init oai-nr-ue
kpt live apply oai-nr-ue --reconcile-timeout=2m --output=table 
```

In case you want to install it on a particular cluster then use `--context` flag

Details: https://kpt.dev/reference/cli/live/
