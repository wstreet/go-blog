package conf

import (
	"fmt"
	"go-blog/model"
	"strings"

	"github.com/spf13/viper"
)

var Conf = new(TotalConfig)

type TotalConfig struct {
	*MySQLConfig  `mapstructure:"mysql"`
	*ServerConfig `mapstructure:"server"`
}
type ServerConfig struct {
	Port string `mapstructure:"port"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Suffix       string `mapstructure:"suffix"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init() {
	viper.SetConfigFile("./conf/conf.toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	fmt.Println(Conf.MySQLConfig)

	path := strings.Join([]string{Conf.Username, ":", Conf.Password, "@tcp(", Conf.Host, ":", fmt.Sprintf("%d", Conf.MySQLConfig.Port), ")/", Conf.Dbname, "?", Conf.Suffix}, "")
	model.Database(path)
	// TODO: cache redis
}
