import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";

const deployController: DeployFunction = async function (
    hre: HardhatRuntimeEnvironment
) {
    const {deployments, getNamedAccounts} = hre;
    const {deploy, execute} = deployments;
    const {admin} = await getNamedAccounts();
    await deploy("HiveController.sol", {
        from: admin,
        args: [],
        log: true,
    });

    const controllerContract = await deployments.get("HiveController.sol");
    const storageContract = await deployments.get("HiveStorage.sol");
    const usersContract = await deployments.get("HiveUsers.sol");
    const mediationContract = await deployments.get("HiveMediationRandom.sol");
    const paymentsContract = await deployments.get("HivePayments.sol");
    const jobCreatorContract = await deployments.get("HiveOnChainJobCreator.sol");

    await execute(
        "HiveController.sol",
        {
            from: admin,
            log: true,
        },
        "initialize",
        storageContract.address,
        usersContract.address,
        paymentsContract.address,
        mediationContract.address,
        jobCreatorContract.address
    );

    await execute(
        "HiveStorage.sol",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        controllerContract.address
    );

    await execute(
        "HivePayments.sol",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        controllerContract.address
    );

    await execute(
        "HiveMediationRandom.sol",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        controllerContract.address
    );

    return true;
};

deployController.id = "deployController";

export default deployController;
