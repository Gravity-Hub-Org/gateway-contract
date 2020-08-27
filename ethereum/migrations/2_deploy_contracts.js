const fs = require("fs");
const Nebula = artifacts.require("./Nebula/Nebula.sol");
const Gravity = artifacts.require("./Gravity/Gravity.sol");
const IBPort = artifacts.require("./IBPort/IBPort.sol");
const Token = artifacts.require("./IBPort/Token.sol");
const Queue = artifacts.require("./libs/QueueLib");

module.exports = async function(deployer, network, accounts) {
  await deployer.deploy(Queue);
  await deployer.link(Queue, Nebula);
  await deployer.link(Queue, Gravity);

  const gravity = await deployer.deploy(Gravity, accounts, 1);
  const nebula = await deployer.deploy(Nebula, accounts.slice(0, 5), gravity.address, 3);
  const token = await deployer.deploy(Token, "TEST", "TST", 8);

  const ibport = await deployer.deploy(IBPort, nebula.address, token.address);

  await token.addMinter(ibport.address);

  // const sub = await nebula.subscribe(ibport.address, 0, "100000000000000")
  const sub = await nebula.subscribe(ibport.address, 0, "0")
  await web3.eth.sendTransaction({ from: accounts[0], to: sub.address, value:  "10000000000000000000" });
  fs.writeFileSync("nebula.json", JSON.stringify({
    address: nebula.address,
    abi: JSON.stringify(nebula.abi),
    tokenAbi: JSON.stringify(token.abi),
    ibportAddress: ibport.address,
    tokenAddress: token.address,
    subscriptionId: sub.receipt.logs[0].args.id,
  }));
};