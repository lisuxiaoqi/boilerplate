const Web3 = require('web3');
const URL='http://localhost:8545';
const CHAINID = 66;
const FROM = "0x8d63Bd4B972794b67aA83CDaf09dB5655b4e00CC"
const PRIVATE_KEY = "08a59286ea6759517e9cd2f01faf9625f0c7502d5401656a2ef5f3121c977f82"
const TO = "0xe1A3cD33b25487fdcf7ac5d347Dbd2A2Fe18bd07"

// 节点
const web3 = new Web3(URL);

async function sendRawTransactionAsync(signedTx) {
    return new Promise((resolve, reject) => {
        web3.currentProvider.send(
            {
                jsonrpc: '2.0',
                method: 'eth_sendRawTransaction',
                params: [signedTx.rawTransaction],
                id: 253,
            },
            (err, result) => {
                if (err) {
                    console.log("Error", err);
                    reject(err);
                } else {
                    //console.log("OK", result);
                    resolve(result.result);
                }
            }
        );
    });
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function getTxReceipt(txHash, delayMs) {
    while (true) {
        const result = await web3.eth.getTransactionReceipt(txHash);
        if(result != null){
            console.log("TransactionReceipt", result);
            break;
        }
        await sleep(delayMs);
    }
}

async function transfer(){
    let b_from =  web3.utils.fromWei(await web3.eth.getBalance(FROM));
    let b_to =  web3.utils.fromWei(await web3.eth.getBalance(TO));
    console.log("before transfer, %s:%s", FROM, b_from);
    console.log("before transfer, %s:%s", TO, b_to);

    const nonce = await web3.eth.getTransactionCount(FROM, 'latest'); // nonce starts counting from 0

    const transaction = {
        'to': TO, // faucet address to return eth
        'value': 100000000000000000000, // 100 ETH
        'gas': 30000,
        'nonce': nonce,
        'maxPriorityFeePerGas':100,
        'maxFeePerGas':1000,
        'chainId': CHAINID
        // optional data field to send message or execute smart contract
    };

    const signedTx = await web3.eth.accounts.signTransaction(transaction, PRIVATE_KEY);
    //console.log(signedTx.rawTransaction);

    //It's OK to use sendSignedTransaction here
    //  Be careful sendSignedTransaction will dead loop if the txHash is invalid,
    //  for it always request for the receipt.
    const txHash = await sendRawTransactionAsync(signedTx);
    console.log("TransactionHash:", txHash);

    if(txHash == '0x0000000000000000000000000000000000000000000000000000000000000000'){
        console.log("Send Tx failed!")
        return;
    }

    await getTxReceipt(txHash, 2);
    
    b_from =  web3.utils.fromWei(await web3.eth.getBalance(FROM));
    b_to =  web3.utils.fromWei(await web3.eth.getBalance(TO));
    console.log("after transfer, %s:%s", FROM, b_from);
    console.log("after transfer, %s:%s", TO, b_to);
}

async function main(){
    await transfer();
}

main();
