import bluebird from 'bluebird'
import {ACCOUNTS, getAccount,} from '../utils/accounts'
import {connectToken, DEFAULT_TOKENS_PER_ACCOUNT, transferTokens,} from '../utils/web3'

async function main() {
    const token = await connectToken()
    await bluebird.mapSeries(ACCOUNTS, async (toAccount) => {
        if (toAccount.name === 'admin') return
        await transferTokens(getAccount('admin'), toAccount, DEFAULT_TOKENS_PER_ACCOUNT)
    })
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
