package infrastructure

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/witwoywhy/go-cores/vipers"
)

var AppConfig AppConfigInfo

func InitConfig() {
	vipers.Init()
	initAppConfig()
}

type AppConfigInfo struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

func initAppConfig() {
	var config AppConfigInfo
	if err := viper.UnmarshalKey("app", &config); err != nil {
		panic(fmt.Errorf("failed to loaded app config: %v", err))
	}

	AppConfig = config
}
