// contracts/ExampleToken.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

// https://openzeppelin.com/contracts/
contract UsdToken is ERC20, Ownable {
    using SafeMath for uint256;

    constructor ()
    ERC20("UsdToken", "USDK")
    {}

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply. This is done after dollars are deposited into the reserve.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     */
    function mint(address account, uint256 amount) public onlyOwner {
        // TODO: fill out this function for exercise 1!
        // Hint: Smart Contracts can call functions on contracts they inherit
    }

    /**
     * @dev Destroys `amount` tokens from `account`, reducing the
     * total supply. This is done before dollars are withdrawn from the reserve.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function burn(address account, uint256 amount) public onlyOwner {
        return _burn(account, amount);
    }

}
