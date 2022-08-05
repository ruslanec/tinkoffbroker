# Dockerfile for gRPC proto files compile for golang
FROM golang:1.18 AS builder

RUN apt-get update
RUN apt-get install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /home
RUN apt-get install -y git 

RUN git clone https://github.com/Tinkoff/investAPI.git

WORKDIR /home/investAPI/src/docs/contracts/
RUN protoc --go_out=. ./*.proto
RUN protoc --go-grpc_out=. ./*.proto

FROM builder AS commiter

ARG email
ARG username

RUN apt-get install -y git

WORKDIR /home
RUN --mount=type=secret,id=my_env ls /run/secrets/
RUN --mount=type=secret,id=my_env cat /run/secrets/my_env

RUN git config --global user.email $email
RUN git config --global user.name $username
RUN git clone https://github.com/ruslanec/tinkoffbroker.git
COPY --from=builder /home/investAPI/src/docs/contracts/* /home/tinkoffbroker/proto/

WORKDIR /home/tinkoffbroker
RUN git add .
RUN git commit -m "update proto files"

SHELL ["/bin/bash", "-c"] 
RUN --mount=type=secret,id=my_env source /run/secrets/my_env \
    && git push https://${GIT_TOKEN}@github.com/ruslanec/tinkoffbroker.git