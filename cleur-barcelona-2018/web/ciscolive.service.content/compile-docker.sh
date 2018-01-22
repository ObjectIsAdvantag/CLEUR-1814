rm microservice/main.linux
env GOOS=linux GOARCH=amd64 go build -o microservice/main.linux -v microservice/main.go
docker build . -f Dockerfile --tag ciscolive.service.content:latest