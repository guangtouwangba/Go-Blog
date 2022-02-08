package db

import (
	"Go-Blog/config"
	"Go-Blog/internal/domain/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type MysqlTool struct {
	Db *gorm.DB
}

var instance *MysqlTool
var once sync.Once

func (m *MysqlTool) CreateDatabaseConnection() (success bool) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Println(config.GetYamlConfig().GetDb() + config.GetYamlConfig().Mysql.Params)
	dbase := config.GetYamlConfig().GetDb() + config.GetYamlConfig().Mysql.Params
	Db, err := gorm.Open(mysql.Open(dbase), &gorm.Config{})
	m.Db = Db
	if err != nil {
		log.Panicln("err:", err.Error())
		return false
	}
	Db.Exec("CREATE DATABASE IF NOT EXISTS " + config.GetYamlConfig().Mysql.Database)
	if Db.Error != nil {
		log.Panicln("err:", Db.Error.Error())
		return false
	}
	m.InitTables(Db)
	log.Println("Connect database success")
	return true
}

func (m *MysqlTool) InitTables(db *gorm.DB) {
	db.Exec("USE " + config.GetYamlConfig().Mysql.Database)
	err := db.AutoMigrate(&po.User{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}

func (m *MysqlTool) GetInstance() *MysqlTool {
	once.Do(func() {
		instance = &MysqlTool{}
		instance.CreateDatabaseConnection()
	})
	return instance
}

func (m *MysqlTool) GetDb() *gorm.DB {
	return m.Db
}
