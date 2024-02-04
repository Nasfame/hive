#!/bin/sh


export GOOS=linux
export GOARCH=amd64
export binName=hive-$GOOS-$GOARCH

go build -v -ldflags="\
  -X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
  -X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
" -o bin/$binName
./bin/$binName version
./bin/$binName run cowsay:v0.0.1 -i Message="Hiro"