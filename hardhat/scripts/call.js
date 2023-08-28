const { ethers } = require("hardhat");

async function call(contractAddress, deployer) {
  const Contract = await ethers.getContractFactory("Sample",deployer);
  const contract = await Contract.attach(contractAddress);
  const number = 22;
  console.log("Write to contract:", number);
  const tx = await contract.set(number);

  receipt = await tx.wait();
  console.log(receipt);

  const counter = await contract.get();
  console.log("Read from contract:", counter);
}

async function main(){
  const contractAddress = "0x0b41A792d2baFAC1f2b64032Bc2aa6502D694Aa7";
  const [deployer] = await ethers.getSigners();
  console.log(
      "Deploying contracts with the account:",
      deployer.address
  );

  const address = await call(contractAddress, deployer);
}

main() .then(() => process.exit(0))
    .catch(error => {
      console.error(error);
      process.exit(1);
    });
