package api

import (
	"chainup.com/node-exchange/chain"
	"chainup.com/node-exchange/common/logger"
	"chainup.com/node-exchange/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecipet(c *gin.Context) {
	var resp *entity.Response

	defer func() {
		c.JSON(http.StatusOK, resp)
	}()

	var reqData entity.TxRequestBody
	if err := c.ShouldBindJSON(&reqData); err != nil {
		logger.Error("ShouldBindJSON occur exception: {}", err)
		resp = entity.GetErrResponse(entity.JsonDecodeErr, "")
		return
	}

	sync, err := chain.InitSyncer("https://smart.zeniq.network:9545", 383414847825)
	if err != nil {
		resp = entity.GetErrResponse(entity.InitSyncErr, "")
		return
	}
	// 访问rpc
	switch reqData.Method {
	case "eth_getTransactionReceipt":
		// 获取交易收据
		c, err := sync.TransactionReceipt(fmt.Sprintf("%v", reqData.Params[0]))
		if err != nil {
			logger.Error("TransactionReceipt occur exception: {}", err)
			resp = entity.GetErrResponse(entity.GetTransactionErr, "")
			return
		}
		resp = entity.GetResponse(entity.OK, "", c)
		return
	case "eth_getTransactionByHash":
		// 获取交易详情
		c, _, err := sync.TransactionByHash(fmt.Sprintf("%v", reqData.Params[0]))
		if err != nil {
			logger.Error("GetTransactionByHash occur exception: {}", err)
			resp = entity.GetErrResponse(entity.GetTransactionErr, "")
			return
		}
		resp = entity.GetResponse(entity.OK, "", c)
		return
	case "eth_getBlockByHash":
		// 获取区块详情
		c, err := sync.GetBlockByHash(fmt.Sprintf("%v", reqData.Params[0]))
		if err != nil {
			logger.Error("GetBlockByHash occur exception: {}", err)
			resp = entity.GetErrResponse(entity.GetTransactionErr, "")
			return
		}
		resp = entity.GetResponse(entity.OK, "", c)
		return
	case "eth_getBlockByNumber":
		c, err := sync.GetBlockByHash(fmt.Sprintf("%v", reqData.Params[0]))
		if err != nil {
			logger.Error("GetBlockByNumber occur exception: {}", err)
			resp = entity.GetErrResponse(entity.GetTransactionErr, "")
			return
		}
		resp = entity.GetResponse(entity.OK, "", c)
		return
	case "eth_call":
		// 调用合约
		c, err := sync.EthCall(
			fmt.Sprintf("%v", reqData.Params[0]),
			fmt.Sprintf("%v", reqData.Params[1]),
			fmt.Sprintf("%v", reqData.Params[2]),
		)
		if err != nil {
			logger.Error("EthCall occur exception: {}", err)
			resp = entity.GetErrResponse(entity.GetTransactionErr, "")
			return
		}
		resp = entity.GetResponse(entity.OK, "", c)
		return
	default:
		resp = entity.GetErrResponse(entity.NotSupportTransType, "")
		return
	}
}
