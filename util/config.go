package util

import (
	"time"

	"github.com/spf13/viper"
)
type Config struct {
    DBDriver      string `mapstructure:"DB_DRIVER"`
    DBSource      string `mapstructure:"DB_SOURCE"`
    HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDR"`
	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDR"`
	KeyString string `mapstructure:"KEY_STRING"`
	AccessTokenDuration time.Duration	`mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("app")
    viper.SetConfigType("env")
	viper.AutomaticEnv();

	err=viper.ReadInConfig();
	if err!=nil{
		return
	}
	err=viper.Unmarshal(&config);
	return
}
