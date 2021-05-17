# Makefile for build multiple target binaries
#

hello:
	echo "hello, world!"

build:
	go build -o bin/kafkee examples/kafkee_example/kafkee_example.go
	go build -o bin/kafkor examples/kafkor_example/kafkor_example.go

fmt:
	go fmt ...

