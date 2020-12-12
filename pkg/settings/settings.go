package settings

import (
	"fmt"
	"github.com/spf13/viper"
)

type yamlSettings struct {
	Log log
	Bot bot
}

type log struct {
	Level string
	File  string
}

type bot struct {
	UseProxy bool
}

func Init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.AddConfigPath("./config")
	config.AddConfigPath("../../config")
	config.AddConfigPath("/config")
	config.AddConfigPath("/")
	_ = config.ReadInConfig()
	unmarshal(config)
}

// 将配置文件映射到 yamlSettings
func unmarshal(v *viper.Viper) {
	if err := v.Unmarshal(&Settings); err != nil {
		_ = fmt.Errorf("filed to unmarshal config file: %w", err)
	}
}
