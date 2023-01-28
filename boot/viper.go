package boot

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	g "tiktok/app/global"
)

const (
	configEnv  = "TIKTOK_CONFIG_PATH"
	configFile = "manifest/config/config.yaml"
)

func ViperSetup(path ...string) {
	var configPath string

	// get config file path
	// priority: param > command line > environment > default
	if len(path) != 0 {
		// param
		configPath = path[0]
	} else {
		// command line
		flag.StringVar(&configPath, "c", "", "set config path")
		flag.Parse()

		if configPath == "" {
			if configPath = os.Getenv(configEnv); configPath != "" {
				// environment
			} else {
				// default
				configPath = configFile
			}
		}
	}
	fmt.Printf("get config path: %s", configPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("get config file failed, err: %v", err))
	}

	if err := v.Unmarshal(&g.Config); err != nil {
		panic(fmt.Errorf("unmarshal config failed, err: %v", err))
	}
}
