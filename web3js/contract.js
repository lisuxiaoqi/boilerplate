// 引入web3.js库
const Web3 = require('web3');

// 连接以太坊节点
const web3 = new Web3('http://localhost:8545');

const contractAddress = "0xe5CC80FC5BD03655c4BaD29C6cDCf39dec240190"

const personalAddress = "0x8d63Bd4B972794b67aA83CDaf09dB5655b4e00CC"

const personalKey = "08a59286ea6759517e9cd2f01faf9625f0c7502d5401656a2ef5f3121c977f82"

const contractABI = [
  {
    "inputs": [],
    "name": "get",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "x",
        "type": "uint256"
      }
    ],
    "name": "set",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]

const contract = new web3.eth.Contract(contractABI, contractAddress);

async function run(){
  const txCount = await web3.eth.getTransactionCount(personalAddress);
  const txObject = {
    from: personalAddress,
    to: contractAddress,
    data: contract.methods.set(15).encodeABI(),
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
  console.log("Call result2:",  web3.utils.toDecimal(callResult2))
}

run()

