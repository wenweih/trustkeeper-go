package configure

import (
	"fmt"
	"strings"
	"trustkeeper-go/library/vault"
	"github.com/btcsuite/btcd/rpcclient"
)

type Conf struct {
	DBInfo				string
	ConsulAddress	string
	Redis         string
	BTCconnCfg    *rpcclient.ConnConfig
	ETHRPC        string
	MQ            string
}

func New() (*Conf, error) {
	vc, path, err := vault.NewVault()
	if err != nil {
		return nil, fmt.Errorf("fail to connect vault" + err.Error())
	}
	// ListSecret
	data, err := vc.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("vaule read error" + err.Error())
	}

	host := strings.Join([]string{"host", data.Data["host"].(string)}, "=")
	port := strings.Join([]string{"port", data.Data["port"].(string)}, "=")
	user := strings.Join([]string{"user", data.Data["username"].(string)}, "=")
	password := strings.Join([]string{"password", data.Data["password"].(string)}, "=")
	dbname := strings.Join([]string{"dbname", data.Data["dbname"].(string)}, "=")
	sslmode := strings.Join([]string{"sslmode", data.Data["sslmode"].(string)}, "=")
	dbInfo := strings.Join([]string{host, port, user, dbname, password, sslmode}, " ")
	consulAddr := data.Data["consuladdr"].(string)
	redis := data.Data["redis"].(string)

	bitcoinHost := data.Data["btchost"].(string)
	bitcoinUsr := data.Data["btcuse"].(string)
	bitcoinPass := data.Data["btcpass"].(string)
	mq := data.Data["mq"].(string)

	eth_rpc := data.Data["eth_rpc"].(string)

	return &Conf{
		DBInfo:			dbInfo,
		ConsulAddress: consulAddr,
		Redis: redis,
		MQ: mq,
		ETHRPC: eth_rpc,
		BTCconnCfg: &rpcclient.ConnConfig{
			Host: bitcoinHost,
			User: bitcoinUsr,
			Pass: bitcoinPass,
			HTTPPostMode: true,
			DisableTLS: true,
		},
	}, nil
}
