import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";
import * as fs from "fs";

import {network} from "hardhat";
import {getAccount} from "../utils/accounts";
import {execSync} from "child_process";


//TODO: use bindings
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
    const tokenContract = await deployments.get("HiveToken");

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


    execSync(`npx hardhat dapp --network ${hre.network.name}`)

    return true;
};

deployController.id = "deployController";

export default deployController;

function writeToFile(data: string, filename: string) {
    fs.writeFileSync(filename, data);

    console.log(`Wrote to ${filename}`);
}
