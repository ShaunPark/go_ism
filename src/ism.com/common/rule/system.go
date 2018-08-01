package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type System struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	LoadBalanceType string `json:"lbType"`
	UseHealth       string `json:"useHealth"`
	AutoRecovery    string `json:"autoRecovery"`
	CheckGid        string `json:"checkGid"`
	CallLimit       int    `json:"callLimit"`
}

type Application struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Port     int           `json:"port"`
	UserId   string        `json:"userId"`
	Password string        `json:"password"`
	AppType  string        `json:"appType"`
	DBApp    DBApplication `json:"dbApp"`
}

type DBApplication struct {
	DBName        string `json:"dbName"`
	JDBCType      string `json:"jdbcType"`
	DBType        string `json:"dbType"`
	DBParameter   string `json:"dbParam"`
	AppSystemName string `json:"appSysName"`
}

type Server struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IpAddress string `json:"ip"`
}

type AppInfo struct {
	AppIndex       int    `json:"appIdx"`
	IsMaster       string `json:"isMaster"`
	CallWeight     int    `json:"callWeight"`
	HealthStatus   string `json:"healthSt"`
	BackupIdx      int    `json:"backupIdx"`
	HealthInterval int    `json:"healthInterval"`
	ApplicationId  string `json:"appId"`
	ServerId       string `json:"svrId"`
}

func GetServer(id string) (string, error) {
	var svr Server
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(svr_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&svr.Id, &svr.Name, &svr.IpAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	b, err := json.Marshal(svr)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func GetApplication(id string) (string, error) {
	var app Application
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(app_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&app.Id, &app.Name, &app.Port, &app.UserId, &app.Password, &app.AppType)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	switch app.AppType {
	case "D":
		if app.DBApp, err = getDBApplication(id, dbConn); err != nil {
			panic(err)
		}
	}

	b, err := json.Marshal(app)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func GetSystem(id string) (string, error) {
	var sys System
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(sys_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&sys.Id, &sys.Name, &sys.LoadBalanceType, &sys.UseHealth, &sys.AutoRecovery, &sys.CheckGid, &sys.CallLimit)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	b, err := json.Marshal(sys)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getDBApplication(id string, dbConn *sql.DB) (DBApplication, error) {
	var app DBApplication

	stmt, err := dbConn.Prepare(dbApp_sql)
	if err != nil {
		panic(err.Error())
	}
	err = stmt.QueryRow(id).Scan(&app.DBName, &app.JDBCType, &app.DBType, &app.DBParameter, &app.AppSystemName)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
		} else {
			panic(err.Error())
		}
	}

	return app, nil
}
