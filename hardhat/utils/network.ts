import hre from "hardhat";

export const getNetwork = () => {
    const network = hre.network;
    return network;
};
