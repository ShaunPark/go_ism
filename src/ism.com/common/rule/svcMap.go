package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type ServiceMap struct {
	Id              string `json:"id"`
	OutSvcId        string `json:"oSvc"`
	DataStructureId string `json:"dstrt"`

	InputServices []InputService `json:"inSvcs"`
	MasterMap     [][]DataMap    `json:"mMap"`
	DetailMap     [][][]DataMap  `json:"dMap"`
}

type DataMap struct {
	Dataindex      int        `json:"dIdx"`
	Detailindex    int        `json:"dtIdx"`
	Columnindex    int        `json:"cIdx"`
	FieldId        NullString `json:"fld"`
	CustomFunction NullString `json:"cFunc"`

	Sources []SourceColumn `json:"sCols"`
}

type SourceColumn struct {
	MappingIndex       int        `json:"mIdx"`
	SourceDataIndex    int        `json:"sdIdx"`
	SourceDetailIndex  int        `json:"sdtIdx"`
	SourceFieldIndex   int        `json:"sfIdx"`
	DefaultValue       NullString `json:"dValue"`
	FieldId            NullString `json:"fld"`
	SourceMessageIndex int        `json:"smIdx"`
	Path               NullString `json:"path"`
	IsXml              string     `json:"isXml"`
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

	var dMap []DataMap
	if dMap, err = getDataMap(id, dbConn); err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	for i, item := range dMap {
		if dMap[i].Sources, err = getSourceColumns(id, item.Dataindex, item.Detailindex, item.Columnindex, dbConn); err != nil {
			fmt.Printf("Error: %s", err)
			return "", err
		}

		if item.Dataindex < 0 {
		} else {
			if item.Detailindex < 0 {
				if len(sMap.MasterMap) <= item.Dataindex {
					sMap.MasterMap = append(sMap.MasterMap, make([]DataMap, 0))
				}
				sMap.MasterMap[item.Dataindex] = append(sMap.MasterMap[item.Dataindex], item)
			} else {
				if len(sMap.DetailMap) <= item.Dataindex {
					sMap.DetailMap = append(sMap.DetailMap, make([][]DataMap, 0))
				}
				if len(sMap.DetailMap[item.Dataindex]) <= item.Detailindex {
					sMap.DetailMap[item.Dataindex] = append(sMap.DetailMap[item.Dataindex], make([]DataMap, 0))
				}
				sMap.DetailMap[item.Dataindex][item.Detailindex] = append(sMap.DetailMap[item.Dataindex][item.Detailindex], item)
			}
		}
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

	stmt, err := dbConn.Prepare(dMap_sql)
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

	stmt, err := dbConn.Prepare(sCol_sql)
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(id, dIdx, dtIdx, cIdx)
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
