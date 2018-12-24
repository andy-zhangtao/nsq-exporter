
.PHONY: build
name = nsq-exporter
version = v0.1.0

build:
	go build -ldflags "-X main._VERSION_=$(shell date +%Y%m%d-%H%M%S)" -o $(name)

run: build
	./$(name)

release: *.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._BUILD_=$(shell date +%Y%m%d) -X main._VERSION_=$(version)" -a -o bin/$(name)
	docker build -t vikings/$(name):$(version) .
	docker push vikings/$(name):$(version)
