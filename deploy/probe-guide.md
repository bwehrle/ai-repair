Readiness and liveness probes must be added to ensure that the appropriate path is used.
```yaml
readiness:
   enabled: true
   path: /status/200
   port: 80
   initialDelaySeconds: 5
   periodSeconds: 5
   failureThreshold: 3
liveness:
  enabled: true
  path: /status/200
  initialDelaySeconds: 5
  periodSeconds: 5
  failureThreshold: 5
```