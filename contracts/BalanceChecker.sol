// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title BalanceChecker
 * @dev 批量查询ETH和ERC20代币余额的合约
 */
contract BalanceChecker {
    
    /**
     * @dev 查询单个地址的ETH余额
     * @param target 目标地址
     * @return 返回ETH余额（以wei为单位）
     */
    function getETHBalance(address target) public view returns (uint256) {
        return target.balance;
    }
    
    /**
     * @dev 批量查询多个地址的ETH余额
     * @param targets 目标地址数组
     * @return 返回ETH余额数组（以wei为单位）
     */
    function getETHBalances(address[] calldata targets) public view returns (uint256[] memory) {
        uint256[] memory balances = new uint256[](targets.length);
        
        for (uint256 i = 0; i < targets.length; i++) {
            balances[i] = targets[i].balance;
        }
        
        return balances;
    }
    
    /**
     * @dev 查询单个地址的ERC20代币余额
     * @param token 代币合约地址
     * @param target 目标地址
     * @return 返回代币余额
     */
    function getERC20Balance(address token, address target) public view returns (uint256) {
        return IERC20(token).balanceOf(target);
    }
    
    /**
     * @dev 批量查询多个地址的ERC20代币余额
     * @param token 代币合约地址
     * @param targets 目标地址数组
     * @return 返回代币余额数组
     */
    function getERC20Balances(address token, address[] calldata targets) public view returns (uint256[] memory) {
        uint256[] memory balances = new uint256[](targets.length);
        
        for (uint256 i = 0; i < targets.length; i++) {
            balances[i] = IERC20(token).balanceOf(targets[i]);
        }
        
        return balances;
    }
    
    /**
     * @dev 批量查询多个代币合约的多个地址余额
     * @param tokens 代币合约地址数组
     * @param targets 目标地址数组
     * @return 返回二维数组，第一维是代币，第二维是地址
     */
    function getMultipleERC20Balances(
        address[] calldata tokens, 
        address[] calldata targets
    ) public view returns (uint256[][] memory) {
        uint256[][] memory balances = new uint256[][](tokens.length);
        
        for (uint256 i = 0; i < tokens.length; i++) {
            balances[i] = new uint256[](targets.length);
            for (uint256 j = 0; j < targets.length; j++) {
                balances[i][j] = IERC20(tokens[i]).balanceOf(targets[j]);
            }
        }
        
        return balances;
    }
    
    /**
     * @dev 查询地址的ETH余额和指定ERC20代币余额
     * @param target 目标地址
     * @param tokens 代币合约地址数组
     * @return ethBalance ETH余额
     * @return tokenBalances 代币余额数组
     */
    function getAddressBalances(
        address target, 
        address[] calldata tokens
    ) public view returns (uint256 ethBalance, uint256[] memory tokenBalances) {
        ethBalance = target.balance;
        tokenBalances = new uint256[](tokens.length);
        
        for (uint256 i = 0; i < tokens.length; i++) {
            tokenBalances[i] = IERC20(tokens[i]).balanceOf(target);
        }
    }
    
    /**
     * @dev 批量查询多个地址的ETH余额和指定ERC20代币余额
     * @param targets 目标地址数组
     * @param tokens 代币合约地址数组
     * @return ethBalances ETH余额数组
     * @return tokenBalances 代币余额二维数组
     */
    function getMultipleAddressBalances(
        address[] calldata targets, 
        address[] calldata tokens
    ) public view returns (
        uint256[] memory ethBalances, 
        uint256[][] memory tokenBalances
    ) {
        ethBalances = new uint256[](targets.length);
        tokenBalances = new uint256[][](targets.length);
        
        for (uint256 i = 0; i < targets.length; i++) {
            ethBalances[i] = targets[i].balance;
            tokenBalances[i] = new uint256[](tokens.length);
            
            for (uint256 j = 0; j < tokens.length; j++) {
                tokenBalances[i][j] = IERC20(tokens[j]).balanceOf(targets[i]);
            }
        }
    }
}
