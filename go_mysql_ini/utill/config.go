package utill

import (
	"github.com/go-ini/ini"
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

func init() {
	var err error
	File, err = ini.Load("conf/config.ini")
	if err != nil {
		panic("error config file not found:".err)
	}
}

func loadDatabase() {
	sec, err := File.GetSection("database")
	if err != nil {
		panic("section database get error".err)
	}
	Host = sec.Key("host").String()
	Port = sec.Key("port").MustInt(3306)
	User = sec.Key("user").String()
	Password = sec.Key("password").String()
	DataBase = sec.Key("database").String()
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(300)) * time.Second
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(300)) * time.Second
}
