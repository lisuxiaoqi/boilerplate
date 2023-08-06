// 引入web3.js库
const Web3 = require('web3');

const fs = require('fs');

// Read the contents of the JSON file
const jsonData = fs.readFileSync('../build/contracts/Simple.json');

// Parse the JSON data
const contractJSON = JSON.parse(jsonData);

// 连接以太坊节点
const web3 = new Web3('https://okbtestrpc.okbchain.org');

const contractAddress = contractJSON.networks["195"].address

console.log(contractAddress)

const personalAddress = "0xA1869624A6EdAcdB03cD970719B20c737c248d2a"

const personalKey = "13977ec5c2fd6f2fa064a54919ea6f1b2efaf2982082e14560cf7536bd1ad670"

const contractABI = contractJSON.abi


const contract = new web3.eth.Contract(contractABI, contractAddress);

async function run(){
  const txCount = await web3.eth.getTransactionCount(personalAddress);
  const txObject = {
    from: personalAddress,
    to: contractAddress,
    data: contract.methods.set(100).encodeABI(),
    nonce: web3.utils.toHex(txCount),
    gas: 2000000
  };

  // 签名交易
  const signedTx = await web3.eth.accounts.signTransaction(txObject, personalKey);

  // 发送交易
  const txResult = await web3.eth.sendSignedTransaction(signedTx.rawTransaction);

  console.log("txResult:", txResult)

  const callResult = await contract.methods.get().call()

  console.log("Call result1:", callResult)

  const callResult2 = await web3.eth.call({
    to:contractAddress,
    data:contract.methods.get().encodeABI()
  })
  console.log("Call result2:", callResult)
}

run()

