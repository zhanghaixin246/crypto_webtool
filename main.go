package main

import (
	"crypto_webtool/core"
	"crypto_webtool/global"
	"crypto_webtool/handler"
	"crypto_webtool/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Viper
	global.CW_VP = core.Viper()
	// 初始化日志包
	global.CW_LOG = core.Zap()
	/*
		  example:
			curl -X POST http://localhost:8083/cryptogen/generate \
			  -F "file=@/Users/zhang/go/src/fabric-samples/test-network/organizations/cryptogen/crypto-config-org2.yaml" \
			  -H "Content-Type: multipart/form-data"
	*/
	//api 模式
	if global.CW_CONFIG.System.ApiMode {
		//中间件来代替gin框架默认的Logger()和Recovery()
		r := gin.New()
		r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
		//r := gin.Default()
		{
			r.GET("/", handler.HealthCheck)

			//cryptogen generate --config=/Users/zhang/go/src/fabric-samples/test-network/organizations/cryptogen/crypto-config-org2.yaml --output="organizations"
			r.POST("/cryptogen/generate", handler.GenerateApi)
			r.POST("/cryptogen/multi-generate", handler.MultiGenerateApi)

			//configtxlator proto_decode --input fabric_block.pb --type common.Block
			r.POST("/protolator/decode", handler.ProtoDecode)

			//configtxgen -profile TwoOrgsOrdererGenesis -channelID system-channel -outputBlock ./system-genesis-block/genesis.block
			//configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/mychannel.tx -channelID mychannel
			//configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID mychannel -asOrg Org1MSP
			r.POST("/configtxgen/profile", handler.ProfileGenerate)
			r.Any("/test", handler.Test)
		}

		port := global.CW_CONFIG.System.Port
		//port := ":8083"
		r.Run(port)
		return
	}

	panic("Fatal error check config file")
}
