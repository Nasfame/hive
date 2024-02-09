# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.21 AS builder

LABEL authors="Hiro <laciferin@gmail.com>"
LABEL maintainer="Hiro <laciferin@gmail.com>"

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR /app
COPY . .
RUN mkdir -p ./bin

ARG TARGETOS
ARG TARGETARCH

RUN go env -w CGO_ENABLED=0 \
    && go get \
    && go install github.com/goreleaser/goreleaser@latest \
    && goreleaser build --single-target --clean -o ./bin/hive --snapshot \
    && ./bin/hive version

FROM --platform=$BUILDPLATFORM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/hive /app/hive

ENTRYPOINT ["/app/hive"]
CMD ["run", "cowsay:v0.1.0", "-i", "Message=Hiro"]
