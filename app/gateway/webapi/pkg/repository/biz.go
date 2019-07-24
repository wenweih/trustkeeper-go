package repository

// SimpleAsset token resp
type SimpleAsset struct {
  AssetID  string  `json:"AssetID"`
  Symbol   string  `json:"Symbol"`
  Status   bool    `json:"Status"`
}

// GroupAsset tokens correspond with chain resp
type GroupAsset struct {
  ChainID  string  `json:"ChainID"`
	Name     string  `json:"Name"`
  Coin     string  `json:"Coin"`
  Status   bool    `json:"Status"`
  SimpleAssets []*SimpleAsset `json:"SimpleAssets"`
}
