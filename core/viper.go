package core

import (
	"crypto_webtool/global"
	"fmt"
	"github.com/spf13/viper180"
	"github.com/fsnotify/fsnotify"
)

func Viper() *viper.Viper {
	config := global.ConfigFile
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CW_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CW_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
