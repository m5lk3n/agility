# docker build -t lttl.dev/k8s-df .
FROM alpine:3.13.0
COPY deployments-watcher/deployments-watcher /
COPY node-exporter/node-exporter /
COPY web-app/static/ /static/
COPY web-app/ /
ENV DF_LOG="color,debug"
#ENV DF_LOG="info"
#ENV GIN_MODE=release
ENV PORT=80
COPY start.sh /
CMD ["/start.sh"]