package configure

import (
	"fmt"
	"strings"
	"trustkeeper-go/library/vault"
)

type Conf struct {
	Redis         string
	DBInfo        string
	JWTKey        string
	ConsulAddress string
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
	jwtkey := data.Data["jwtkey"].(string)
	redis := data.Data["redis"].(string)
	consuladdr := data.Data["consuladdr"].(string)
	return &Conf{
		Redis:         redis,
		DBInfo:        dbInfo,
		JWTKey:        jwtkey,
		ConsulAddress: consuladdr,
	}, nil
}
