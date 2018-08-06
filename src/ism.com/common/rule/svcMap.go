package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type ServiceMap struct {
	Id              string
	OutSvcId        string
	DataStructureId string

	InputServices []InputService
	MasterMap     [][]DataMap
	DetailMap     [][][]DataMap
}

type DataMap struct {
	Dataindex      int
	Detailindex    int
	Columnindex    int
	FieldId        string
	CustomFunction string

	Sources []SourceColumn
}

type SourceColumn struct {
	MappingIndex       int
	SourceDataIndex    int
	SourceDetailIndex  int
	SourceFieldIndex   int
	DefaultValue       string
	FieldId            string
	SourceMessageIndex int
	Path               string
	IsXml              string
}

type InputService struct {
	SvcId           string `json:"svcId"`
	DataStructureId string `json:"dstrId"`
}

func GetServiceMap(id string) (string, error) {
	var sMap ServiceMap
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(sMap_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&sMap.Id, &sMap.OutSvcId, &sMap.DataStructureId)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	if sMap.InputServices, err = getInputServices(id, dbConn); err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	b, err := json.Marshal(sMap)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getDataMap(id string, dbConn *sql.DB) ([]DataMap, error) {
	var dMaps []DataMap

	stmt, err := dbConn.Prepare(inSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		// get RawBytes from data
		var dMap DataMap
		if err := rows.Scan(&dMap.Dataindex, &dMap.Detailindex, &dMap.Columnindex, &dMap.FieldId, &dMap.CustomFunction); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		if dMap.Sources, err = getSourceColumns(id, dMap.Dataindex, dMap.Detailindex, dMap.Columnindex, dbConn); err != nil {
			fmt.Printf("Error: %s", err)
			return nil, err
		}

		dMaps = append(dMaps, dMap)
	}
	return dMaps, nil
}

func getSourceColumns(id string, dIdx int, dtIdx int, cIdx int, dbConn *sql.DB) ([]SourceColumn, error) {
	var scs []SourceColumn

	stmt, err := dbConn.Prepare(inSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		// get RawBytes from data
		var sc SourceColumn
		if err := rows.Scan(&sc.MappingIndex, &sc.SourceDataIndex, &sc.SourceDetailIndex, &sc.SourceFieldIndex,
			&sc.DefaultValue, &sc.FieldId, &sc.SourceMessageIndex, &sc.Path, &sc.IsXml); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		scs = append(scs, sc)
	}
	return scs, nil
}

func getInputServices(id string, dbConn *sql.DB) ([]InputService, error) {
	var inputs []InputService

	stmt, err := dbConn.Prepare(inSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		// get RawBytes from data
		var input InputService
		if err := rows.Scan(&input.SvcId, &input.DataStructureId); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		inputs = append(inputs, input)
	}
	return inputs, nil
}
