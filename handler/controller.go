package handler

import (
	"crypto_webtool/fabric/v2.2/bccsp/factory"
	"crypto_webtool/fabric/v2.2/internal-tools/configtxgen/genesisconfig"
	"crypto_webtool/global"
	"crypto_webtool/handler/response"
	"crypto_webtool/tool"
	"crypto_webtool/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	//kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/golang/protobuf/proto"
	_ "github.com/hyperledger/fabric-protos-go/msp"
	_ "github.com/hyperledger/fabric-protos-go/orderer"
	_ "github.com/hyperledger/fabric-protos-go/orderer/etcdraft"
	_ "github.com/hyperledger/fabric-protos-go/peer"
)

var (
//outputDir     *string
//genConfigFile **os.File
)

func HealthCheck(c *gin.Context) {
	response.OkWithMessage("It works!", c)
	return
}

func GenerateApi(c *gin.Context) {
	// 接收http请求，赋值 outputDir 和 genConfigFile
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		response.FailWithMessage(err.Error()+" or file is null", c)
		return
	}
	tmpFile := &global.CW_CONFIG.System.TempFile
	err = c.SaveUploadedFile(file, *tmpFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	outputDir := &global.CW_CONFIG.System.GenerateDir
	//存在则先删除一下
	//if isExit, _ := utils.PathExists(*outputDir); isExit {
	//	err = utils.RemoveDir(*outputDir)
	//	if err != nil {
	//		global.CW_LOG.Error("GenerateApi", zap.String("remove ", err.Error()))
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": err.Error(),
	//		})
	//		return
	//	}
	//	global.CW_LOG.Info("成功删除文件夹：" + *outputDir)
	//}
	err = generate(tmpFile, outputDir)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var src = *outputDir
	var dst = fmt.Sprintf("%s.tar.gz", src)
	// 生成 gzip
	if err := utils.Tar(src, dst); err != nil {
		global.CW_LOG.Error("generateApi", zap.String("gzip ", err.Error()))
		response.FailWithMessage(err.Error(), c)
		return
	}
	//返回gzip
	c.File(dst)
	response.OkWithMessage("/cryptogen/generate   ok!", c)
	return
}

func generate(fileName, outputDir *string) error {
	file, err := os.Open(*fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	err = tool.Generate(&file, outputDir)
	if err != nil {
		return err
	}
	return nil
}

func MultiGenerateApi(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	files := form.File["file"]
	outputDir := &global.CW_CONFIG.System.GenerateDir
	//检测上传文件并批量生成所有文件
	if err = multiGen(c, files, outputDir); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var src = *outputDir
	var dst = fmt.Sprintf("%s.tar.gz", src)
	// 生成 gzip
	if err := utils.Tar(src, dst); err != nil {
		global.CW_LOG.Error("generateApi", zap.String("gzip ", err.Error()))
		response.FailWithMessage(err.Error(), c)
		return
	}
	//返回gzip
	c.File(dst)
	response.Ok(c)
	return
}

func multiGen(c *gin.Context, files []*multipart.FileHeader, outputDir *string) (err error) {
	if len(files) == 0 {
		return errors.New("no files uploaded")
	}
	tmpFile := &global.CW_CONFIG.System.TempFile
	//存在则删除
	if isExit, _ := utils.PathExists(*outputDir); isExit {
		err = utils.RemoveDir(*outputDir)
		if err != nil {
			global.CW_LOG.Error("multiGen", zap.String("remove ", err.Error()))
			return err
		}
		global.CW_LOG.Info("成功删除文件夹：" + *outputDir)
	}
	for _, file := range files {
		if err = c.SaveUploadedFile(file, *tmpFile); err != nil {
			return err
		}
		if err = generate(tmpFile, outputDir); err != nil {
			return err
		}
		if err = os.Remove(*tmpFile); err != nil {
			return err
		}
	}
	return nil
}

func ProtoDecode(c *gin.Context) {
	//todo
	msgName := c.Query("msgName")
	fmt.Println("msgName:", msgName)

	msgType := proto.MessageType(msgName) //这里版本需要和fabric中引入一致！！！
	fmt.Println("msgType:", msgType)
	if msgType == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "msgType nil",
		})
		return
	}
	msg := reflect.New(msgType.Elem()).Interface().(proto.Message)
	r := c.Request
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "read err",
		})
		return
	}
	err = proto.Unmarshal(buf, msg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Unmarshal err",
		})
		return
	}
	//todo

	c.JSON(http.StatusOK, gin.H{
		"message": "protoDecode",
	})
	return
}

