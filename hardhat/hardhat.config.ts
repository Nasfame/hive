import {HardhatUserConfig, task} from "hardhat/config";
import "@typechain/hardhat";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-ethers";
import "@nomicfoundation/hardhat-chai-matchers";
import "hardhat-deploy";
import * as dotenv from "dotenv";
import * as process from "process";

const ENV_FILE = process.env.CONFIG || "./.env";

console.log(`ENV_FILE is ${ENV_FILE}`);

dotenv.config({ path: ENV_FILE });

import {ACCOUNT_ADDRESSES, PRIVATE_KEYS} from "./utils/accounts";

const NETWORK = process.env.NETWORK || "hardhat";
const INFURA_KEY = process.env.INFURA_KEY || "";

console.log(`infura key is ${INFURA_KEY}`);

const config: HardhatUserConfig = {
  solidity: "0.8.21",
  defaultNetwork: "hardhat",
  namedAccounts: ACCOUNT_ADDRESSES,
  networks: {
    hardhat: {
      chainId: 1337,
      accounts: [
        {
          privateKey:
            process.env.PRIVATE_KEY ||
            "beb00ab9be22a34a9c940c27d1d6bfe59db9ab9de4930c968b16724907591b3f",
          balance: `${1000000000000000000000000n}`,
        },
        ...PRIVATE_KEYS.map((privateKey) => {
          return {
            privateKey: privateKey,
            balance: `${1000000000000000000000000n}`,
          };
        }),
      ],
    },
    geth: {
      url: "http://localhost:8545",
      chainId: 1337,
      accounts: PRIVATE_KEYS,
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_KEY}`,
      chainId: 11155111,
      accounts: PRIVATE_KEYS,
    },
    calibration: {
      chainId: 314159,
      url: "https://api.calibration.node.glif.io/rpc/v1",
      accounts: PRIVATE_KEYS,
    },
    fvm: {
      chainId: 314,
      url: "https://api.node.glif.io",
      accounts: PRIVATE_KEYS,
    },
    coophive: {
      chainId: 1337,
      url: "http://testnet.co-ophive.network:8545",
      accounts: PRIVATE_KEYS,
    },
    chaos: {
      //skale testnet
      chainId: 1351057110,
      url: "https://staging-v3.skalenodes.com/v1/staging-fast-active-bellatrix",
      accounts: PRIVATE_KEYS,
      // faucet: "https://sfuel.skale.network/staging/chaos",
      // explorer: "https://staging-fast-active-bellatrix.explorer.staging-v3.skalenodes.com",
    },
    titanAI: {
      //skale testnet
      chainId: 1020352220,
      url: "https://testnet.skalenodes.com/v1/aware-fake-trim-testnet",
      accounts: PRIVATE_KEYS,
      // https://testnet.portal.skale.space/chains/titan
    },
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  // sourcify: {
  //     // Disabled by default
  //     // Doesn't need an API key
  //     enabled: true
  // }
};

import "./tasks";

module.exports = config;
