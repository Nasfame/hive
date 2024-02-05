import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";

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

    return true;
};

deployController.id = "deployController";

export default deployController;
