include .env
export

binName = hive-$(shell uname -s)-$(shell uname -m)

build-ci:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
	" -o ./bin/hive-ci .
	./bin/hive-ci version

export VERSION=$(git describe --tags --abbrev=0)
export COMMIT_SHA=$(git rev-parse HEAD)

build:
	goreleaser build --single-target --clean -o bin/hive1 --snapshot

prerelease:
	echo "Version is $(VERSION)"
	goreleaser check
	goreleaser build --single-target --clean

release:
	goreleaser release --clean

install-hive:
	goreleaser build --single-target --clean -o ./bin/${binName} --snapshot

make-bin:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
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


generate-sol-bindings-for-go:
	./stack go-bindings;


plugin-autoacceptdealer:
	#this plugin is very platform specific and can only be build in that specific platform
	./stack build-plugin-autoaccept;


plugin-websocket:
	./stack build-plugin-websocket


cleanup_github:
	git tag -l | grep pr | xargs -I {} sh -c 'git tag -d {} & git push origin --delete {} & gh release delete {} --yes'
	#git tag -l | grep v0.0.0-br | xargs -I {} sh -c 'git tag -d {} & git push origin --delete {} & gh release delete {} --yes'

github:
	exit 1
	#manuall triggers
	gh workflow run .github/workflows/publish-gcr.yml --ref v0.2.6
	git tag -l | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | xargs -I {} sh -c "gh workflow run .github/workflows/publish-gcr.yml --ref {}"
