// 引入web3.js库
const Web3 = require('web3');

// 连接以太坊节点
const web3 = new Web3('http://localhost:8545');

const contractAddress = "0xe547d6DF6C6efD549Ac71b333a49459585C7e087"

const personalAddress = "0xbbE4733d85bc2b90682147779DA49caB38C0aA1F"

const personalKey = "8ff3ca2d9985c3a52b459e2f6e7822b23e1af845961e22128d5f372fb9aa5f17"

const contractABI = [
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "c",
        "type": "uint256"
      }
    ],
    "name": "increment",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCount",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "c",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]

const contract = new web3.eth.Contract(contractABI, contractAddress);

async function run(){
  const txCount = await web3.eth.getTransactionCount(personalAddress);
  const txObject = {
    from: personalAddress,
    to: contractAddress,
    data: contract.methods.increment(1).encodeABI(),
    nonce: web3.utils.toHex(txCount),
    gas: 2000000
  };

  // 签名交易
  const signedTx = await web3.eth.accounts.signTransaction(txObject, personalKey);

  // 发送交易
  const txResult = await web3.eth.sendSignedTransaction(signedTx.rawTransaction);

  console.log("txResult:", txResult)

  const callResult = await contract.methods.getCount().call()

  console.log("Call result1:", callResult)

  const callResult2 = await web3.eth.call({
    to:contractAddress,
    data:contract.methods.getCount().encodeABI()
  })
  console.log("Call result2:", callResult)
}

run()

