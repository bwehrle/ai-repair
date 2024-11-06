Readiness and liveness probes are known as `health checks` must be added to ensure that the appropriate path is used.  In the example, placeholder values are provided and these are not correct.


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

The value of the `READINESS` and `LIVENESS-URL` are specific to the platform.

## References
* [Configuring readiness and liveness urls](configure-healthcheck)
* [URL validation](validate-url.md)