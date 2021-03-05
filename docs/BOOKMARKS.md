# Bookmarks

## Major Resources

### K8s

Watcher:

- [Building stuff with the Kubernetes API (Part 4) - Using Go](https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899) and related [pvcwatch in go](https://github.com/vladimirvivien/k8s-client-examples/blob/master/go/pvcwatch/main.go)
- client-go
  - [client-go](https://github.com/kubernetes/client-go)
  - [CUD k8s-deployment in main.go](https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go)

ConfigMap:

- [K8S Read config map via go API](https://stackoverflow.com/questions/59234194/k8s-read-config-map-via-go-api)
- [A simple Go client for Kubernetes](https://github.com/ericchiang/k8s#a-simple-go-client-for-kubernetes)

### Prometheus Node Exporter

- [Writing exporters](https://prometheus.io/docs/instrumenting/writing_exporters/)
- [Naming conventions](https://prometheus.io/docs/practices/naming/)
- [Best practices](https://prometheus.io/docs/practices/instrumentation/#things-to-watch-out-for)
- [A Noob's Guide to Custom Prometheus Exporters](https://rsmitty.github.io/Prometheus-Exporters/)
- [A Noob's Guide to Custom Prometheus Exporters (Revamped!)](https://rsmitty.github.io/Prometheus-Exporters-Revamp/)
- [Prometheus' node_exporter.go](https://github.com/prometheus/node_exporter/blob/master/node_exporter.go)
- [Scheduling](https://prometheus.io/docs/instrumenting/writing_exporters/#scheduling): "*Metrics should only be pulled from the application when Prometheus scrapes them, exporters should not perform scrapes based on their own timers. That is, all scrapes should be synchronous.*"

### kind

- [Loading an Image Into Your Cluster](https://kind.sigs.k8s.io/docs/user/quick-start/#loading-an-image-into-your-cluster)

### K8s API Documentation

- [k8s client-go apps-v1](https://godoc.org/k8s.io/client-go/kubernetes/typed/apps/v1)
- [k8s api-v1](https://godoc.org/k8s.io/api/core/v1)

## Miscellaneous

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

## Unnamed

- https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/
- https://kind.sigs.k8s.io/docs/user/quick-start/  
- https://godoc.org/github.com/prometheus/client_golang/prometheus
- https://github.com/prometheus/client_golang/blob/master/prometheus/examples_test.go
- https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899
- https://github.com/vladimirvivien/k8s-client-examples/tree/master/go
- https://prometheus.io/docs/instrumenting/writing_exporters/
- https://github.com/ericchiang/k8s#a-simple-go-client-for-kubernetes
- https://github.com/kubernetes/client-go/blob/master/listers/core/v1/configmap.go
- https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/configmap.go
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
- https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#-em-deployment-em-
- https://medium.com/containerum/configuring-permissions-in-kubernetes-with-rbac-a456a9717d5d
- https://kubernetes.io/docs/reference/access-authn-authz/rbac/
- https://prometheus.io/docs/instrumenting/writing_exporters/
- https://rsmitty.github.io/Prometheus-Exporters/
- https://rsmitty.github.io/Prometheus-Exporters-Revamp/
- https://github.com/prometheus/node_exporter/blob/master/node_exporter.go
- https://stackoverflow.com/questions/47138461/get-total-requests-in-a-period-of-time
- https://thenewstack.io/how-devops-affects-business-stakeholders-and-leaders/
