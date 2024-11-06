Alerts should be created for all services and added to the values.yaml file.  The minimum alert is called `crash-loop`.

```yaml
 alerts:
  - name: CrashLoop
    annotations:
        description: 'Pod {{ $labels.namespace }}/{{ $labels.pod }} ({{ $labels.container
          }}) is in waiting state (reason: "CrashLoopBackOff").'
        runbook_url: https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubepodcrashlooping
        summary: Pod is crash looping.
    expr: max_over_time(kube_pod_container_status_waiting_reason{reason="CrashLoopBackOff",
        job="kube-state-metrics", namespace=~".*"}[5m]) >= 1
    for: 15m
    labels:
      severity: warning
```