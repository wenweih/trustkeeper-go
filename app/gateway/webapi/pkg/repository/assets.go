package repository

// SimpleAsset token resp
type SimpleAsset struct {
  AssetID  string  `json:"AssetID"`
  Symbol   string  `json:"Symbol"`
  Status   bool    `json:"Status"`
  Identify string  `json:"Identify"`
  Decimal  uint64  `json:"Decimal"`
}

// ERC20Token Ethereum erc20 token info
type ERC20Token struct {
  Symbol   string `json:"Symbol"`
  Name     string `json:"Name"`
  Decimals uint32 `json:"Decimals"`
}

// GroupAsset tokens correspond with chain resp
type GroupAsset struct {
  ChainID  string  `json:"ChainID"`
	Name     string  `json:"Name"`
  Coin     string  `json:"Coin"`
  Status   bool    `json:"Status"`
  Decimal  uint64  `json:"Decimal"`
  SimpleAssets []*SimpleAsset `json:"SimpleAssets"`
}
