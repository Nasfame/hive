# Setup hive  on your machine  [For devs]

## Setup repo

```shell
gh repo clone Coophive/hive
cd hive
```

```bash
make
cp .env.example .env
./hive version
```

## Setup env vars for services

### Generate addresses for services

```shell
cd hardhat
pnpm gen-env
```

1. Open `hardhat/.env.gen`
2. Pick the private keys and paste it appropriately to .env file

### For `HIVE_SOLVER`

```shell
cd hardhat
pnpm account $PRIVATE_KEY
```

> paste the address you


[//]: # ()

[//]: # (> You can use private keys generated at `hardhat/.env.gen` in your .env file)

[//]: # (> Ensure you have got funds from your faucet for all addresses)

> To get address from

```bash


```