package common

import (
  "net"
  "github.com/gin-gonic/gin"
)

const (
  // AccountSrv account service register in etcd
  AccountSrv = "/services/account/"
  // DashboardSrv dashboard service register in etcd
  DashboardSrv = "/services/dashboard/"

  TxSrv = "/services/tx/"

  // WalletKeySrv wallet_key register in etcd
  WalletKeySrv = "/services/walletkey/"

  WalletManagementSrv = "/services/walletmanagement/"

  LevelDBDir = "tk_w"

  // jobs name
  SignUpJobs = "signup"
  WalletMnemonicJob = "wn"


  BTC = 0
  LTC = 2
  ETH = 60
  ETC = 61
  BCH = 145
  EOS = 194
  BSV = 236
  BNB = 714
  XLM = 148
  ADA = 1815
  XMR = 128
  ATOM = 118
  XRP = 144
)

// LocalIP returns the non loopback IP of the host
// https://stackoverflow.com/a/31551220/6998584
func LocalIP() string {
  addrs, err := net.InterfaceAddrs()
  if err != nil {
    return ""
  }
  for _, add := range addrs {
    if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        return ipnet.IP.String()
      }
    }
  }
  return ""
}


// GinRespException bad response util
func GinRespException(c *gin.Context, code int, err error) {
  c.AbortWithStatusJSON(code, &JSONAbortMsg{
    Code: code,
    Msg: err.Error(),
  })
}

// JSONAbortMsg about json
type JSONAbortMsg struct {
  Code  int `json:"code"`
  Msg   string `json:"msg"`
}
