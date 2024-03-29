package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"trustkeeper-go/app/service/chains_query/pkg/model"
	"trustkeeper-go/app/service/chains_query/pkg/repository"

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
		originBlock.Number().Int64(),
		"OriginBlockHash", originBlock.Hash().String())

	for blockNumber := orderHeight.Int64(); blockNumber <= originBlock.Number().Int64(); blockNumber++ {
		block, err := svc.EthereumBlock(ctx, big.NewInt(blockNumber))
		if err != nil {
			e := errors.New(strings.Join([]string{
				"Get block error, height:",
				strconv.FormatInt(blockNumber, 10)}, ""))
			return big.NewInt(blockNumber), e
		}
		data, err := model.EncodeETHBlock(*block)
		// data, err := encodeBlock(*block)
		if err != nil {
			fmt.Println(err.Error())
			e := errors.New(strings.Join([]string{"json Marshal raw ethereum block error", err.Error()}, ""))
			return big.NewInt(blockNumber), e
		}
		if err := svc.MQPublish(
			data,
			repository.ExchangeNameForEthereumBestBlock,
			"fanout",
			repository.BindKeyEthereum); err != nil {
			e := errors.New(strings.Join([]string{"EtherumPublishError", err.Error()}, ""))
			return big.NewInt(blockNumber), e
		}
		orderHeight.Add(orderHeight, big.NewInt(1))
	}
	return orderHeight, nil
}
