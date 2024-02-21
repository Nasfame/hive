import {getAccount} from "./accounts";
import {HardhatRuntimeEnvironment} from "hardhat/types";
import fs from "fs";

export async function syncDapps(hre: HardhatRuntimeEnvironment) {
    const {deployments, network} = hre
    const controllerContract = await deployments.get("HiveController");
    const storageContract = await deployments.get("HiveStorage");
    const usersContract = await deployments.get("HiveUsers");
    const mediationContract = await deployments.get("HiveMediationRandom");
    const paymentsContract = await deployments.get("HivePayments");
    const jobCreatorContract = await deployments.get("HiveOnChainJobCreator");
    const tokenContract = await deployments.get("HiveToken");

    // @ts-ignore
    const netUrl = hre.network.config.url ?? "http://localhost:8545";
    // @ts-ignore
    let websocketUrl = hre.network.config.ws ?? netUrl.replace('http', 'ws');

    const solvers = [
        getAccount("solver").address,
        // getAccount("admin").address,
    ]

    const content = `
HIVE_CONTROLLER=${controllerContract.address}
HIVE_MEDIATION_RANDOM=${mediationContract.address}
HIVE_SOLVER=${solvers.join(",")}
WEB3_RPC_URL=${websocketUrl} 
WEB3_RPC_HTTP=${netUrl}
WEB3_CHAIN_ID=${network.config.chainId}

HIVE_TOKEN=${tokenContract.address}
`.trim();

    // the below can be derived from controller contract
    /***
     HIVE_STORAGE=${storageContract.address}
     HIVE_PAYMENT=${paymentsContract.address}
     HIVE_JOBCREATOR=${jobCreatorContract.address}
     */

    console.log(content);

    writeToFile(content, `../config/dApps/${network.name}.env`);
}

function writeToFile(data: string, filename: string) {
    fs.writeFileSync(filename, data);

    console.log(`Wrote to ${filename}`);
}
