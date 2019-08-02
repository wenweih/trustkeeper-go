```
/opt/bitcoin-0.15.2/bin/bitcoin-cli -datadir=/home/bitcoin/.bitcoin/regtest -rpcuser=btcrpc -rpcpassword=123456 -rpcport=8431 -regtest sendtoaddress n2R3nEPb4RzNVVX4vBFbbze5TeaNVbG2Zt 0.5

/opt/bitcoin-0.15.2/bin/bitcoin-cli -datadir=/home/bitcoin/.bitcoin/regtest -rpcuser=btcrpc -rpcpassword=123456 -rpcport=8431 -regtest generate 1

omnicore-cli -rpcuser=btcrpc -rpcpassword=123456 -rpcport=8432 -regtest omni_sendissuancefixed mhJbuA1g19mw6EYjuV2jFhkzRZUesfgLgQ 2 2 0 "category" "subcategory" "omni_first_token" "huangwenwei.com" "Just do it" 100000000

omnicore-cli -rpcuser=btcrpc -rpcpassword=123456 -rpcport=8432 -regtest omni_gettransaction b7cc479db230e5b425bb2f90800bc6ab9c47052fe63c485f0fca326f95d256c1
```

- [bitcoin.conf](https://github.com/bitcoin/bitcoin/blob/master/share/examples/bitcoin.conf)
- [Accessing Bitcoin's ZeroMQ interface](https://bitcoindev.network/accessing-bitcoins-zeromq-interface/)
- [Block and Transaction Broadcasting with ZeroMQ](https://github.com/bitcoin/bitcoin/blob/master/doc/zmq.md)
- [Mac下为本地回环地址添加别名](https://github.com/hhxsv5/dev-tool/blob/master/LoopbackAlias(Mac%E4%B8%8B%E4%B8%BA%E6%9C%AC%E5%9C%B0%E5%9B%9E%E7%8E%AF%E5%9C%B0%E5%9D%80%E6%B7%BB%E5%8A%A0%E5%88%AB%E5%90%8D)/README.md)
- [omnicore JSON-RPC API](https://github.com/OmniLayer/omnicore/blob/master/src/omnicore/doc/rpc-api.md)
