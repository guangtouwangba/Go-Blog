package db

import (
	"Go-Blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func CreateDatabaseConnection() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Println("Connecting to database...", config.GetYamlConfig().GetDb())
	dbase := config.GetYamlConfig().GetDb()
	Db, err := gorm.Open(mysql.Open(dbase), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	log.Println("create database...", config.GetYamlConfig().Mysql.Database)
	Db.Exec("CREATE DATABASE IF NOT EXISTS " + config.GetYamlConfig().Mysql.Database)
	if Db.Error != nil {
		log.Panicln("err:", Db.Error.Error())
	}

	log.Println("Connect database success")
}
