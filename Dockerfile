# Dockerfile for gRPC proto files compile for golang
FROM golang:1.18 AS builder

RUN apt-get update
RUN apt-get install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /home
RUN apt-get install -y git 

RUN git clone https://github.com/Tinkoff/investAPI.git

#RUN cp /home/investAPI/src/docs/contracts/* /home/tinkoffbroker/proto/

WORKDIR /home/investAPI/src/docs/contracts/
RUN protoc --go_out=. ./*.proto
RUN protoc --go-grpc_out=. ./*.proto

FROM golang:1.18 AS commiter

ARG email
#ENV GIT_EMAIL=$email

ARG username
#ENV GIT_NAME=$username
#ARG token

RUN apt-get install -y git
SHELL ["/bin/bash", "-c"] 

WORKDIR /home
RUN --mount=type=secret,id=my_env ls /run/secrets/
RUN --mount=type=secret,id=my_env cat /run/secrets/my_env
RUN --mount=type=secret,id=my_env source /run/secrets/my_env
RUN git config --global user.email $email
RUN git config --global user.name $username
RUN git clone https://github.com/ruslanec/tinkoffbroker.git
COPY --from=builder /home/investAPI/src/docs/contracts/* /home/tinkoffbroker/proto/

WORKDIR /home/tinkoffbroker
RUN git add .
RUN git commit -m "update proto files"
RUN git push https://${GIT_TOKEN}@github.com/ruslanec/tinkoffbroker.git