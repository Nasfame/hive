
FROM golang:1.21 AS builder

LABEL authors="Hiro <laciferin@gmail.com>"
LABEL maintainer="Hiro <laciferin@gmail.com>"

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR /app
COPY . .

RUN go get
#RUN go install github.com/goreleaser/goreleaser@latest
# TODO: this works, make sure to match the bin name : RUN  goreleaser build --single-target --clean -o bin/hive --snapshot
RUN mkdir -p ./bin
RUN make build-ci


FROM golang:1.21-alpine

WORKDIR /app

COPY --from=builder /app/bin/hive-ci  /app/hive

ENTRYPOINT ["/app/hive"]
CMD ["run", "cowsay:v0.1.0", "-i", "Message=Hiro"]