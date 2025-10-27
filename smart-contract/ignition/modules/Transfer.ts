import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("TransferModule", (m) => {
  const transfer = m.contract("Transfer");

  m.call(transfer, "getBalance", []);

  return { transfer };
});
