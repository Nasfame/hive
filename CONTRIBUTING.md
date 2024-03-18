# Development Setup

Before starting development, ensure that the following dependencies are installed on your system:

- Docker
- Node.js
- Golang

## Installation

First, install the required Node.js modules and generate a local `.env` file containing private keys for various
services. Run the following commands:

```bash
(cd hardhat && pnpm install)
./setup print-env > .env
```

## Booting the Stack

### 1 - Bacalhau

To run a Bacalhau node on the same machine as the resource provider, follow these steps:

```bash
curl -sL https://get.bacalhau.org/install.sh | sudo bash
# Set the IPFS data path by exporting the `BACALHAU_SERVE_IPFS_PATH` variable to your desired location
export BACALHAU_SERVE_IPFS_PATH=/tmp/hive/data/ipfs
mdir -p $BACALHAU_SERVE_IPFS_PATH
# Run Bacalhau as both a compute node and a requester node
bacalhau serve --node-type compute,requester --peer none --private-internal-ipfs=false --job-selection-accept-networked --web-ui --web-ui-port 1080
```

### 2 - CoopHive

To initiate the boot sequence for CoopHive, run the following command:

```bash
./setup boot
```

This command performs the following four phases within the boot sequence:

#### 2.1 - Geth

During this phase, the following tasks are executed:

```bash
./setup geth
./setup fund-admin
./setup fund-services-ether
./setup balances
```

These commands start Geth, allocates eth to the admin, faucet, solver, mediator, resource_provider, job_creator, and
directory accounts.

#### 2.2 - Compile Contracts

```bash
./setup compile-contracts
```

This compiles the smart contracts and generates Go bindings in `pkg/contract/bindings/contracts`.

#### 2.3 - Deploy Contracts

```bash
./setup deploy-contracts
```

This deploys the smart contracts. Note that services will not have any tokens at this point.

#### 2.4 - Fund Tokens

```bash
./setup fund-services-tokens
./setup balances
```

This funds the services with tokens and prints the balances.

### 3 - Run Services

Run the following commands in separate terminal windows:

```bash
./setup solver
```

Wait for the solver to start when `🟡 SOL solver registered` is logged, and then run:

```bash
./setup mediator
```

If you have a GPU, run the following command in a separate terminal window:

```bash
./setup rp --offer-gpu 1
```

Otherwise, if you don't have a GPU:

```bash
./setup rp
```

Run Cowsay:

```bash
./setup run cowsay:v0.0.1 -i Message="moo"
```

Run SDXL:

```bash
./setup runsdxl sdxl:v0.1.0 PROMPT="beautiful view of iceland with a record player"
```

### 4 - Run Cowsay On-Chain

Start the on-chain Job Creator:

```bash
./setup jc
```

```bash
./setup run-cowsay-onchain
```

## Stopping the Stack

To stop Geth at any time, use the following command:

```bash
./setup geth-stop
```

To reset Geth data, effectively performing a complete restart, use the following command:

```bash
./setup clean
```

Please note that after running `clean`, you will need to re-run the `fund-admin` and `fund-services` commands.

## Unit Tests

Run the smart contract unit tests with the following command:

```bash
./setup unit-tests
```

## Regenerating Go Bindings

Whenever you make changes to the smart contracts, regenerate the Go bindings in `pkg/contract/bindings/contracts` by
running:

```bash
./setup compile-contracts
```

## Troubleshooting

### Receive Buffer Size Error

`failed to sufficiently increase receive buffer size`
See https://github.com/quic-go/quic-go/wiki/UDP-Receive-Buffer-Size for details. Fix for Linux:

```
sudo sysctl -w net.core.rmem_max=2500000
sudo sysctl -w net.core.wmem_max=2500000
```
