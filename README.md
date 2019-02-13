# go-meetup

To run:

Add reddit credentials to `src/python/input-stream/app.env`
then:
```
cd src/python/grpc-base/
docker build . -t grpc-base
cd ../../../
docker-compose build
docker-compose up
```
