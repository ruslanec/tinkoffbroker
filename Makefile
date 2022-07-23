.PHONY: all
all: build

.PHONY: build_proto
build_proto:
	docker build -t protobuilder ./Dockerfile.build.proto

.PHONY: vet
vet:
	go vet ./cmd/tinkoffbroker/

.PHONY: build
build:
	go build -ldflags="-X 'main.Version=v1.0.0'" -o ./bin/ ./cmd/tinkoffbroker/

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