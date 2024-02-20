import {task} from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {getAccount, getBalance} from "../utils/accounts";
import * as fs from "fs";
import {syncDapps} from "../utils/syncDapps";

task("dapp", "Sync Dapp: always pass network its a must")
    .setAction(async ({}, hre) => {
        await syncDapps(hre)
    });


