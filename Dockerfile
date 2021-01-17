# docker build -t lttl.dev/k8s-df .

FROM alpine:3.7
COPY deployments-watcher/deployments-watcher /
COPY node-exporter/node-exporter /
COPY start.sh /
CMD ["/start.sh"]