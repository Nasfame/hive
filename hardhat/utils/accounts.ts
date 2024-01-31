// IMPORTANT: we cannot import hardhat directly here
// because it will cause a circular dependency
import {Account} from "./types";
import {Wallet} from "ethers";

export const loadEnv = (name: string, defaultValue: string) => {
    return process.env[name] || defaultValue;
};

export const loadPrivateKey = (name: string, defaultValue: string) => {
    return loadEnv(`${name.toUpperCase()}_PRIVATE_KEY`, defaultValue).trim();
};

export const loadAddress = (name: string, privateKey: string) => {
    let address = loadEnv(`${name.toUpperCase()}_ADDRESS`, "").trim();

    if (!address) {
        try {
            const wallet = new Wallet(privateKey);
            address = wallet.address;
            console.log(name + ": " + address);
        } catch (error: any) {
            console.error(
                `Error deriving address from private key for ${name}: ${error.message}`
            );
        }
    }

    return address;
};

// the default values here are the hardhat defualt insecure accounts
// this means that we get a reproducable dev environment between hardhat and geth
export const ACCOUNTS: Account[] = [
    {
        name: "admin",
        privateKey: loadPrivateKey(
            "admin",
            "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
        ),
    metadata: {},
    },
    {
        name: "faucet",
        privateKey: loadPrivateKey(
            "faucet",
            "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
        ),
    metadata: {},
    },
    {
        name: "solver",
        privateKey: loadPrivateKey(
            "solver",
            "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
        ),
    metadata: {},
    },
    {
        name: "mediator",
        privateKey: loadPrivateKey(
            "mediator",
            "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
        ),
    metadata: {},
    },
    {
        name: "resource_provider",
        privateKey: loadPrivateKey(
            "resource_provider",
            "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a"
        ),
    metadata: {},
    },
    {
        name: "job_creator",
        privateKey: loadPrivateKey(
            "job_creator",
            "0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba"
        ),
    metadata: {},
    },
    {
        name: "directory",
        privateKey: loadPrivateKey(
            "directory",
            "0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e"
        ),
    metadata: {},
    },
].map((account: Account) => {
    // Check if the address is not already present
    if (!account.address) {
        // Derive the address using the loadAddress function
        account.address = loadAddress(account.name, account.privateKey);
    }
    return account;
});

// map of account name -> account
export const NAMED_ACCOUNTS = ACCOUNTS.reduce<Record<string, Account>>(
    (all, acc) => {
        all[acc.name] = acc;
        return all;
    },
    {}
);

// map of account name -> account address
export const ACCOUNT_ADDRESSES = ACCOUNTS.reduce<Record<string, string>>(
    (all, acc) => {
        all[acc.name] = acc.address;
        return all;
    },
    {}
);

// flat list of private keys in order
export const PRIVATE_KEYS = ACCOUNTS.map((acc) => acc.privateKey);

export const getAccount = (name: string) => {
    const account = NAMED_ACCOUNTS[name];
    if (!account) {
        throw new Error(`Unknown account ${name}`);
    }
    return account;
};
