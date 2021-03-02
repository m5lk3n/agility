# agility

This is a cloud-native implementation of how to measure agility (read: digital transformation) using Kubernetes *Deployment Frequency* (DF) (or *Throughput*).

## Specification

### DF KPI

- Purely based upon K8s deployments (directly via `kubectl deploy`, or indirectly via Helm or CD tooling like ArgoCD)
- Per cluster
- Per namespace (*support include/exclude pattern!*)
- Per app (name) (*support include/exclude pattern!*)
- Per aggregate (*support name pattern!*)

### Tech

- Deploy and run on Kubernetes only
- Store settings in configmap (offer restart/reload option later)

## Implementation

### Metric idea

1. `deployed{app=<name>,namespace=<name>}=<UnixTimeOfDeployment>`:

- same as `kube_deployment_created`
- doesn't capture all counts, only the latest timestamp => misses hits in between scraps
- less intuitive than increasing count
- ties it to timestamp database

1. `deployed{app=<name>,namespace=<name>}=<numOfDeployments>`:

- atomic count, requires history
- conceptually easier to understand than `UnixTimeOfDeployment` approach

=> 2. but with in-memory limitation, negligible in the long-run as we're primarily interested in recent agility (keep it above certain threshold)

Example:

1. Get the increase in number of deployments (for an app), e.g. for the last 24h: `increase(deployed_count[24h])`

2. Compute the frequency, e.g. for the last week: `increase(deployed_count[7d])/7`:

### Design idea

3 layers: K8s controller watching K8s deployments -> storing data in map <- node exporter for Prometheus exporting data/frequency

[Scheduling](https://prometheus.io/docs/instrumenting/writing_exporters/#scheduling): "*Metrics should only be pulled from the application when Prometheus scrapes them, exporters should not perform scrapes based on their own timers. That is, all scrapes should be synchronous.*"

### Prerequisites

Local:

- Make
- Go 1.14+
- kubectl v1.20.2
- Helm v3.4.2

K8s:

- [Kubernetes 1.19.1 installed](docs/K8S.md) with [Prometheus and Grafana deployed](docs/MONITORING.md)

## Import Prometheus Dashboard in Grafana

Under Grafana -> Dashboards -> Manage -> [Import](http://localhost:3000/dashboard/import) (assumes Grafana to be available under [http://localhost:3000](helper/expose-grafana.sh)):

- Upload [these dashboards](grafana-dashboards/)
- Optional: load dashboard ID `1860`

## Build

### deployments watcher

- [Add client-go as a dependency](https://github.com/jtestard/client-go/blob/master/INSTALL.md#add-client-go-as-a-dependency):

```bash
# pwd is ~/go/src/lttl.dev/agility/deployments-watcher
$ make init
go mod init
go: creating new go.mod: module lttl.dev/agility/deployments-watcher
$ make get
go get k8s.io/client-go@v0.19.1
# ...
```

## Bookmarks

### Major Resources

#### K8s

Watcher:

- [Building stuff with the Kubernetes API (Part 4) - Using Go](https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899) and related [pvcwatch in go](https://github.com/vladimirvivien/k8s-client-examples/blob/master/go/pvcwatch/main.go)
- client-go
  - [client-go](https://github.com/kubernetes/client-go)
  - [CUD k8s-deployment in main.go](https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go)

ConfigMap:

- [K8S Read config map via go API](https://stackoverflow.com/questions/59234194/k8s-read-config-map-via-go-api)
- [A simple Go client for Kubernetes](https://github.com/ericchiang/k8s#a-simple-go-client-for-kubernetes)

#### Prometheus Node Exporter

- [Writing exporters](https://prometheus.io/docs/instrumenting/writing_exporters/)
- [Naming conventions](https://prometheus.io/docs/practices/naming/)
- [Best practices](https://prometheus.io/docs/practices/instrumentation/#things-to-watch-out-for)
- [A Noob's Guide to Custom Prometheus Exporters](https://rsmitty.github.io/Prometheus-Exporters/)
- [A Noob's Guide to Custom Prometheus Exporters (Revamped!)](https://rsmitty.github.io/Prometheus-Exporters-Revamp/)
- [Prometheus' node_exporter.go](https://github.com/prometheus/node_exporter/blob/master/node_exporter.go)
- kind: [Loading an Image Into Your Cluster](https://kind.sigs.k8s.io/docs/user/quick-start/#loading-an-image-into-your-cluster)

### K8s API Documentation

- [k8s client-go apps-v1](https://godoc.org/k8s.io/client-go/kubernetes/typed/apps/v1)
- [k8s api-v1](https://godoc.org/k8s.io/api/core/v1)

#### Backlog

- client-go
  - [client-go examples](https://github.com/kubernetes/client-go/tree/master/examples)
  - [Available clientsets](https://github.com/kubernetes/client-go/blob/master/kubernetes/clientset.go)
  - [Deployment informer](https://github.com/kubernetes/client-go/blob/master/informers/apps/v1/deployment.go)
- [Controllers architecture](https://kubernetes.io/docs/concepts/architecture/controller/)
- [Sample controller](https://github.com/kubernetes/sample-controller)
- [A deep dive into Kubernetes controllers](https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html)
- Kubewatch
  - [Kubewatch, an example of Kubernetes custom controller](https://engineering.bitnami.com/articles/kubewatch-an-example-of-kubernetes-custom-controller.html)
  - [Kubewatch controller.go](https://github.com/bitnami-labs/kubewatch/blob/master/pkg/controller/controller.go)
- [Using Kubernetes API from Go](https://rancher.com/using-kubernetes-api-go-kubecon-2017-session-recap)
- [How to write Kubernetes custom controllers in Go](https://medium.com/speechmatics/how-to-write-kubernetes-custom-controllers-in-go-8014c4a04235)

### Reading List

- PromQL [Query examples](https://prometheus.io/docs/prometheus/latest/querying/examples/)

## Deployment

```bash
# from this project directory
$ make install
```

## Access

```bash
# forward local port 8088 to df-frontend port (80), to enable: curl localhost:8088
$ helper/expose-df-frontend.sh
# forward local port 8089 to df-backend (node-exporter) port (8080), to enable: curl localhost:8089/metrics
$ helper/expose-df-backend.sh
```

## Configure DF Node Exporter in Prometheus

Add the following snippet to `prometheus-server`'s Configmap under *scrape_configs*:

```yaml
scrape_configs:
- job_name: 'df_node_exporter_metrics'
  scrape_interval: 5s # for testing only
  metrics_path: /metrics
  static_configs:
    - targets:
      - agility-df-backend.agility.svc:8080
  # optional, remove system-generated labels:
  metric_relabel_configs:
    - source_labels: [ job ]
      target_label: job
      action: replace
      replacement: ''
    - source_labels: [ instance ]
      target_label: instance
      action: replace
      replacement: ''
```

Remove the `prometheus-server` pod to force a restart.

## Testing

Run a K8s cluster locally.

```bash
# shell 1
$ make all
# ...
$ go run main.go
```

```bash
# shell 2
# run some dummy deployments ...
$ helper/deploy-dummy-apps.sh
```

## Troubleshooting

### Helm

```bash
# from this project folder
$ helm install --namespace agility --debug --dry-run agility ./chart
```
