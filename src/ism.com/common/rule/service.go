package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type Service struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	SvcBlock      string `json:"svcBlock"`
	InDstrId      string `json:"inDstrId"`
	OutDstrId     string `json:"outDstrId"`
	ErrDstrId     string `json:"errDstrId"`
	ServiceType   string `json:"svcType"`
	InterfaceType string `json:"infType"`
	SendRcvType   string `json:"srType"`

	DBSvc DBService `json:"dbSvc"`
}

type DBService struct {
	CRUDType  string `json:"crudType"`
	TableName string `json:"tName"`
	Query     string `json:"query"`
	QueryType string `json:"qryType"`
	// DateManageClass    string
	// TableNameInitClass string
	// QueryHandlerClass  string
	SPType string `json:"srType"`
	// FilterHandler      string
	DefaultCRUDType  string `json:"defaultCRUD"`
	FilterCheckType  string `json:"fChkType"`
	TargetFetchCount int    `json:"tgtFetchCnt"`
}

func GetService(id string) (string, error) {
	var svc Service
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(svc_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&svc.Id, &svc.Name, &svc.SvcBlock, &svc.InDstrId, &svc.OutDstrId, &svc.ErrDstrId,
		&svc.ServiceType, &svc.InterfaceType, &svc.SendRcvType)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	switch svc.ServiceType {
	case "D":
		if svc.DBSvc, err = getDBService(id, dbConn); err != nil {
			panic(err)
		}
	}

	b, err := json.Marshal(svc)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getDBService(id string, dbConn *sql.DB) (DBService, error) {
	var svc DBService

	stmt, err := dbConn.Prepare(dbSvc_sql)
	if err != nil {
		panic(err.Error())
	}

	err = stmt.QueryRow(id).Scan(&svc.CRUDType, &svc.TableName, &svc.Query, &svc.QueryType, &svc.SPType,
		&svc.DefaultCRUDType, &svc.FilterCheckType, &svc.TargetFetchCount)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
		} else {
			panic(err.Error())
		}
	}

	return svc, nil
}
