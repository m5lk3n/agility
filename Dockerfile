# docker build -t lttl.dev/k8s-df .
FROM alpine:3.13.0
COPY df-backend/df-backend /
COPY df-frontend/static/ /static/
COPY df-frontend/df-frontend /
ENV DF_LOG="debug"
# ENV DF_LOG="info"
ENV GIN_MODE=release
ENV PORT=80
COPY version.txt /
COPY start.sh /
CMD ["/start.sh"]