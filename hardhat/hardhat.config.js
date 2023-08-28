require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: ["08a59286ea6759517e9cd2f01faf9625f0c7502d5401656a2ef5f3121c977f82"],
    },
    okbtestokx: {
      url: "https://okbtestrpc.okx.com",
      accounts: ["08a59286ea6759517e9cd2f01faf9625f0c7502d5401656a2ef5f3121c977f82"],
    },
    okbtest: {
      url: "https://okbtestrpc.okbchain.org",
      accounts: ["08a59286ea6759517e9cd2f01faf9625f0c7502d5401656a2ef5f3121c977f82"],
    },
  },
  solidity: "0.8.18",
};
