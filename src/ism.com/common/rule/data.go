package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type DataStructure struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Data        []Data            `json:"data"`
	Lengths     []LengthFieldInfo `json:"lengths"`
	MessageType int               `json:"messageType"`
}

type Data struct {
	Id                 string
	DataStrtId         string
	MasterFieldGroupId string
	Detail             []Detail
	RecordDelimeter    string
	Master             FieldGroup
	DataIndex          int
}

type Detail struct {
	Id                    string
	Detail                FieldGroup
	RepeatCount           int
	RepeatCountDataIndex  int
	RepeatCountFieldIndex int
	GroupType             int
}

type LengthFieldInfo struct {
	LengthDataIndex   int
	LengthDetailIndex int
	LengthFieldIndex  int

	DataIndex   int
	DetailIndex int
	FieldIndex  int
	DiffValue   int
}

type FieldGroup struct {
}

type Field struct {
	Id          string
	Name        string
	FieldType   string
	FieldLength int
}

func GetDataStructure(id string) (string, error) {
	var inf DataStructure
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(dStrt_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&inf.Id, &inf.Name, &inf.MessageType)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	b, err := json.Marshal(inf)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getData(id string, dbConn *sql.DB) ([]Data, error) {
	var data []Data

	stmt, err := dbConn.Prepare(rSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	// err = stmt.QueryRow(id).Scan(&svc.MsgLog, &svc.ProcLog, &svc.AsyncRCnt, &svc.AsyncRTOut)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		println("not found")
	// 	} else {
	// 		panic(err.Error())
	// 	}
	// }
	return data, nil
}
