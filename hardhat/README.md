## 说明
本项目是hardhat的使用模版。
hardhat网站： https://hardhat.org/hardhat-runner/docs/getting-started#overview

## 合约代码
contracts/sample.sol

## 执行步骤
1. 安装hardhat, 检查package.json中是否有hardhat作为dev dependency，否则的话安装: npm install --save-dev hardhat
1. 配置hardhat.config.js
2. 执行npx hardhat compile。生成在artifacts目录中
3. 编写部署脚本，scripts/deploy.js。执行
```
npx hardhat run scripts/deploy.js --network localhost
```

## 测试
1. 编写调用脚本，scirpts/call.js。填入合约地址，调用合约方法。执行
```
npx hardhat run scripts/call.js --network localhost
```

## 其他
*  查看帮助： npx hardhat help。
*  hardhat是基于task和plugin系统，有点类似于gradle，执行npx hardhat compile时，是在执行compile task。

