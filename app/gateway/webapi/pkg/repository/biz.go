package repository

// SimpleToken token resp
type SimpleToken struct {
  TokenID  string  `json:"tokenid"`
  Symbol   string  `json:"symbol"`
  Status   bool    `json:"status"`
}

// GroupAsset tokens correspond with chain resp
type GroupAsset struct {
  Chainid  string  `json:"chainid"`
	Name     string  `json:"name"`
  Coin     string  `json:"desc"`
  Status   bool    `json:"status"`
  SimpleTokens []*SimpleToken `json:"simpletokens"`
}
