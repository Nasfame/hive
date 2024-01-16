binName = hive-$(shell uname -s)-$(shell uname -m)

build:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" .

release:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" -o bin/$(binName)
	./bin/$(binName) version


.PHONY: release install-unix install-win build

install-linux:
	export GOOS=linux
	export GOARCH=amd64
	make release
#	ln -s ./bin/hive $$(go env GOBIN)

install-win:
	make release
	cp ./bin/$(binName) ./bin/hive.exe
	cp ./bin/hive.exe $$GOBIN
#Ps1: cmd	cp ./bin/hive.exe $env:GOBIN
