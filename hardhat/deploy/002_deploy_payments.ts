import {HardhatRuntimeEnvironment} from 'hardhat/types'
import {DeployFunction} from 'hardhat-deploy/types'

const deployPayments: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
    const {deployments, getNamedAccounts} = hre
    const {deploy, execute} = deployments
    const {
        admin,
    } = await getNamedAccounts()
    await deploy("HivePayments.sol", {
        from: admin,
        args: [],
        log: true,
    })

    const tokenContract = await deployments.get('HiveToken.sol')
    const paymentsContract = await deployments.get('HivePayments.sol')

    await execute(
        'HivePayments.sol',
        {
            from: admin,
            log: true,
        },
        'initialize',
        tokenContract.address
    )

    await execute(
        'HiveToken.sol',
        {
            from: admin,
            log: true,
        },
        'setControllerAddress',
        paymentsContract.address
    )

    return true
}

deployPayments.id = 'deployPayments'

export default deployPayments