const { ethers } = require("hardhat");
const web3 = require('web3');
async function deploy(deployer) {
  const Contract = await ethers.getContractFactory("Sample", deployer);
  const contract = await Contract.deploy();
  await contract.deployed();
  console.log("Contract deployed to address:", contract.address);
  return contract.address;
}

async function main(){
  const [deployer] = await ethers.getSigners();
  console.log(
      "Deploying contracts with the account:",
      deployer.address
  );

  const address = await deploy(deployer);
}

main() .then(() => process.exit(0))
    .catch(error => {
      console.error(error);
      process.exit(1);
    });


