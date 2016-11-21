# This is how we want to name the binary output
BINARY=go-rabbitmq

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X github.com/sklarsa/go-rabbitmq/core.Version=${VERSION} -X github.com/sklarsa/go-rabbitmq/core.BuildTime=${BUILD_TIME}"

build:
	go build ${LDFLAGS} -o ${BINARY} ./

all: build test_api
	go build ${LDFLAGS} -o ${BINARY} ./

get_test_deps:
	go get github.com/streadway/amqp

test_api: get_test_deps
	go test ./api
