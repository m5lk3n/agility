# Research

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

## Article Refs

- https://itrevolution.com/book/accelerate/ | https://www.google.com/books/edition/_/85XHAQAACAAJ?hl=en
- https://thenewstack.io/how-devops-affects-business-stakeholders-and-leaders/
- https://www.youtube.com/watch?v=iUFpRFvlT2U

From the latter:

DF ("throughput") per day per project

| Percentile | 2020 Value | 2019 Value |
| --- | --- | --- |
| 5p | 0.03 | 0.03 |
| 50p | 0.70 | 0.80 |
| 90p | 16.03 | 13.00 |
| 95p | 32.125 | 25.47 |
| **Mean** | **8.22** | **5.76** |

From [State of DevOps 2019](https://services.google.com/fh/files/misc/state-of-devops-2019.pdf):

"*Deployment frequencyThe elite group reported that it routinely deploys on-demand and performs multiple deployments per day, consistent with the last several years. By comparison, low performers reported deploying between once per month (12 per year) and once per six months (two per year), which is a decrease in performance from last year. The normalized annual deployment numbers range from 1,460 deploys per year **(calculated as four deploys per day x 365 days) for the highest performers** to seven deploys per year for low performers (average of 12 deploys and two deploys). Extending this analysis shows that elite performers deploy code 208 times more frequently than low performers. It's worth noting that four deploys per day is a conservative estimate when comparing against companies such as CapitalOne that report deploying up to 50 times per day for a product, **8 for companies such as Amazon, Google, and Netflix** that deploy thousands of times per day (aggregated over the hundreds of services that comprise their production environments).*"
