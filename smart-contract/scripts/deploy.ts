import { network } from "hardhat";

async function main() {
  const { viem } = await network.connect({
    network: "hardhatOp",
    chainType: "op",
  });
  //   const publicClient = await viem.();

  //   const Transfer = await hre. .getContractFactory("Transfer");
  //   const transfer = await Transfer.deploy();
  //   await transfer.deployed();
  //   console.log("Transfer deployed to:", transfer.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