func ProfileGenerate(c *gin.Context) {
	//http://localhost:8083/configtxgen/profile?outputBlock=genesis123.block&profile=TwoOrgsOrdererGenesis&channelID=system-channel
	//todo
	profile := c.Query("profile")
	channelID := c.Query("channelID")
	outputBlock := c.Query("outputBlock")
	outputChannelCreateTx := c.Query("outputCreateChannelTx")
	outputAnchorPeersUpdate := c.Query("outputAnchorPeersUpdate")
	asOrg := c.Query("asOrg")

	fmt.Println("profile :", profile)
	fmt.Println("channelID :", channelID)
	fmt.Println("outputBlock :", outputBlock)
	fmt.Println("outputCreateChannelTx :", outputChannelCreateTx)
	fmt.Println("outputAnchorPeersUpdate :", outputAnchorPeersUpdate)

	if channelID == "" && (outputBlock != "" || outputChannelCreateTx != "") {
		response.FailWithMessage("Missing channelID, please specify it with '-channelID'", c)
		return
	}
	err := factory.InitFactories(nil)
	if err != nil {
		response.FailWithMessage("Error on initFactories: "+err.Error(), c)
		return
	}
	var profileConfig *genesisconfig.Profile
	if outputBlock != "" || outputChannelCreateTx != "" || outputAnchorPeersUpdate != "" {
		if profile == "" {
			response.FailWithMessage("The '-profile' is required when '-outputBlock', '-outputChannelCreateTx', or '-outputAnchorPeersUpdate' is specified", c)
			return
		}
		// 接收配置文件
		file, err := c.FormFile("file")
		if err != nil || file == nil {
			response.FailWithMessage("file not upload ", c)
			return
		}
		//是否接收 crypto-config文件
		if global.CW_CONFIG.System.CryptoConfig {
			// todo
		}

		// 配置环境变量
		pwd, _ := os.Getwd()
		cfgPath := filepath.Join(pwd, global.CW_CONFIG.System.ConfigTxPath)
		err = os.Setenv("FABRIC_CFG_PATH", cfgPath)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		fmt.Println("FABRIC_CFG_PATH :", os.Getenv("FABRIC_CFG_PATH"))

		// 把配置文件保存到相应位置
		configTxFile := filepath.Join(os.Getenv("FABRIC_CFG_PATH"), global.CW_CONFIG.System.ConfigTxFile)
		err = c.SaveUploadedFile(file, configTxFile)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}

		//加载配置
		profileConfig, err = genesisconfig.Load(profile)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		//res, _ := json.Marshal(profileConfig)
		//fmt.Println("profileConfig:", string(res))
	}
	if outputBlock != "" {
		//文件名做限制
		//if outputBlock[0:7] != "genesis" && outputBlock[len(outputBlock)-5:] != "block" {
		if outputBlock[0:7] != "genesis" {
			response.FailWithMessage("outputBlock name must start with 'genesis'!", c)
			return
		}
		//生成
		if err := tool.DoOutputBlock(profileConfig, channelID, outputBlock); err != nil {
			global.CW_LOG.Info("Error on DoOutputBlock: %s" + err.Error())
			response.FailWithMessage(err.Error(), c)
			return
		}
		//check
		if !utils.CheckFileIsExist(outputBlock) {
			global.CW_LOG.Info("Error on outputBlock: get outputBlock file fail!")
			response.FailWithMessage("get outputBlock file fail!", c)
			return
		}
		//返回
		c.File(outputBlock)

		response.OkWithData("ProfileGenerate success", c)
		return
	}
	// add outputChannelCreateTx ;2021-07-08   生成channel源文件
	if outputChannelCreateTx != "" {
		if !outputTxNameCheck(outputChannelCreateTx) {
			response.FailWithMessage("outputChannelCreateTx name must end with 'tx'!", c)
			return
		}
		//生成
		if err := tool.DoOutputChannelCreateTx(profileConfig, nil, channelID, outputChannelCreateTx); err != nil {
			global.CW_LOG.Info("Error on DoOutputChannelCreateTx: %s" + err.Error())
			response.FailWithMessage(err.Error(), c)
			return
		}
		//check
		if !utils.CheckFileIsExist(outputChannelCreateTx) {
			global.CW_LOG.Info("Error on outputChannelCreateTx: get outputChannelCreateTx file fail!")
			response.FailWithMessage("get outputChannelCreateTx file fail!", c)
			return
		}
		//返回
		c.File(outputChannelCreateTx)
		response.OkWithMessage("ProfileGenerate outputChannelCreateTx success ", c)
		return
	}
	// add outputAnchorPeersUpdate ;2021-07-14 生成组织的锚节点文件
	if outputAnchorPeersUpdate != "" {
		if !outputTxNameCheck(outputAnchorPeersUpdate) {
			response.FailWithMessage("outputAnchorPeersUpdate name must end with 'tx'!", c)
			return
		}
		//生成
		if err := tool.DoOutputAnchorPeersUpdate(profileConfig, channelID, outputAnchorPeersUpdate, asOrg); err != nil {
			global.CW_LOG.Info("Error on DoOutputAnchorPeersUpdate: %s" + err.Error())
			response.FailWithMessage("Error on inspectChannelCreateTx: "+err.Error(), c)
			return
		}
		//check
		if !utils.CheckFileIsExist(outputAnchorPeersUpdate) {
			global.CW_LOG.Info("Error on outputAnchorPeersUpdate: get outputAnchorPeersUpdate file fail!")
			response.FailWithMessage("get outputAnchorPeersUpdate file fail!", c)
			return
		}
		//返回
		c.File(outputAnchorPeersUpdate)
		response.OkWithMessage("ProfileGenerate outputAnchorPeersUpdate success ", c)
		return
	}

	response.Fail(c)
	return
}

func outputTxNameCheck(handler string) bool {
	//文件名限制
	return handler[len(handler)-2:] == "tx"
}

func Test(c *gin.Context) {
	//file,err := c.FormFile("file")
	//fmt.Println("err : ",err)
	//fmt.Println("fileName :",file.Filename)
	outputBlock := c.Query("outputBlock")
	outputBlock = "go.mod"
	fmt.Println("outputBlock :", outputBlock)
	//第一种下载方式
	//c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", "attachment; filename="+outputBlock)
	//c.Header("Content-Transfer-Encoding", "binary")
	//c.Header("Cache-Control", "no-cache")
	//c.Header("Content-Length","600")
	// 这里使用postman或是api类工具会导致panic，浏览器没有问题
	//返回文件流
	c.File(outputBlock)

	response.Ok(c)
	return
}
