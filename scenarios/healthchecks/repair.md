Readiness and liveness probes are known as `health checks` are defined in the `values.yaml` file.  In the example, placeholder values are provided and platform-specific values must be used.

```yaml
readiness:
   enabled: true
   path: READINESS-URL
   port: 80
   initialDelaySeconds: 5
   periodSeconds: 5
   failureThreshold: 3
liveness:
  enabled: true
  path: LIVENESS-URL
  initialDelaySeconds: 5
  periodSeconds: 5
  failureThreshold: 5
```

## References
* [Configuring readiness and liveness urls](configure-healthchecks.md)
* [URL validation](validate-url.md)