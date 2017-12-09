var Migrations = artifacts.require("./Migrations.sol");
var AIDataAssetRegister = artifacts.require("./DataAssetRegister/AIDataAssetRegister.sol");
var BTOToken = artifacts.require("./BTOToken/BTOToken.sol");
var DataExchangeManager = artifacts.require("./DataExchangeManager/DataExchangeManager.sol");
var DataRequirementManager = artifacts.require("./DataRequirement/DataRequirementManager.sol");
var DataStore = artifacts.require("./DataStore/DataStore.sol");
var LibDataRequirement = artifacts.require("./DataRequirement/LibDataRequirement.sol");
var TokenManager = artifacts.require("./TokenManager/TokenManager.sol");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(AIDataAssetRegister);
  //deployer.deploy(BTOToken);
  deployer.deploy(DataExchangeManager);
  deployer.deploy(LibDataRequirement);
  deployer.link(LibDataRequirement, DataRequirementManager);
  deployer.deploy(DataRequirementManager);
  //deployer.deploy(DataStore);
  deployer.deploy(TokenManager);
};
