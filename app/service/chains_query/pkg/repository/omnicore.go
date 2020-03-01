package repository

import (
	"encoding/json"
)

// OmniProperty omnicore property
type OmniProperty struct {
	Propertyid      int64  `json:"Propertyid"`
	Name            string `json:"Name"`
	Category        string `json:"Category"`
	Subcategory     string `json:"Subcategory"`
	URL             string `json:"URL"`
	Divisible       bool   `json:"Divisible"`
	Issuer          string `json:"Issuer"`
	Creationtxid    string `json:"Creationtxid"`
	Fixedissuance   bool   `json:"Fixedissuance"`
	Managedissuance bool   `json:"Managedissuance"`
	Freezingenabled bool   `json:"Freezingenabled"`
	Totaltokens     string `json:"Totaltokens"`
}

func (repo *repo) QueryOmniProperty(propertyid int64) (*OmniProperty, error) {
	propertyID, err := json.Marshal(propertyid)
	if err != nil {
		return nil, err
	}
	var params []json.RawMessage
	{
		params = []json.RawMessage{propertyID}
	}
	resp, err := repo.omniClient.RawRequest("omni_getproperty", params)
	if err != nil {
		return nil, err
	}

	var omniProperty OmniProperty
	if err := json.Unmarshal(resp, &omniProperty); err != nil {
		return nil, err
	}
	return &omniProperty, nil
}
