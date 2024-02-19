import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";
import {DEFAULT_TOKEN_SUPPLY} from "../utils/web3";

const deployToken: DeployFunction = async function (
    hre: HardhatRuntimeEnvironment,
) {
    const {deployments, getNamedAccounts} = hre;
    const {deploy} = deployments;
    const {admin} = await getNamedAccounts();
    // log the admin address
    console.log(`admin: ${admin}`);
    await deploy("HiveToken", {
        from: admin,
        args: ["Hive Token", "HIVE", DEFAULT_TOKEN_SUPPLY],
        log: true,
        waitConfirmations: hre.network.config.chainId == 1337 ? 1 : 6,
    });
    return true;
};

deployToken.id = "deployToken";

export default deployToken;

module.exports.tags = ["all", "hiveToken"];
