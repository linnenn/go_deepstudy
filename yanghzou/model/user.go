package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
	conf "yanghzou/utill"
)

// User 用户模型,注意标签，其中json为json返回数据，db为对应的数据库中字段
// 可以使用sqlx或者gorm,都可以很好的满足目前的工作
type User struct {
	// ID 自增ID
	ID uint32 `json:"id" db:"id"`
	// Company 用户公司
	Company string `json:"company" db:"company"`
	// Type 用户类型
	Type uint8 `json:"type" db:"type"`
	// Title 用户类型名称
	Title string `json:"title" db:"title"`
	// Passwd 密码(md5)
	Passwd string `json:"passwd" db:"passwd"`
	// Name 姓名
	Name string `json:"name" db:"name,omitempty"`
	// NickName 用户别名
	NickName *string `json:"nick_name" db:"nick_name"`
	// Email 邮箱
	Email *string `json:"email" db:"email"`
	// Phone 电话
	Phone *string `json:"phone" db:"phone" `
	// State 用户状态(1新建 2正常 3本平台禁用)
	State int8 `json:"state" db:"state"`
	// CreateTime 创建时间
	CreateTime time.Time `json:"create_time" db:"create_time"`
	// ModifyTime 最后修改时间
	ModifyTime time.Time `json:"modify_time" db:"modify_time"`
	// DeletedAt 软删除,注意sql.NullTime跟指针类型不一致，数据库返回结构上是有区别的
	//*返回 null ,sql.NullType包括：值和valid两个
	//"NickName": {
	//                "String": "运营人员",
	//                "Valid": true
	//            },
	//            "Email": {
	//                "String": "admin@moviebook.cn",
	//                "Valid": true
	//            },
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at,omitempty"`
}

var connectionString string

// Db 数据库连接串
var Db *sqlx.DB

func init() {
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", conf.User, conf.Password, conf.Host, conf.Port, conf.DataBase)
	var err error
	Db, err = sqlx.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("connection error", err)
		return
	}
}

// GetRecords 记录
func GetRecords() interface{}{
	var users []User
	err := Db.Select(&users, "select * from user")
	if err != nil {
		fmt.Fprintln(os.Stdin, "get result error", err)
	}
	return users
}

func Insert()  {
	r,err := Db.Exec("insert into user (`company`,`type`,`title`,`passwd`,`name`,`email`) values(?,?,?,?,?,?) ","影普",1,"test","ejoiafjoaf","yanglei","dan@126.com")
	if err != nil {
		fmt.Println("insert error",err)
		return
	}
	id,err := r.LastInsertId()
	if err != nil {
		fmt.Println("get result error",err)
	}
	fmt.Println("insert successful last_id:",id)
}

func Update()  {
	r,err := Db.Exec("update user set nick_name = ? where id = ?","小格格1",3)
	if err != nil{
		fmt.Println("update dailed:",err)
		return
	}
	records,err := r.RowsAffected()
	if err != nil {
		fmt.Println("update failed",err)
	}
	fmt.Println("update successful the effect rows:",records)
}

func GetUserRecord(uid int) interface{} {
	var user User
	err := Db.Get(&user,"select * from user where id = ?",uid)
	if err != nil{
		fmt.Fprintf(os.Stdout,"get record failed %s",err)
	}
	return user
}