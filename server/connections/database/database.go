package database

import (
	"os"
	"path/filepath"

	"log"

	"github.com/burdzand/hash/server/cfg"
	"github.com/burdzand/hash/server/connections/mysql"
	"github.com/burdzand/hash/server/connections/postgres"
	"github.com/burdzand/hash/server/connections/sqlite"
	"github.com/jinzhu/gorm"
)

var connSQLite *gorm.DB

var database *string

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
func init() {
	database = cfg.GetString("connections/database", "postgres")

	// Start Connections
	if *database == "postgres" {
		postgres.New()
		connSQLite = postgres.GetConn()
	} else if *database == "mysql" {
		mysql.New()
		connSQLite = mysql.GetConn()
	} else if *database == "sqlite" {
		path := cfg.GetString("connections/sqlite/path", filepath.Join(os.TempDir(), "testing.db"))
		sqlite.New(path)
		connSQLite = sqlite.GetConn()
	}
}

//GetConn connection
func GetConn() *gorm.DB {
	return connSQLite
}

//Close connection
func Close() {
	defer cleanup()
	if connSQLite != nil {
		connSQLite.Close()
	}
}
