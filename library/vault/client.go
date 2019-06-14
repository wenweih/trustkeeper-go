package vault

import (
	"log"
	"github.com/caarlos0/env"
	"github.com/hashicorp/vault/api"
)

var (
	cfg config
)

// Config vault server info
type config struct {
	Address	string	`env:"v_address"`
	Token		string	`env:"v_token"`
	Path		string	`env:"v_path"`
}

// Client vault client
type Client struct {
	*api.Client
}

// NewVault new vault client
func NewVault() (c *Client, path string, err error) {
	client, err := api.NewClient(&api.Config{
		Address: cfg.Address,
	})
	if err != nil {
		return nil, "", err
	}

	client.SetToken(cfg.Token)
	return &Client{client},  cfg.Path, nil
}

func init()  {
	cfg = config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln(err.Error())
	}
}
