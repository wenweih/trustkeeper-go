package main

import (
  "fmt"
  "context"
  "net/http"
  "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
  common "trustkeeper-go/library/util"
)

func btcBestBlockNotifyHandle(c *gin.Context) {
  strHash := c.Query("hash")
  blockHash, err := chainhash.NewHashFromStr(strHash)
  if err != nil {
    logger.Log("NewHashFromStr error", err.Error())
    common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("NewHashFromStr %s", err))
    return
  }
  block, err := svc.BitcoincoreBlock(context.Background(), blockHash)
  if err != nil {
    logger.Log("query BitcoincoreBlock error", err.Error())
    common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("BitcoincoreBlock %s", err))
    return
  }
  body, err := json.Marshal(block)
  if err != nil {
    logger.Log("json Marshal raw block error", err.Error())
    common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("json Marshal %s", err))
    return
  }
  if err := svc.MQPublish(body, "bestblock", "direct", "bitcoincore", "bitcoincore_best_block_queue");
    err != nil {
      logger.Log("fail to publish bitcoin block msg to mq", err.Error())
      common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("MQPublish %s", err))
      return
  }
  c.JSON(http.StatusOK, gin.H {
    "status": http.StatusOK,
    "hash": block.Hash,
  })
}
