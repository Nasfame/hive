# Check if .env file exists
ifeq ($(wildcard .env),)
    ENV_EXISTS := false
else
    ENV_EXISTS := true
endif

# Include .env file if it exists
ifeq ($(ENV_EXISTS),true)
    include .env
    export
endif



binName ?= hive-$(shell uname -s)-$(shell uname -m)


setup-dev:
	go install github.com/goreleaser/goreleaser@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	cd hardhat && pnpm install && pnpm gen-env
	cd ..
	go generate
	go build

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

make-bin:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
	" -o bin/$(binName)
	./bin/$(binName) version

release-linux:
	sh scripts/release-linux.sh

.PHONY: release install-unix install-win build release release-linux make-bin install install-all install-hive-latest install-hive

deps:
	go mod tidy && go work sync

install:
ifeq ($(OS),Windows_NT)
	@goreleaser build --clean --snapshot
else
	@goreleaser build --single-target --clean -o ./bin/$(binName) --snapshot
endif

snapshot:
	goreleaser build --clean --snapshot

host ?= hive
hiveDir ?= /tmp/

sync:
	#host=${host:-"hive"}
	echo "host is ${host}"
	make snapshot
	#scp dist/hive_linux_amd64_v1/hive hive:/usr/local/bin/hive #permission issue
	#scp dist/hive_linux_amd64_v1/hive hive:./bin/hive

	scp dist/hive_linux_amd64_v1/hive ${host}:${hiveDir}
#	ssh ${host} 'cd ${hiveDir} && sudo chmod +x hive && sudo cp hive /usr/local/bin/'
	ssh ${host} 'cd ${hiveDir} && sudo chmod +x hive && sudo rsync --force ./hive /usr/local/bin/'


	scp *.yml ${host}:.
	scp .env.* ${host}:.

	#scp dist/hive_linux_amd64_v1/hive hive1:/usr/local/bin/hive

	#scp .env.prod hive:.env
	#scp .env.prod hive1:.env
sync-solver:
	make sync host=solver

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
	./setup go-bindings;


plugin-autoacceptdealer:
	#this plugin is very platform specific and can only be build in that specific platform
	./scripts/build-plugin autoaccept


plugin-websocket:
	./scripts/build-plugin websocket

cleanup_github:
	git tag -l | grep -- "-pr[0-9]" | xargs -I {} sh -c 'git tag -d {} & git push origin --delete {} & gh release delete {} --yes'
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
        --job-selection-accept-networked \
       	--web-ui \
       	--web-ui-port 1080

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
	hive run cowsay:v0.1.2

.PHONY: solver run cowsay h b

solver:
	echo $$SOLVER_PRIVATE_KEY
	hive solver --web3-private-key $$SOLVER_PRIVATE_KEY

run:
	hive run cowsay:v0.1.2


rp:
	hive rp


test:
	cd test
	go test -v -count 1 .


sdxl-subt:
	go run . run sdxl:v0.3.0-alpha.1 -i Prompt="hiro saves the hive" -i Seed=20;

deploy-sepolia:
	docker-compose up sepolia -d --wait
	export CONFIG_FILE=.env.sepolia
	hive solver
