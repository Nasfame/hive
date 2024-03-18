set shell := ["sh", "-c"]
set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]
#set allow-duplicate-recipe
set positional-arguments
#set dotenv-load
#set dotenv-filename := ".env"
#set export


default:
    go build

release:
	go build -v -ldflags="\
		-X 'github.com/CoopHive/hive/config.version=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/CoopHive/hive/config.commitSha=$$(git rev-parse HEAD)' \
	" .
	./hive version

echo:
    @echo just

setup-geth:
    echo "Setup geth"


go-bindings:
    #!/usr/bin/env bash
    function go-binding() {
      local name="$1"
      local pkg="$2"

      # compile the sol files into bytecode and ABI
      cd hardhat
      solc \
        --base-path . \
        --include-path node_modules \
        --overwrite \
        --abi --bin \
        "contracts/$name.sol" \
        -o artifacts


    #  sudo chown -R $USER:$USER hardhat/artifacts
      mkdir -p artifacts/bindings/$pkg

      # generate the go bindings
      abigen \
        "--bin=artifacts/$name.bin" \
        "--abi=artifacts/$name.abi" \
        "--pkg=$pkg" "--out=artifacts/bindings/$pkg/$pkg.go"

      cd ..

    #  sudo chown -R $USER:$USER hardhat/artifacts/bindings/$pkg
    #  sudo chmod 0644 hardhat/artifacts/bindings/$pkg/$pkg.go
      cp -r hardhat/artifacts/bindings/$pkg pkg/web3/bindings/$pkg

      echo "Generated go binding hardhat/artifacts/bindings/$pkg/$pkg.go"
    }

    function go-bindings() {
      # check if the hive-solc image exists
      # and only build it if it doesn't

      rm -rf pkg/web3/bindings
      mkdir -p pkg/web3/bindings
      go-binding HiveToken token
      go-binding HivePayments payments
      go-binding HiveStorage storage
      go-binding HiveUsers users
      go-binding HiveMediationRandom mediation
      go-binding HiveOnChainJobCreator jobcreator
      go-binding HiveController controller

      echo "Generated all go bindings pkg/contract/bindings/"
    }

    function go-bindings-docker() {
      # check if the hive-solc image exists
      # and only build it if it doesn't

      if [[ -z $(docker images -q coophive-solc) ]]; then
        docker build -t coophive-solc hardhat/solc
      fi
      rm -rf pkg/web3/bindings
      mkdir -p pkg/web3/bindings
      go-binding-docker HiveToken token
      go-binding-docker HivePayments payments
      go-binding-docker HiveStorage storage
      go-binding-docker HiveUsers users
      go-binding-docker HiveMediationRandom mediation
      go-binding-docker HiveOnChainJobCreator jobcreator
      go-binding-docker HiveController controller

      echo "Generated all go bindings pkg/contract/bindings/"
    }

    go-bindings