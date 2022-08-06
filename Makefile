.PHONY: all
all: build

.PHONY: update_proto
update_proto:
	docker build -f proto.Dockerfile -t proto \
	--build-arg email=r.idrisov@gmail.com \
	--build-arg username=ruslanec \
	--secret id=my_env,src=.env \
	--no-cache .

.PHONY: vet
vet:
	go vet .

.PHONY: win_build
win_build:
	go build .\examples\portfolio\portfolio.go 
	go build .\examples\marketdatastream\marketdatastream.go
	go build .\examples\users\users.go
	go build .\examples\errors\errors.go

.PHONY: win_clean
clean:
	rm ./portfolio.exe
	rm ./marketdatastream.exe
	rm ./users.exe

.PHONY: build
build:
	go build ./examples/portfolio/portfolio.go 
	go build ./examples/marketdatastream/marketdatastream.go
	go build ./examples/users/users.go
	go build ./examples/errors/errors.go


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