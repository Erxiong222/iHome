package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"ihome/conf"
)
var globalDB *gorm.DB

//一个函数一个功能    50行
func InitDb()error{
	//打开数据库
	//拼接链接字符串
	connString := conf.MysqlName+":"+conf.MysqlPwd+"@tcp("+conf.MysqlAddr+":"+conf.MysqlPort+")/"+conf.MysqlDB+"?parseTime=true"
	db,err  := gorm.Open("mysql",connString)
	if err != nil {
		fmt.Println("链接数据库失败",err)
		return err
	}

	//连接池设置
	//设置初始化数据库链接个数
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(70)
	db.DB().SetConnMaxLifetime(60 * 5)

	//默认情况下表名是复数
	db.SingularTable(true)
	globalDB = db
	return nil
}

func main() {
	InitDb()

}
