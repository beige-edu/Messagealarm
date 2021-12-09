FROM alpine:latest

LABEL maintainer="jinfengyi@51beige.com"

RUN apk add --update --no-cache tzdata

COPY bin/PrometheusAlert app/PrometheusAlert
COPY conf/app.conf app/conf/app.conf
COPY conf/cert app/conf/cert
COPY logs app/logs
COPY static app/static
COPY views app/views

EXPOSE 13013

WORKDIR /app
CMD ["./PrometheusAlert"]