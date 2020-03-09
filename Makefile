init:
	go install golang.org/x/tools/cmd/goimports

fmt:
	goimports -w .
	go fmt ./...

test: fmt
	go test ./...

build: test
	go build -o main
	docker build --tag proxeus/node-proof-of-existence .

push:
	docker push proxeus/node-proof-of-existence:latest