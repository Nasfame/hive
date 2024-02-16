
FROM golang:1.21 AS builder

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR /app
RUN mkdir -p ./bin

RUN go install github.com/goreleaser/goreleaser@latest

COPY go.* ./
RUN go get

COPY . .

RUN goreleaser build --single-target --clean -o ./bin/hive --snapshot

RUN ./bin/hive version


FROM alpine:latest

#ENV WEB3_PRIVATE_KEY; try to pass a hardhat private key here

WORKDIR /app

RUN mkdir -p ./coophive

ENV APP_DIR=/app/coophive

COPY --from=builder /app/bin/hive  /app/bin/hive

RUN ln -s /app/bin/hive /bin/hive

ENTRYPOINT ["/bin/hive"]
CMD ["run", "cowsay:v0.1.0", "-i", "Message=Hiro"]

LABEL authors="Hiro <laciferin@gmail.com>"
LABEL maintainer="Hiro <laciferin@gmail.com>"
