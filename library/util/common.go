package common

import (
  "net"
  "os"
  "strconv"
  "syscall"
  "os/signal"
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

  ChainsQuerySrv = "/services/chainsquery/"

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


// Handles Ctrl+C or most other means of "controlled" shutdown gracefully. Invokes the supplied func before exiting.
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}

func Hex2int(hexStr string) int64 {
  // base 16 for hexadecimal
  result, _ := strconv.ParseUint(hexStr, 16, 64)
  return int64(result)
}
