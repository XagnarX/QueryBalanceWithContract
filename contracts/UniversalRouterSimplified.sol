// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.17;

contract UniversalRouterSimplified {
    
    event ExecutionFailed(uint256 commandIndex, bytes message);
    
    modifier checkDeadline(uint256 deadline) {
        if (block.timestamp > deadline) revert("TransactionDeadlinePassed");
        _;
    }

    /// @notice Execute commands with inputs and deadline
    function execute(bytes calldata commands, bytes[] calldata inputs, uint256 deadline)
        external
        payable
        checkDeadline(deadline)
    {
        execute(commands, inputs);
    }

    /// @notice Execute commands with inputs
    function execute(bytes calldata commands, bytes[] calldata inputs) 
        public 
        payable 
    {
        uint256 numCommands = commands.length;
        if (inputs.length != numCommands) revert("LengthMismatch");

        // loop through all given commands, execute them and pass along outputs as defined
        for (uint256 commandIndex = 0; commandIndex < numCommands;) {
            bytes1 command = commands[commandIndex];
            bytes calldata input = inputs[commandIndex];

            // For simplified version, we'll just emit an event
            // In real implementation, this would dispatch to specific handlers
            emit ExecutionFailed(commandIndex, input);

            unchecked {
                commandIndex++;
            }
        }

        uint256 balance = address(this).balance;
        if ((balance > 0) && (msg.sender != address(this))) {
            // Transfer remaining ETH back to sender
            (bool success, ) = msg.sender.call{value: balance}("");
            require(success, "ETH transfer failed");
        }
    }

    /// @notice To receive ETH from WETH and NFT protocols
    receive() external payable {}
}
