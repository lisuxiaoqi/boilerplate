require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: ["d8611869c1cf0548d412322d5a946b1fa5303d80a9ce48ff0a7b697d1c7f3cd6"],
    },
    okbtestokx: {
      url: "https://okbtestrpc.okx.com",
      accounts: ["13977ec5c2fd6f2fa064a54919ea6f1b2efaf2982082e14560cf7536bd1ad670"],
    },
    okbtest: {
      url: "https://okbtestrpc.okbchain.org",
      accounts: ["13977ec5c2fd6f2fa064a54919ea6f1b2efaf2982082e14560cf7536bd1ad670"],
    },
  },
  solidity: "0.8.18",
};
