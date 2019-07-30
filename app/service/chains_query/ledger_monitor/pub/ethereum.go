package main

import(
  "context"
  "math/big"
  "encoding/json"
  "github.com/ethereum/go-ethereum/core/types"
)

func subHandle(orderHeight *big.Int, head *types.Header) (*big.Int, error) {
	ctx := context.Background()
	number := head.Number
  originBlock, err := svc.EthereumBlock(ctx, head.Number)
	if err != nil {
    logger.Log("FailToGetOriginBlock", err.Error, "Height", number.String())
	}

	if orderHeight.Cmp(big.NewInt(0)) == 0 {
		orderHeight = originBlock.Number()
	}
  logger.Log(
    "orderHeight",
    orderHeight.Int64(),
    "SubscriptBlockHeight",
    originBlock.Number().Int64())

	for blockNumber := orderHeight.Int64();
    blockNumber <= originBlock.Number().Int64();
    blockNumber++ {
      orderHeight.Add(orderHeight, big.NewInt(1))
      block, err := svc.EthereumBlock(ctx,  big.NewInt(blockNumber))
  		if err != nil {
        logger.Log("Get block error, height:", blockNumber)
  			continue
  		}
  		body, err := json.Marshal(block)
      if err != nil {
        logger.Log("json Marshal raw ethereum block error", err.Error())
        continue
      }
  		if err := svc.MQPublish(
        body,
        "bestblock",
        "fanout",
        "ethereum",
        "ethereum_best_block_queue");
        err != nil {
          logger.Log("EtherumPublishError", err.Error())
      }
	}
	return orderHeight, nil
}
