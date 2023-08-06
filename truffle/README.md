## 说明
本项目是truffle的使用模版。
truffle网站：https://trufflesuite.com/docs/truffle/

## 环境
```
npm install web3
npm install @truffle/hdwallet-provider
```
## 合约代码
contracts/simple.sol

## 执行步骤
1. 配置truffle-config.js
2. 执行truffle compile。生成在build目录中
3. 编写部署脚本，migrations/2_deploy_contracts.js。执行truffle migrate --network xxx

## 使用truffle console调用合约
```
truffle console
const SimpleStorage = artifacts.require('Simple');
const simpleStorageInstance = await SimpleStorage.deployed();
await simpleStorageInstance.set(35)
await simpleStorageInstance.get()
```

## 使用web3调用合约
```
cd call
node rpc.js
```
