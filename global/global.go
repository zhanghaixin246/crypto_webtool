package global

import (
	"crypto_webtool/config"
	"github.com/spf13/viper180"
	"go.uber.org/zap"
)

const (
	ConfigFile = "./config/config.yaml"
)

var (
	CW_VP     *viper.Viper
	CW_CONFIG config.Server
	CW_LOG    *zap.Logger
)
