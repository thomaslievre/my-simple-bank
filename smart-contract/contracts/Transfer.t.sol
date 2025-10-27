// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "./Transfer.sol";

contract TransferTest is Test {
    Transfer transfer;
    address payable recipient = payable(address(0xBEEF));
    address sender = address(0xCAFE);

    function setUp() public {
        transfer = new Transfer();
        vm.deal(sender, 10 ether); // Give sender some ether
    }

    function testSendTransfersFundsAndEmitsEvent() public {
        uint256 amount = 1 ether;
        uint256 initialRecipientBalance = recipient.balance;

        vm.prank(sender);
        vm.expectEmit(true, true, false, true);
        emit Transfer.FundsTransferred(sender, recipient, amount);

        transfer.send{value: amount}(recipient);

        assertEq(
            recipient.balance,
            initialRecipientBalance + amount,
            "Recipient balance should increase"
        );
    }

    function testSendZeroAmountReverts() public {
        vm.prank(sender);
        vm.expectRevert("Amount must be > 0");
        transfer.send{value: 0}(recipient);
    }
}
