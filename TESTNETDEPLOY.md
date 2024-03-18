# Testnet Deployment

## Installation

First, install the required Node.js modules and generate a local `.env` file containing private keys for various
services. Run the following commands:

```bash
(cd hardhat && pnpm install && pnpm gen-env)
./stack print-env > .env
```

## Booting the Stack

### 1 - Bacalhau

To run a Bacalhau node on the same machine as the resource provider, follow these steps:

```bash
# install the latest bacalhau which works with GPUs (https://github.com/bacalhau-project/bacalhau/issues/2858)
curl -sL https://get.bacalhau.org/install.sh | sudo bash
# configure this to where you want the ipfs data to be stored
export BACALHAU_SERVE_IPFS_PATH=/tmp/coophive/data/ipfs
# run bacalhau as both compute node and requester node
./stack bacalhau-serve
```

## Create Seven New Accounts

Follow the `README.md` in the `generate_accts` directory to create seven new accounts.

Copy `hardhat/.env.sample` to `.env` and update the private keys. 

> Hint: you can also use the same private key across all services, but the recommended way is to use different accounts.

## Create a new Infura Project

Create a new Infura project and update the following environment variable in `hardhat/.env`:

```
INFURA_KEY=
```

Also add the infura key to the `.env` file:

```
export INFURA_KEY=
```

## Setup Hardhat

Add the NETWORK to the `hardhat/hardhat.config.ts`.

## Fund the Seven New Accounts

Fund the `admin` acccount with .7 ETH.

Fund the remaining six accounts with .1 ETH each.

```bash
./stack fund-services-ether
```

Check the balances

```bash
./stack balances
```

## Compile Contracts

```bash
./stack compile-contracts
```

## Deploy Contracts

```bash
./stack deploy-contracts
```

## Fund Services Tokens

```bash
./stack fund-services-tokens
```

### Run Services

Run the following commands in separate terminals:

```bash
./stack solver
```

Wait for the solver to start when `🟡 SOL solver registered` is logged, and then run:

```bash
./stack mediator
```

If you have a GPU, run the following command in a separate terminal window:

```bash
./stack rp --offer-gpu 1
```

Otherwise, if you don't have a GPU:

```bash
./stack rp
```

Run Cowsay:

```bash
./stack run cowsay:v0.1.2 -i Message="Hiro welcomes you to his Hive"
```

Run SDXL:

```bash
./stack run sdxl:v0.3.0 -i Prompt="hiro saves the hive" -i Seed=16
```

### 4 - Run Cowsay On-Chain

Start the on-chain Job Creator:

```bash
./stack jc
```

```bash
./stack run-cowsay-onchain
```
