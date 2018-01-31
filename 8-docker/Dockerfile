FROM alpine:3.5

LABEL Description="Golang Docker Demo For CLEUR"

RUN apk update && \
    apk upgrade && \
    apk add \
        bash \
        ca-certificates \
    && rm -rf /var/cache/apk/*

COPY ./bin/docker-golang-demo.linux /usr/local/bin/docker-golang-demo
ENTRYPOINT /usr/local/bin/docker-golang-demo