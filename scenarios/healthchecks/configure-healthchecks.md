The readiness and liveness health check urls are, by default, configured as below.  To change them, modify `config.yaml`.
  
```yaml
healthchecks:
  liveness:
    url: /alternate-livez-path
  readiness:
    url: /alternate-readyz-path
```
