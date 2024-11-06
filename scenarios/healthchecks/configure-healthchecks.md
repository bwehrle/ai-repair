The healtchecks are, by default, configured as below.  To change them, modify `config.yaml`.
  
```yaml
healtchecks:
  liveness:
    url: /alternate-livez-path
  readiness:
    url: /alternate-readyz-path
```
