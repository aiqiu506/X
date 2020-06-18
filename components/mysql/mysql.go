//+build mysql  !production

package mysql

import (
	"fmt"
	"github.com/aiqiu506/x/global"
	"github.com/aiqiu506/x/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Config struct {
	Host        string `map:"host"`
	Port        string `map:"port"`
	DBName      string `map:"db"`
	User        string `map:"user"`
	Pwd         string `map:"pwd"`
	MaxCons     int    `map:"maxCons"`
	MaxFreeCons int    `map:"MaxFreeCons"`
}

type mySqlStruct struct {
	DB *gorm.DB
}

func Connect(my *Config) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", my.User, my.Pwd, my.Host, my.Port, my.DBName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	if my.MaxFreeCons != 0 {
		db.DB().SetMaxOpenConns(my.MaxCons)
	}
	if my.MaxCons != 0 {
		db.DB().SetMaxIdleConns(my.MaxFreeCons)
	}
	return db
}

func (m mySqlStruct) NewComponent(config interface{}) {
	mysqlParams := &Config{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, mysqlParams)
		if err != nil {
			log.Fatal(err)
		}
		m.DB = Connect(mysqlParams)
	} else {
		log.Fatal("mysql配置文件错误")
	}

}

var Mysql mySqlStruct

func init() {
	//注册组件
	global.Global.Register("mysql", &Mysql)
}
