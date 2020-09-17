package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"yanghzou/utill"
)

var DbInit = DbWorkerConnect{}

type DbWorkerConnect struct {
	Conn *sqlx.DB
}

func NewDbWorkerConnect() {
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", utill.User, utill.Password, utill.Host, utill.Port, utill.DataBase)

	dataSourceName = dataSourceName + "?parseTime=true&loc=Asia%2FShanghai&charset=utf8-mb4"
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	// 设置理想最大空闲连接数
	db.SetMaxIdleConns(2)
	//设置最大的连接数
	db.SetMaxOpenConns(10)
	DbInit.Conn = db
}
