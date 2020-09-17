package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type User struct {
	// ID 自增ID
	ID uint32 `json:"id"`
	// Company 用户公司
	Company string `json:"company"`
	// Type 用户类型
	Type uint8 `json:"type"`
	// Title 用户类型名称
	Title string `json:"title"`
	// Passwd 密码(md5)
	Passwd string `json:"passwd"`
	// Name 姓名
	Name string `json:"name"`
	// NickName 用户别名
	NickName string `json:"nick_name"`
	// Email 邮箱
	Email string `json:"email"`
	// Phone 电话
	Phone string `json:"phone"`
	// State 用户状态(1新建 2正常 3本平台禁用)
	State int8 `json:"state"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time"`
	// ModifyTime 最后修改时间
	ModifyTime time.Time `json:"modify_time"`
	// DeletedAt 软删除
	DeletedAt time.Time `json:"deleted_at"`
}

var connectionString string
var Db *sqlx.DB

func init() {
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DataBase)
	Db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("connection error", err)
	}
	defer Db.Close()
}

func GetRecords() {
	var users []User
	err := Db.Select(&users, "select * from user")
	if err != nil {
		fmt.Fprintln(os.Stdin, "get result error", err)
	}
	fmt.Println("the result data:", users)
}
