package repository

// SimpleAsset token resp
type SimpleAsset struct {
	AssetID  string `json:"AssetID"`
	Symbol   string `json:"Symbol"`
	Status   bool   `json:"Status"`
	Identify string `json:"Identify"`
	Decimal  uint64 `json:"Decimal"`
}

type Wallet struct {
	Address   string `json:"Address"`
	Status    bool   `json:"Status"`
	ChainName string `json:"ChainName"`
}
