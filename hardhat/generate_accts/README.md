# Generate Ethereum Accounts

## Create a folder locally to store the keystore

```
sudo mkdir -p /mnt/hive/keystore
```

## Build the docker image

```
cd hardhat/generate_accts
docker build -t generate-accts-image .
```

## Run the generate-accts docker image and remove the container after exiting

```
docker run -it --rm --name generate-accts-container -v /mnt/hive/keystore:/root/.ethereum/keystore generate-accts-image
```

## Create seven new accounts

Use the same password for all accounts:

```
geth account new
geth account new
geth account new
geth account new
geth account new
geth account new
geth account new
```

## Get the public and private key for each account

```
./extractkeys <password>
```

## Update `.env` file in `hardhat` directory

Add the addresses and private keys

## Keystore files are locally stored at /mnt/hive/keystore

```
ls /mnt/hive/keystore
```
