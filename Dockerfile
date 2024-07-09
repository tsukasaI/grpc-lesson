FROM golang:1.22.5-bookworm
ARG PROTOBUF_VERSION=3.20.3

RUN apt-get update && apt-get install unzip

# protocのダウンロード
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-aarch_64.zip -o protoc.zip && \
    unzip protoc.zip && \
    mv bin/* /usr/local/bin/ && \
    mv include/* /usr/local/include/

WORKDIR /app

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
