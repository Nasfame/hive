import {task} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {getAccount, getBalance, getBalanceInEther, getPublicAddress} from "../utils/accounts";
import {HiveToken} from "../typechain-types";
import {Account} from "../utils/types";

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


interface fundOut {
    address: string,
    balance: string,
    hive: string,
    newHive?: string,
    newBalance?: string,
}

interface Out {
    admin: fundOut,
    faucet?: fundOut,
    solver?: fundOut,
}
task("fund", "Fund RP, faucet account's balance")
    .addOptionalPositionalParam("eth", "The Eth drip", ".01")
    .addOptionalPositionalParam("hive", "The Hive Drip", "8000")
    .setAction(async ({eth, hive}, hre) => {
        console.log("network", hre.network.name);

        // const balance = await getBalance(, hre);

        // console.log(hre.ethers.formatEther(balance), "ETH");

        const fundingAccount = getAccount("admin");
        const signer = await hre.ethers.getSigner(fundingAccount.address)

        console.log("Funding account:", fundingAccount.address);

        console.log("Funding account bal:", await getBalanceInEther(fundingAccount.address, hre));

        const rcvAccounts = [
            getAccount("faucet"),
            getAccount("mediator"),
            getAccount("solver"),
            getAccount("resource_provider"),
            getAccount(""),
        ]

        const amountInWei = hre.ethers.parseEther(eth);

        const token = await hre.deployments.get("HiveToken");

        console.log("token Address:", token.address)

        const tokenContract: HiveToken = new hre.ethers.Contract(token.address, token.abi, signer)


        const amountInHiveWei = hre.ethers.parseUnits(hive, await tokenContract.decimals()); // Assuming 'token' is the HiveToken contract and 'decimals' is the number of decimal places of the token

        const hiveBal = async (address: string) => await tokenContract.balanceOf(address) + "HIVE"
        const ethBal = async (address: string) => await getBalanceInEther(address, hre)


        let out: Out = {
            admin: {
                address: fundingAccount.address,
                balance: await getBalanceInEther(fundingAccount.address, hre),
                hive: await tokenContract.balanceOf(fundingAccount.address)
            }
        }

        for (const acc of rcvAccounts) {
            // console.log("acc:" + acc.name);
            // console.log("accAddr:" + acc.address)

            const out_: fundOut = {
                address: acc.address,
                balance: await ethBal(acc.address),
                hive: await hiveBal(acc.address),
            }
            out[acc.name] = out_

            const tx = {
                to: acc.address,
                value: amountInWei
            };
            const transactionResponse = await signer.sendTransaction(tx);
            const transactionReceipt = await transactionResponse.wait();
            const {hash: txHash} = transactionReceipt;

            console.log(`Transaction successful: ${txHash}`);

            // Send Hive tokens
            const transferTx = await tokenContract.transfer(acc.address, amountInHiveWei);
            const transferReceipt = await transferTx.wait();
            const {hash: transferTxHash} = transferReceipt;

            console.log(`Hive transfer successful: ${transferTxHash}`);


            out_.newHive = await hiveBal(acc.address)
            out_.newBalance = await ethBal(acc.address)
        }

        out.admin.newBalance = await getBalanceInEther(fundingAccount.address, hre)
        out.admin.newHive = await hiveBal(fundingAccount.address)


        console.table(out)


    });


task("drip", "Drip RP, faucet account's balance")
    .addPositionalParam('address', "The address to drip to")
    .addOptionalPositionalParam("eth", "The amount to drip", "0.01")
    .addOptionalPositionalParam("hive", "The amount to drip", "100")
    .setAction(async ({address, eth, hive}, hre) => {
        console.log("network", hre.network.name);

        const fundingAccount = getAccount("admin");
        const signer = await hre.ethers.getSigner(fundingAccount.address)

        console.log("Funding account:", fundingAccount.address);

        console.log("Funding account bal:", await getBalanceInEther(fundingAccount.address, hre));


        const amountInWei = hre.ethers.parseEther(eth);

        const token = await hre.deployments.get("HiveToken");

        console.log("token Address:", token.address)

        const tokenContract: HiveToken = new hre.ethers.Contract(token.address, token.abi, signer)

        const amountInHiveWei = hre.ethers.parseUnits(hive, await tokenContract.decimals()); // Assuming 'token' is the HiveToken contract and 'decimals' is the number of decimal places of the token
        const hiveBal = async (address: string) => await tokenContract.balanceOf(address) + "HIVE"
        const ethBal = async (address: string) => await getBalanceInEther(address, hre)


        let out: Out = {
            admin: {
                address: fundingAccount.address,
                balance: await getBalanceInEther(fundingAccount.address, hre),
                hive: await tokenContract.balanceOf(fundingAccount.address)
            }
        }
        const acc: Account = {
            name: "dev",
            address: address,
        }


        const out_: fundOut = {
            address: acc.address,
            balance: await ethBal(acc.address),
            hive: await hiveBal(acc.address),
        }
        out[acc.name] = out_

        const tx = {
            to: acc.address,
            value: amountInWei
        };

        if (hre.network.name == "titanAI") {
            // const gasLimit = await signer.estimateGas(tx);
            tx["gasLimit"] = 100000 //+gasLimit //1.5 times or 250K
        }
        const transactionResponse = await signer.sendTransaction(tx);
        const transactionReceipt = await transactionResponse.wait();
        const {hash: txHash} = transactionReceipt;

        console.log(`Transaction successful: ${txHash}`);


        const additonalParams = {}

        if (hre.network.name == "titanAI") {
            additonalParams["gasLimit"] = 100000 //+gasLimit //1.5 times or 250K
        }

        // Send Hive tokens
        const transferTx = await tokenContract.transfer(acc.address, amountInHiveWei, additonalParams);
        const transferReceipt = await transferTx.wait();
        const {hash: transferTxHash} = transferReceipt;

        console.log(`Hive transfer successful: ${transferTxHash}`);


        out_.newHive = await hiveBal(acc.address)
        out_.newBalance = await ethBal(acc.address)
        out.admin.newBalance = await getBalanceInEther(fundingAccount.address, hre)
        out.admin.newHive = await hiveBal(fundingAccount.address)


        console.table(out)


    });
