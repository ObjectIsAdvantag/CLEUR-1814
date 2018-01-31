FROM golang:latest

COPY constants/secrets.yml /go/src/github.com/ObjectIsAdvantag/CLEUR-1814/5-webserver/constants/secrets.yml
COPY microservice/main.linux /usr/local/bin/ciscolive.service.content
COPY entrypoint.sh /

RUN GOPATH=/go
RUN export GOPATH

ENTRYPOINT ["/entrypoint.sh"]