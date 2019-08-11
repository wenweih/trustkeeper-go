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
erc20 token address 0xf0680d66aac362b1e42e21d3098ad61e92c6f43f

```
// transfer eth
eth.sendTransaction({from: eth.coinbase, to: "0xf5AaeE49BF40a5fc0f0373ea036356cc8a240263", value: web3.toWei(1, "ether")})

var abi = [{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_value","type":"uint256"}],"name":"burn","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_value","type":"uint256"}],"name":"burnFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"},{"name":"_extraData","type":"bytes"}],"name":"approveAndCall","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"initialSupply","type":"uint256"},{"name":"tokenName","type":"string"},{"name":"tokenSymbol","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Burn","type":"event"}]

var contract_address = '0xf0680d66aac362b1e42e21d3098ad61e92c6f43f'

var contract_object = eth.contract(abi).at(contract_address)

var coinbase_token = contract_object.balanceOf(eth.coinbase).div(1e18)

contract_object.transfer.sendTransaction('0x6F485F869706a658399A2d5839dC01acc80dD315', 10000000000000000000, {from: web3.eth.accounts[0]})
```
