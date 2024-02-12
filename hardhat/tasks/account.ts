import {task} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {getBalance, getPublicAddress} from "../utils/accounts";

task("balance", "Prints an account's balance")
    .addParam("account", "The account's address")
    .setAction(async ({account}, hre) => {
        console.log("network", hre.network.name);
        const balance = await hre.ethers.provider.getBalance(account);

        console.log(hre.ethers.formatEther(balance), "ETH");
    });

task("bal", "Prints an account's balance")
    .addPositionalParam("account", "The account'address")
    .setAction(async ({account}, hre) => {
        console.log("network", hre.network.name);
        const balance = await getBalance(account, hre);

        console.log(hre.ethers.formatEther(balance), "ETH");
    });

task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners();

    console.log("Loading accounts");

    for (const account of accounts) {
        const bal = await getBalance(account.address, hre);
        const balString = hre.ethers.formatEther(bal) + "ETH";
        console.log("acc", account.address, balString);
    }
});

task("account", "Prints account address from private key")
    .addPositionalParam("privateKey", "The private key")
    .setAction(async ({privateKey}, hre) => {
        const address = getPublicAddress(privateKey, hre);
        console.log("account:", address);
    });
