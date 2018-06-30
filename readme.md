

## sign in btc block

### get testing coin
> :star:[Bitcoin testnet3 faucet](https://testnet.coinfaucet.eu/en/)
> 

https://api.blockcypher.com/v1/btc/test3/addrs/{address}


### build transaction
:heart[balance unspent here](https://api.blockcypher.com/v1/btc/test3/addrs/msXzK5c57uo6aP5pHoaBnGzSbKcmmh9rYe/full?limit=50?unspentOnly=true&includeScript=true)

> broadcast in [here]()


### go get
```
GOPATH : D:\go\gopath  
```

`gopm` (using git for windows)
```
go get -u github.com/gpmgo/gopm

gopm get github.com/btcsuite/btcd/btcec
```

## tools
1. [tx decode](https://btc.com/tools/tx/decode)
2. [BCC](https://www.blocktrail.com/tBCC/address/)
