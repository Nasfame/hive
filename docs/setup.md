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
source .env
cd hardhat
pnpm account $SOLVER_PRIVATE_KEY
```

1. Copy the account address
2. Paste it to HIVE_SOLVER in the `.env` file

### Setup Metamask

1. **Install MetaMask**: If you haven't already, install the MetaMask extension for your web browser or the MetaMask
   mobile app from the respective app store.

2. **Open MetaMask**: Launch the MetaMask extension or app on your device.

3. **Access Settings**: Look for the settings menu in MetaMask. This is usually represented by a gear or three dots
   icon.

4. **Select "Import Account"**: In the settings menu, find the option to import an account. Click on it to proceed.

5. **Enter Private Key**: You'll be prompted to enter the private key associated with the account you want to import.
   Make sure you have the correct private key.

6. **Complete Import**: Follow the prompts to complete the import process. You may need to confirm your action with a
   password or additional verification.

7. **Verify and Access Imported Account**: Once the import process is complete, you should see the imported account in
   your MetaMask wallet along with any existing accounts you have.

8. **Ensure Security**: After importing your account, it's important to ensure the security of your MetaMask wallet.
   Make sure to keep your private key secure and never share it with anyone.

> Do the same for all the private keys

## Setup CoopHive Services

### Setup Solver

`hive solver`

### Setup Resource Provider

`hive rp`

### Run a coophive module on the rp

`hive run cowsay:v0.1`

##  

###  

[//]: # ()

[//]: # (> You can use private keys generated at `hardhat/.env.gen` in your .env file)

[//]: # (> Ensure you have got funds from your faucet for all addresses)

