package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v2"
)

var dbConn *sql.DB
var dbInf dbInfo

type dbs struct {
	Dbs []dbInfo
}
type dbInfo struct {
	DbName   string
	Url      string
	DbId     string
	DbPasswd string
}

func init() {
	path := os.Getenv("GO_PROP")
	if path == "" {
		path = "."
	}
	yamlFile, err := ioutil.ReadFile(fmt.Sprint(path + "/mysql.yml"))
	if err != nil {
		log.Printf("Read config file error #%v", err)
	}
	var dInfos *dbs
	err = yaml.Unmarshal(yamlFile, &dInfos)
	if err != nil {
		log.Printf("Unmarshal config file error #%v", err)
	}
	dbInf = dInfos.Dbs[0]
}

func GetDatabase() *sql.DB {
	var err error
	if dbConn == nil {
		dbConn, err = sql.Open("mysql", fmt.Sprint(dbInf.DbId, ":", dbInf.DbPasswd, "@", dbInf.Url, "/", dbInf.DbName))
		// if there is an error opening the connection, handle it
		if err != nil {
			panic(err.Error())
		}

		dbConn.Ping()
		dbConn.SetMaxIdleConns(5)
		dbConn.SetMaxOpenConns(100)
	}
	return dbConn
}
