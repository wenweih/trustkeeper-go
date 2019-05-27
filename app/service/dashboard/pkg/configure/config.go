package configure

import (
	"fmt"
	"strings"
	"trustkeeper-go/library/vault"
)

type Conf struct {
	DBInfo			string
	EtcdServer	string
	Instance string
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

	host := strings.Join([]string{"host", data.Data["host"].(string)}, "=")
	port := strings.Join([]string{"port", data.Data["port"].(string)}, "=")
	user := strings.Join([]string{"user", data.Data["username"].(string)}, "=")
	password := strings.Join([]string{"password", data.Data["password"].(string)}, "=")
	dbname := strings.Join([]string{"dbname", data.Data["dbname"].(string)}, "=")
	sslmode := strings.Join([]string{"sslmode", data.Data["sslmode"].(string)}, "=")
	dbInfo := strings.Join([]string{host, port, user, dbname, password, sslmode}, " ")
	etcdServer := data.Data["etcdserver"].(string)
	instance := data.Data["accountinstance"].(string)
	return &Conf{
		DBInfo:			dbInfo,
		EtcdServer: etcdServer,
		Instance: instance,
	}, nil
}
