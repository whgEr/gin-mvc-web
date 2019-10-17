package config

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

var _db *gorm.DB

func init() {

	var err error
	_db, err = gorm.Open("mysql", "root:123@tcp(localhost:3306)/test?charset=utf8")
	//不能使用下面这句，否则全局_db赋值不成功
	//_db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/test?charset=utf8")

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	//debug mod
	_db.LogMode(true)

	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(10)
	_db.DB().SetMaxIdleConns(5)

	//defer DB.Close() //连接池模式不能关闭，导致整个连接池关闭，没有可用的连接。
	
	//return _db
}

func GetDB() *gorm.DB {

	return _db

}