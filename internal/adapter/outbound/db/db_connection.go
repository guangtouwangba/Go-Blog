package db

import (
	"Go-Blog/config"
	"Go-Blog/internal/domain/do"
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
	isUpdate := config.GetYamlConfig().Mysql.Update

	nameToObj := map[string]interface{}{
		"users":   &do.User{},
		"article": &do.ArticleDO{},
		"comment": &do.CommentDO{},
		"tag":     &do.TagDO{},
	}

	for k, v := range nameToObj {
		if isUpdate {
			if dropTable(db, k) {
				continue
			}
		}

		err := db.AutoMigrate(v)
		if err != nil {
			log.Panicln("err:", err.Error())
		}
	}

}

func dropTable(db *gorm.DB, tableName string) bool {
	if db.Migrator().HasTable(tableName) {
		err := db.Migrator().DropTable(tableName)
		if err != nil {
			return true
		}
	}
	return false
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
