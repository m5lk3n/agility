# docker build -t lttl.dev/k8s-df .

FROM scratch
LABEL maintainer="m5l.k3n@gmail.com"
COPY deployments-watcher/deployments-watcher /
CMD ["/deployments-watcher"]