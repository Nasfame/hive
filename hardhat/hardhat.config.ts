import {HardhatUserConfig, task} from "hardhat/config";
import "@typechain/hardhat";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-chai-matchers";
import "@nomicfoundation/hardhat-ethers";
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
      saveDeployments: true,
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
        ws: `wss://sepolia.infura.io/ws/v3/${INFURA_KEY}`,
        chainId: 11155111,
        accounts: PRIVATE_KEYS,
        saveDeployments: true,
    },
    /* coophive: {
       chainId: 1337,
       url: "http://aurora.co-ophive.network:8545",
       ws: "ws://aurora.co-ophive.network:8546",
       accounts: PRIVATE_KEYS,
       saveDeployments: true,
     },*/
    aurora: {
      chainId: 1337,
      url: "http://aurora.co-ophive.network:8545",
      ws: "ws://aurora.co-ophive.network:8546",
      accounts: PRIVATE_KEYS,
      saveDeployments: true,
    },
    halcyon: { //coophive testnet
      chainId: 1337,
      url: "http://halcyon.co-ophive.network:8545",
      ws: "ws://halcyon.co-ophive.network:8546",
      faucet: "http://halcyon-faucet.co-ophive.network:8085",
      accounts: PRIVATE_KEYS,
      saveDeployments: true,
    },
    /*chaos: {
      //skale testnet
      chainId: 1351057110,
      url: "https://staging-v3.skalenodes.com/v1/staging-fast-active-bellatrix",
      accounts: PRIVATE_KEYS,
      // faucet: "https://sfuel.skale.network/staging/chaos",
      // explorer: "https://staging-fast-active-bellatrix.explorer.staging-v3.skalenodes.com",
    },*/
    /*    titanAI: {
          //skale testnet
          chainId: 1020352220,
          url: "https://testnet.skalenodes.com/v1/aware-fake-trim-testnet",
          ws: "wss://testnet.skalenodes.com/v1/ws/aware-fake-trim-testnet",
          accounts: PRIVATE_KEYS,
          saveDeployments: true,
          // https://testnet.portal.skale.space/chains/titan
          // https://aware-fake-trim-testnet.explorer.testnet.skalenodes.com
        },*/
        calibration: {
          chainId: 314159,
          url: "https://api.calibration.node.glif.io/rpc/v1",
          ws: 'wss://wss.calibration.node.glif.io/apigw/lotus/rpc/v1',
          accounts: PRIVATE_KEYS,
          saveDeployments: true,
        },
    /*fvm: {
      chainId: 314,
      url: "https://api.node.glif.io",
      accounts: PRIVATE_KEYS,
    },*/
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  // sourcify: {
  //     // Disabled by default
  //     // Doesn't need an API key
  //     enabled: true
  // }
  typechain: {
    outDir: 'typechain-types',
    target: 'ethers-v6',
    alwaysGenerateOverloads: false, // should overloads with full signatures like deposit(uint256) be generated always, even if there are no overloads?
    externalArtifacts: ['externalArtifacts/*.json'], // optional array of glob patterns with external artifacts to process (for example external libs from node_modules)
    dontOverrideCompile: false // defaults to false
  },
};

import "./tasks";

module.exports = config;
