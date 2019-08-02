package repository

import(
  "encoding/json"
)

// OmniProperty omnicore property
type OmniProperty struct {
  Propertyid    int64    `json:"propertyid"`
  Name          string   `json:"name"`
  Category      string   `json:"category"`
  Subcategory   string   `json:"subcategory"`
  URL           string   `json:"url"`
  Divisible     bool     `json:"divisible"`
  Issuer        string   `json:"issuer"`
  Creationtxid  string   `json:"creationtxid"`
  Fixedissuance bool     `json:"fixedissuance"`
  Managedissuance bool   `json:"managedissuance"`
  Freezingenabled bool   `json:"freezingenabled"`
  Totaltokens string     `json:"totaltokens"`
}

func (repo *repo ) QueryOmniProperty(propertyid int64) (*OmniProperty, error) {
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
