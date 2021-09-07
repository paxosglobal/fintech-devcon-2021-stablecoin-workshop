# fintech-devcon-2021-stablecoin-workshop

Fintech DevCon 2021 Stablecoin Workshop

# Slides

https://github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/raw/master/slides.pdf

# Setup

- install golang https://golang.org/doc/install
- install docker https://docs.docker.com/get-docker/
- install node https://nodejs.org/en/download/

# This is Exercise 1

## The ERC-20 smart contract

The USDK smart contract we're using complies with the ERC-20 standard. ERC stands for "Ethereum request for comment" and
in this case represents a standard for defining a "token" on the Ethereum blockchain. To learn more about the standard
visit [Ethereum.org's documentation](https://ethereum.org/en/developers/docs/standards/tokens/erc-20/) on the subject.

1. Look for `TODO` in `USDK.sol` to complete the `mint` function

HINT: Smart Contracts can call functions on contracts they inherit. This contract imports
[ERC20.sol](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol),
[Ownable.sol](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol), and
[SafeMath.sol](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/math/SafeMath.sol).
