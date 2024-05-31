# Mysql Database for OAI-UDR

This mysql database is preconfigured with some user subscription information. 

## Usage

### Fetch the package

`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] database`

Details: https://kpt.dev/reference/cli/pkg/get/

### View package content

`kpt pkg tree database`

Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package

```
kpt live init database
kpt live apply database --reconcile-timeout=2m --output=table
```

Details: https://kpt.dev/reference/cli/live/
