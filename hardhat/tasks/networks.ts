import {task} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {getBalance} from "../utils/accounts";


import {exec} from "child_process";
import hre from "hardhat";
import {getNetworks} from "../utils/network";
import {deployToAllNetworks} from "../scripts/deployAll";



task("networks", "Show all network")
    .setAction(async ({}, hre) => {
        let networkNames = getNetworks(hre)
        console.log(`Networks: ${networkNames.join(', ')}`);
        // console.log(hre.config.networks);
        for (const name of networkNames) {
            const network = hre.config.networks[name];
            console.log(`\nNetwork: ${name}`);
            // @ts-ignore
            console.log(`URL: ${network.url ?? 'http://localhost:8545'}`);
            // Access other properties of the network configuration as needed
        }
    });


interface deployArgs {
    skip?: string
    skipNetworks?: string
}

task("deployAll", "Deploy to network")
    .addOptionalParam("skip", "The networks to skip put in, separated string")
    .addOptionalParam("skipNetworks", "The networks to skip put in, separated string")
    .setAction(async ({skip, skipNetworks}: deployArgs, hre) => {


        let totalSkipNetworks = skip ?? '' + skipNetworks ?? ''
        let skippedNetworks = totalSkipNetworks.split(',') ?? []

        // skippedNetworks = skippedNetworks.concat(['localhost', 'chaos'])

        console.log("skipNetworks", skippedNetworks)

        await deployToAllNetworks(hre, skippedNetworks)

        console.log("Deployment completed")

    });
