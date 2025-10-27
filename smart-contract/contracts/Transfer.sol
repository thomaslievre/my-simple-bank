// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Transfer {
    event FundsTransferred(
        address indexed from,
        address indexed to,
        uint256 amount
    );

    function send(address payable to) external payable {
        require(msg.value > 0, "Amount must be > 0");
        to.transfer(msg.value);
        emit FundsTransferred(msg.sender, to, msg.value);
    }

    // Permet de consulter le solde du contrat
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }

    // Fonction pour recevoir de l'ether
    receive() external payable {}
}
