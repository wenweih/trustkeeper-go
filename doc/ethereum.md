- [command line options](https://github.com/ethereum/go-ethereum/wiki/Command-Line-Options)
-
- [Running ethereum in Docker](https://github.com/ethereum/go-ethereum/wiki/Running-in-Docker)
- [Ethereum 私有链和 web3.js 使用](https://huangwenwei.com/blogs/ethereum-private-chain-and-web3js)
- [Ethereum private chain resets back to block 0 when restarted](https://ethereum.stackexchange.com/questions/39922/ethereum-private-chain-resets-back-to-block-0-when-restarted)「Every time you terminate geth, up to 128 blocks can be lost, with default settings.」
- [Ethereum Development with Go](https://goethereumbook.org/en/)
- [Native DApps: Go bindings to Ethereum contracts](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts)

#### Go bindings to Ethereum contracts
```
cd {chains_query service repository directly}
wget https://github.com/ConsenSys/Tokens/raw/fdf687c69d998266a95f15216b1955a4965a0a6d/contracts/eip20/EIP20.sol

wget https://github.com/ConsenSys/Tokens/raw/fdf687c69d998266a95f15216b1955a4965a0a6d/contracts/eip20/EIP20Interface.sol

solc --abi EIP20.sol -o ./

abigen --abi=EIP20.abi --pkg=token --out=erc20.go
```
