const { ethers } = require("hardhat");

async function main() {
  const contractAddress = "0xc7a09F62Ff29f9ccD903400CaF0a43B08be6fccF";
  const Contract = await ethers.getContractFactory("Sample");
  const contract = await Contract.attach(contractAddress);
  const ret = await contract.set(33);
  console.log(ret) 
  const counter = await contract.get();
  console.log("Counter:", counter);
}

main()
