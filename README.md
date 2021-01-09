# k8s-df

## Prerequisites

- Make
- Go 1.14+
- Kubernetes 1.19.1
- [Add client-go as a dependency](https://github.com/jtestard/client-go/blob/master/INSTALL.md#add-client-go-as-a-dependency):

```bash
michael@x250:~/go/src/lttl.dev/k8s-df
$ make init
go mod init
go: creating new go.mod: module lttl.dev/k8s-df
$ make get
go get k8s.io/client-go@v0.19.1
# ...
```

## Bookmarks

### Major Resources

- [Building stuff with the Kubernetes API (Part 4) - Using Go](https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899)
- client-go
  - [client-go](https://github.com/kubernetes/client-go)
  - [CUD k8s-deployment in main.go](https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go)

### API doc

- [k8s client-go apps-v1](https://godoc.org/k8s.io/client-go/kubernetes/typed/apps/v1)
- [k8s api-v1](https://godoc.org/k8s.io/api/core/v1)

### Backlog

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

## Reading List

- [Understanding go.sum and go.mod file in Go (Golang)](https://golangbyexample.com/go-mod-sum-module/)

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
$ kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
```
