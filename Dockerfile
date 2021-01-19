# docker build -t lttl.dev/k8s-df .
FROM alpine:3.13.0
# DEV-only add curl:
RUN apk --no-cache add curl
COPY deployments-watcher/deployments-watcher /
COPY node-exporter/node-exporter /
COPY web-app/static/ /static/
COPY web-app/ /
ENV GIN_MODE=release
ENV PORT=80
COPY start.sh /
CMD ["/start.sh"]