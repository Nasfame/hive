import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";
import * as fs from "fs";

import {network} from "hardhat";
import {getAccount} from "../utils/accounts";

const deployController: DeployFunction = async function (
    hre: HardhatRuntimeEnvironment,
) {
    const {deployments, getNamedAccounts} = hre;
    const {deploy, execute} = deployments;
    const {admin} = await getNamedAccounts();
    await deploy("HiveController", {
        from: admin,
        args: [],
        log: true,
    });

    const controllerContract = await deployments.get("HiveController");
    const storageContract = await deployments.get("HiveStorage");
    const usersContract = await deployments.get("HiveUsers");
    const mediationContract = await deployments.get("HiveMediationRandom");
    const paymentsContract = await deployments.get("HivePayments");
    const jobCreatorContract = await deployments.get("HiveOnChainJobCreator");

    await execute(
        "HiveController",
        {
            from: admin,
            log: true,
        },
        "initialize",
        storageContract.address,
        usersContract.address,
        paymentsContract.address,
        mediationContract.address,
        jobCreatorContract.address,
    );

    await execute(
        "HiveStorage",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        controllerContract.address,
    );

    await execute(
        "HivePayments",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        controllerContract.address,
    );

    await execute(
        "HiveMediationRandom",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        controllerContract.address,
    );

    console.log("Deployed Contracts:");

    // @ts-ignore
    const netUrl = hre.network.config.url ?? "http://localhost:8545"

    const content = `
HIVE_CONTROLLER=${controllerContract.address}
HIVE_MEDIATION_RANDOM=${mediationContract.address}
HIVE_SOLVER=${getAccount("solver").address}
WEB3_RPC_URL=${netUrl} 
WEB3_CHAIN_ID=${network.config.chainId}
`.trim();

    // the below can be derived from controller contract
    /***
     HIVE_STORAGE=${storageContract.address}
     HIVE_PAYMENT=${paymentsContract.address}
     HIVE_JOBCREATOR=${jobCreatorContract.address}
     */

    console.log(content);

    writeToFile(content, `../config/dApps/${network.name}.env`);

    return true;
};

deployController.id = "deployController";

export default deployController;

function writeToFile(data: string, filename: string) {
    fs.writeFileSync(filename, data);

    console.log(`Wrote to ${filename}`);
}
