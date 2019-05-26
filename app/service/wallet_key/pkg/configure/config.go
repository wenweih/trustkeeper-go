package configure

import (
	"fmt"
	"trustkeeper-go/library/vault"
)

type Conf struct {
	EtcdServer	string
	AccountInstance string
}

func New() (*Conf, error) {
	vc, err := vault.NewVault()
	if err != nil {
		return nil, fmt.Errorf("fail to connect vault" + err.Error())
	}
	// ListSecret
	data, err := vc.Logical().Read("kv1/db_trustkeeper_account")
	if err != nil {
		return nil, fmt.Errorf("vaule read error" + err.Error())
	}

	etcdServer := data.Data["etcdserver"].(string)
	accountInstance := data.Data["accountinstance"].(string)
	return &Conf{
		EtcdServer: etcdServer,
		AccountInstance: accountInstance,
	}, nil
}
