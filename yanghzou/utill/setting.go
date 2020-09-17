package utill

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/spf13/viper"
	"time"
)

var (
	File         *ini.File
	Host         string
	Port         int
	User         string
	Password     string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	DataBase     string
)
var (
	RedisHost string
	RedisPort int
	RedisAuth string
	RedisWriteTimeout time.Duration
	RedisReadTimeout  time.Duration
)

func init() {
	//加载配置文件，读取配置
	viper.SetConfigFile("./config/dev_env.yaml")
	//viper.AddConfigPath("./config/dev_env.yaml") // ini.Load("")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("error config file not found:%s",err)
	}
	loadDatabase()
	loadRedis()
}

func loadDatabase() {
	sec, err := File.GetSection("database")
	if err != nil {
		fmt.Println("section database get error",err)
	}
	Host = sec.Key("host").String()
	Port = sec.Key("port").MustInt(3306)
	User = sec.Key("user").String()
	Password = sec.Key("password").String()
	DataBase = sec.Key("database").String()
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(300)) * time.Second
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(300)) * time.Second
}

func loadRedis()  {
	sec,err := File.GetSection("redis")
	if err != nil {
		fmt.Println("section redis get error",err)
	}
	RedisHost = sec.Key("host").String()
	RedisPort = sec.Key("port").MustInt(6379)
	RedisAuth = sec.Key("auth").String()
	RedisWriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(300))*time.Second
	RedisReadTimeout  = time.Duration(sec.Key("read_timeout").MustInt(300))*time.Second
}
