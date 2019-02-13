# go-meetup

To run:

Add reddit credentials to `src/python/input-stream/app.env`
then:
```
cd src/go
GOOS=linux GOARCH=amd64 go build -o pipeline/service pipeline/cmd/pipeline/main.go
cd ../python/grpc-base/
docker build . -t grpc-base
cd ../../../
docker-compose build
docker-compose up
```
(sorry for the lack of a makefile)
