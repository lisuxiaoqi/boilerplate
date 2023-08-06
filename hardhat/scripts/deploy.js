const { ethers } = require("hardhat");

async function main() {
  const Contract = await ethers.getContractFactory("Sample");
  const contract = await Contract.deploy();
  console.log("Contract deployed to address:", contract.address);
}

main();

