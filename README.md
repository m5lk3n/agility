# k8s-df

**`export IMAGE_VER` before flight!**

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

- `deployed{app=<name>,namespace=<name>}=<UnixTimeOfWhen>?` (observation time? deployment time?)

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

- Kubernetes 1.19.1 with Prometheus and Grafana deployed:
  
  | app name | chart version | app version |
  | --- | --- | --- |
  | grafana | grafana-6.1.17 | 7.3.5 |
  | prometheus | prometheus-11.12.1 | 2.20.1 |

#### Install kind

kind v0.7.0+ is required, but set up was v0.9.0 (which comes with Kubernetes 1.19.1):

```bash
$ curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64
$ chmod +x kind
$ mv kind ~/bin
$ kind create cluster
# ...
```

#### Install Prometheus

```bash
$ kubectl create ns monitoring
$ helm repo update
$ helm repo add stable https://charts.helm.sh/stable
$ helm install prometheus stable/prometheus --namespace monitoring
#
# ...
#
# Get the Prometheus server URL by running these commands in the same shell:
#  export POD_NAME=$(kubectl get pods --namespace monitoring -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
#  kubectl --namespace monitoring port-forward $POD_NAME 9090
#
# ...
#
# To uninstall:
#   $ helm delete prometheus --namespace monitoring
#
```

#### Install Grafana

```bash
$ kubectl create ns monitoring
$ helm repo add grafana https://grafana.github.io/helm-charts
$ helm repo update
$ helm install grafana grafana/grafana --namespace monitoring
#
# ...
#
# To uninstall:
#   $ helm delete grafana --namespace monitoring
#
```

#### Configure Prometheus as a Data Source in Grafana

![Connect Grafana Prometheus](prometheus_as_data_source_in_grafana.png)

#### Import Prometheus Dashboard in Grafana

Under [Import](http://localhost:3000/dashboard/import) load dashboard ID `1860`.

## Build

### deployments watcher

- [Add client-go as a dependency](https://github.com/jtestard/client-go/blob/master/INSTALL.md#add-client-go-as-a-dependency):

```bash
michael@x250:~/go/src/lttl.dev/k8s-df/deployments-watcher
$ make init
go mod init
go: creating new go.mod: module lttl.dev/k8s-df/deployments-watcher
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

- [Understanding go.sum and go.mod file in Go (Golang)](https://golangbyexample.com/go-mod-sum-module/)

## Deploying

```bash
$ cd k8s
$ ./k8s
$ export POD_NAME=$(kubectl get pods --namespace k8s-df -l "app=k8s-df" -o jsonpath="{.items[0].metadata.name}")
# forward local port 8088 to web-app port (80), to enable: curl localhost:8088
$ kubectl --namespace k8s-df port-forward $POD_NAME 8088:80
# forward local port 8089 to node-exporter port (8080), to enable: curl localhost:8089/metrics
$ kubectl --namespace k8s-df port-forward $POD_NAME 8089:8080
```

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
# deploy ...
$ kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
# ... delete ...
$ kubectl delete deployment hello-node
# ... deploy
$ kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
```
