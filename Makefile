build:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" .

release:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" -o bin/
	./bin/hive version


.PHONY: release install-unix install-win build

install-unix:
	make release
	ln -s ./bin/hive $$(go env GOBIN)

install-win:
	make release
	cp ./bin/hive.exe $$GOBIN
#Ps1: cmd	cp ./bin/hive.exe $env:GOBIN
