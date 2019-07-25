package repository

// SimpleAsset token resp
type SimpleAsset struct {
  AssetID  string  `json:"AssetID"`
  Symbol   string  `json:"Symbol"`
  Status   bool    `json:"Status"`
  Identify string  `json:"Identify"`
}
