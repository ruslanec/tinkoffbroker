.PHONY: all
all: build

.PHONY: build_proto
build_proto:
	docker build -t proto \
	--build-arg email=r.idrisov@gmail.com \
	--build-arg username=protobuilder \
	--secret id=my_env,src=.env \
	--no-cache .

.PHONY: vet
vet:
	go vet ./cmd/tinkoffbroker/

.PHONY: build
build:
	go build .\examples\portfolio\portfolio.go 
	go build .\examples\marketdatastream\marketdatastream.go
	go build .\examples\users\users.go

.PHONY: clean
clean:
	del portfolio.exe
	del marketdatastream.exe
	del users.exe

.PHONY: build_compose
build_compose:
	docker-compose build myapp

.PHONY: build_docker
build_docker:
	docker build -t myapp .

.PHONY: docker_lint
docker_lint:
	docker run --rm -i hadolint/hadolint hadolint --ignore DL3008 - < .\Dockerfile.build.proto

.PHONY: run
run:
	docker-compose up myapp

.PHONY: test
test:
	go test -v ./...