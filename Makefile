include .env
export

binName = hive-$(shell uname -s)-$(shell uname -m)


setup-dev:
	go install github.com/goreleaser/goreleaser@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	cd hardhat
	pnpm install
	cd ..
	go generate

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

.PHONY: release install-unix install-win build release release-linux make-bin install install-all install-hive-latest install-hive


install:
	goreleaser build --single-target --clean -o ./bin/${binName} --snapshot

snapshot:
	goreleaser build --clean --snapshot

sync:
	make snapshot

	scp dist/cli_linux_amd64_v1/bin hive:/usr/local/bin/hive
	scp dist/cli_linux_amd64_v1/bin hive1:/usr/local/bin/hive

	scp .env.prod hive:.env
	scp .env.prod hive1:.env

sync-env:
	scp .env.prod hive:.env
	scp .env.prod hive1:.env

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

setup-bacalhau:
	mkdir -p /tmp/coophive/data/ipfs
	export BACALHAU_SERVE_IPFS_PATH=/tmp/coophive/data/ipfs
	bacalhau serve \
        --node-type compute,requester \
        --peer none \
        --private-internal-ipfs=false \
        --job-selection-accept-networked

	#bacalhau serve --node-type requester --private-internal-ipfs --peer none

test-b:
	 echo "make sure you have pasted the env vars otherwise it points to the public bacalhau cluster"
	# sudo docker run alpine echo hello
	bacalhau docker run alpine echo hi;

install-all:
	curl -sSf https://raw.githubusercontent.com/CoopHive/hive/main/install.sh | sh -s -- all


install-hive-latest:
	curl -sSf https://raw.githubusercontent.com/CoopHive/hive/main/install.sh | sh -s -- hive


b:
	make setup-bacalhau

h:
	make install-hive-latest

cowsay:
	hive run cowsay:v0.1.0

.PHONY: solver run cowsay h b

solver:
	echo $$SOLVER_PRIVATE_KEY
	hive solver --web3-private-key $$SOLVER_PRIVATE_KEY

run:
	hive run cowsay:v0.1.0


rp:
	hive rp

