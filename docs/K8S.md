# Kubernetes

In my case, I used [kind](https://github.com/kubernetes-sigs/kind) (*k*ubernetes *in* *d*ocker) for this project to have a Kubernetes 1.19.1 at hand. Probably, any other flavor should do it.

## Install kind

kind v0.7.0+ is required, but my setup was on v0.9.0 (which comes with Kubernetes 1.19.1), installed as follows:

```bash
$ curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64
$ chmod +x kind
$ mv kind ~/bin
$ kind create cluster
# ...
```
