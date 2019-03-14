package postgres

import (
	"fmt"

	"log"

	"github.com/burdzand/hash/server/cfg"
	"github.com/jinzhu/gorm"

	//_ load dialects postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connPostgres *gorm.DB

var postgresHost *string
var postgresPort *string
var postgresDatabase *string
var postgresUser *string
var postgresPassword *string

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
	postgresHost = cfg.GetString("connections/postgres/host", "localhost")
	postgresPort = cfg.GetString("connections/postgres/port", "5432")
	postgresDatabase = cfg.GetString("connections/postgres/database", "postgres")
	postgresUser = cfg.GetString("connections/postgres/user", "postgres")
	postgresPassword = cfg.GetString("connections/postgres/password", "postgres")
}

//GetConn connection
func GetConn() *gorm.DB {

	defer cleanup()

	if connPostgres != nil {
		log.Print("Reapproving connection postgres")
		return connPostgres
	}

	log.Printf("Postgres connecting ... (host=%v, port=%v, database=%v, user=%v) !!!", *postgresHost, *postgresPort, *postgresDatabase, *postgresUser)
	config := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", *postgresHost, *postgresPort, *postgresUser, *postgresDatabase, *postgresPassword)
	conn, err := gorm.Open("postgres", config)

	checkError(&err)
	log.Printf("Postgres connected in %v !!!", *postgresHost)
	connPostgres = conn
	return connPostgres
}

//Close connection
func Close() {
	defer cleanup()
	if connPostgres != nil {
		connPostgres.Close()
	}
}
