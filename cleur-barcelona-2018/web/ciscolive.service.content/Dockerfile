FROM golang:latest

COPY constants/secrets.yml /go/src/cleur-barcelona-2018/web/ciscolive.service.content/constants/secrets.yml
COPY microservice/main.linux /usr/local/bin/ciscolive.service.content
COPY entrypoint.sh /

RUN GOPATH=/go
RUN export GOPATH

ENTRYPOINT ["/entrypoint.sh"]