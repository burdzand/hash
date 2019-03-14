package mysql

import (
	"fmt"

	"log"

	"github.com/burdzand/hash/server/cfg"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var connMysql *gorm.DB

var mysqlHost *string
var mysqlPort *string
var mysqlDatabase *string
var mysqlUser *string
var mysqlPassword *string

func init() {
}

func checkError(err *error) {
	if *err != nil {
		panic(*err)
	}
}

func cleanup() {
	if r := recover(); r != nil {
		log.Print(r)
	}
}

/**
Configurado por variavel de amviente
TODO: switch to configuration file
**/
func New() {
	mysqlHost = cfg.GetString("connections/mysql/host", "localhost")
	mysqlPort = cfg.GetString("connections/mysql/port", "3306")
	mysqlDatabase = cfg.GetString("connections/mysql/database", "mysql")
	mysqlUser = cfg.GetString("connections/mysql/user", "mysql")
	mysqlPassword = cfg.GetString("connections/mysql/password", "mysql")
}

//GetConn connection
func GetConn() *gorm.DB {

	defer cleanup()

	if connMysql != nil {
		log.Print("Reapproving connection mysql")
		return connMysql
	}

	log.Printf("%v @tcp(%v:%v)/%v", *mysqlUser, *mysqlHost, *mysqlPort, *mysqlDatabase)
	config := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", *mysqlUser, *mysqlPassword, *mysqlHost, *mysqlPort, *mysqlDatabase)

	conn, err := gorm.Open("mysql", config)

	checkError(&err)
	log.Printf("mysql connected in %v !!!", *mysqlHost)
	connMysql = conn
	return connMysql
}

//Close connection
func Close() {
	defer cleanup()
	if connMysql != nil {
		connMysql.Close()
	}
}
