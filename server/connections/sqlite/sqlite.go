package sqlite

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var connSQLite *gorm.DB

var sqlitePath *string

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
func New(path *string) {
	sqlitePath = path
}

//GetConn connection
func GetConn() *gorm.DB {

	defer cleanup()

	if connSQLite != nil {
		log.Print("Reapproving connection sqlite")
		return connSQLite
	}

	log.Printf("SQLite connecting ... (path=%v) !!!", *sqlitePath)

	conn, err := gorm.Open("sqlite3", *sqlitePath)

	checkError(&err)

	log.Printf("SQLite connected in %v !!!", *sqlitePath)

	connSQLite = conn

	return connSQLite
}

//Close connection
func Close() {
	defer cleanup()
	if connSQLite != nil {
		connSQLite.Close()
	}
}
