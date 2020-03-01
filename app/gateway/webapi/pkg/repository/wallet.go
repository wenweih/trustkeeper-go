package repository

type ChainWithWallets struct {
	ChainName string    `json:"ChainName"`
	TotalSize int32     `json:"TotalSize"`
	Wallets   []*Wallet `json:"Wallets"`
}

type Wallet struct {
	ID        string `json:"ID,omitempty"`
	Address   string `json:"Address"`
	Status    bool   `json:"Status"`
	ChainName string `json:"ChainName,omitempty"`
}
