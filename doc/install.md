## 安装
```shell
git clone
go mod tidy
go mod download
go mod vendor
rm -rf vendor/github.com/btcsuite/btcd
git clone https://github.com/wenweih/btcd_m_backup.git vendor/github.com/btcsuite/btcd
```
