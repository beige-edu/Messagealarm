FROM alpine:latest

LABEL maintainer="jinfengyi@51beige.com"

RUN apk add --update --no-cache tzdata

RUN mkdir -p app/logs
RUN touch app/logs/prometheusalertcenter.log

COPY bin/PrometheusAlert app/PrometheusAlert
COPY conf/app-example.conf app/conf/app.conf
COPY conf/cert app/conf/cert
COPY static app/static
COPY views app/views

EXPOSE 13013

WORKDIR /app
CMD ["./PrometheusAlert"]