package main

import (
  "fmt"
  "context"
  "net/http"
  // "encoding/json"
  "github.com/gin-gonic/gin"
  "github.com/btcsuite/btcd/chaincfg/chainhash"
  common "trustkeeper-go/library/util"
)

func btcBestBlockNotifyHandle(c *gin.Context) {
  strHash := c.Query("hash")
  blockHash, err := chainhash.NewHashFromStr(strHash)
  if err != nil {
    fmt.Println(err.Error())
    common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("NewHashFromStr %s", err))
    return
  }
  block, err := svc.BitcoincoreBlock(context.Background(), blockHash)
  if err != nil {
    fmt.Println(err.Error())
    common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("NewHashFromStr %s", err))
    return
  }
  // rawBlock, err := svc.BitcoincoreBlock(context.Background(), strHash)
  // // rawBlock, err := btcClient.Client.GetBlockVerboseTx(blockHash)
  // if err != nil {
  //   logger.Log(err.Error())
  //   common.GinRespException(c, http.StatusInternalServerError, fmt.Errorf("GetBlockVerboseTxM %s", err))
  //   return
  // }
  // body, err := json.Marshal(rawBlock)
  // if err != nil {
  //   logger.Log("json Marshal raw block error", err.Error())
  // }
  // messageClient.Publish(body, "bestblock", "fanout", "bitcoincore", "bitcoincore_best_block_queue")
  c.JSON(http.StatusOK, gin.H {
    "status": http.StatusOK,
    "hash": block.Hash,
  })
}
