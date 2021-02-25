# To do

- Read PromQL
- Relabel measurements to app@namespace
- Add to Grafana dashboard:
  - `sum(increase(deployed_total[1y]))`?
  - `sum_over_time(deployed_total[1d])`?

## Check

- https://www.weave.works/blog/how-many-kubernetes-replicasets-are-in-your-cluster-
- https://prometheus.io/docs/prometheus/latest/querying/functions/
- https://stackoverflow.com/questions/47138461/get-total-requests-in-a-period-of-time