const usdk = artifacts.require("UsdToken");

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(usdk);
};
