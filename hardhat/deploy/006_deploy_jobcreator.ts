import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";

const deployJobCreator: DeployFunction = async function (
    hre: HardhatRuntimeEnvironment,
) {
    const {deployments, getNamedAccounts} = hre;
    const {deploy, execute} = deployments;
    const {admin, solver} = await getNamedAccounts();
    await deploy("HiveOnChainJobCreator", {
        from: admin,
        args: [],
        log: true,
    });

    await deploy("ExampleClient", {
        from: admin,
        args: [],
        log: true,
    });

    const tokenContract = await deployments.get("HiveToken");
    const jobCreator = await deployments.get("HiveOnChainJobCreator");

    await execute(
        "HiveOnChainJobCreator",
        {
            from: admin,
            log: true,
        },
        "initialize",
        tokenContract.address,
    );

    await execute(
        "ExampleClient",
        {
            from: admin,
            log: true,
        },
        "initialize",
        jobCreator.address,
    );

    // we set the controller of the job creator to be the solver
    // because it will be the one pulling jobs from it
    await execute(
        "HiveOnChainJobCreator",
        {
            from: admin,
            log: true,
        },
        "setControllerAddress",
        solver,
    );
    return true;
};

deployJobCreator.id = "deployJobCreator";

export default deployJobCreator;
