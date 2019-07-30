package main

import(
  "strconv"
  "errors"
  "strings"
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
    e := errors.New(strings.Join([]string{"FailToGetOriginBlock ", err.Error(), " Height ", number.String()}, ""))
    return big.NewInt(head.Number.Int64() - 1), e
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
      e := errors.New(strings.Join([]string{
        "Get block error, height:",
        strconv.FormatInt(blockNumber, 10)}, ""))
			return big.NewInt(blockNumber), e
		}
		body, err := json.Marshal(block)
    if err != nil {
      e := errors.New(strings.Join([]string{"json Marshal raw ethereum block error", err.Error()}, ""))
			return big.NewInt(blockNumber), e
    }
		if err := svc.MQPublish(
      body,
      "bestblock",
      "fanout",
      "ethereum",
      "ethereum_best_block_queue");
      err != nil {
        e := errors.New(strings.Join([]string{"EtherumPublishError", err.Error()}, ""))
        return big.NewInt(blockNumber), e
      }
    }
  return orderHeight, nil
}
