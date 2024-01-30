binName = hive-$(shell uname -s)-$(shell uname -m)

build-ci:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" -o ./bin/hive-ci .
	./bin/hive-ci version

export VERSION=$(git describe --tags --abbrev=0)
export COMMIT_SHA=$(git rev-parse HEAD)

build:
	goreleaser build --single-target --clean -o bin/hive --snapshot

prerelease:
	echo "Version is $(VERSION)"
	goreleaser check
	goreleaser build --single-target --clean

release:
	goreleaser release --clean


make-bin:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/cmd/hive.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/cmd/hive.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" -o bin/$(binName)
	./bin/$(binName) version

release-linux:
	sh scripts/release-linux.sh

.PHONY: release install-unix install-win build release release-linux make-bin


#	ln -s ./bin/hive $$(go env GOBIN)
install-win:
	make release
	cp ./bin/$(binName) ./bin/hive.exe
	cp ./bin/hive.exe $$GOBIN
#Ps1: cmd	cp ./bin/hive.exe $env:GOBIN
