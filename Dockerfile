
FROM golang:1.21 AS builder

LABEL authors="Hiro <laciferin@gmail.com>"
LABEL maintainer="Hiro <laciferin@gmail.com>"

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR /app
COPY . .
RUN mkdir -p ./bin

RUN go get

RUN go install github.com/goreleaser/goreleaser@latest
RUN goreleaser build --single-target --clean -o ./bin/hive --snapshot
RUN ./bin/hive version


FROM golang:1.21-alpine

WORKDIR /app

COPY --from=builder /app/bin/hive  /app/hive

ENTRYPOINT ["/app/hive"]
CMD ["run", "cowsay:v0.1.0", "-i", "Message=Hiro"]