package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type DataStructure struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	MessageType int    `json:"messageType"`

	Data    []Data            `json:"data"`
	Lengths []LengthFieldInfo `json:"lengths"`
}

type Data struct {
	DataIndex          int        `json:"dataIndex"`
	RecordDelimeter    NullString `json:"rDelimeter"`
	MasterFieldGroupId string     `json:"flgId"`

	Detail []Detail `json:"detail"`
}

type Detail struct {
	DetailFieldGroupId    string  `json:"flgId"`
	RepeatCount           int     `json:"rCount"`
	RepeatCountDataIndex  int     `json:"rCountDataIndex"`
	RepeatCountFieldIndex int     `json:"rCountFldIndex"`
	GroupType             NullInt `json:"groupType"`
}

type LengthFieldInfo struct {
	LengthDataIndex   int `json:"lengthDataIndex"`
	LengthDetailIndex int `json:"lengthDetailIndex"`
	LengthFieldIndex  int `json:"lengthFieldIndex"`
	DataIndex         int `json:"dataIndex"`
	DetailIndex       int `json:"detailIndex"`
	FieldIndex        int `json:"fieldIndex"`
	DiffValue         int `json:"diffvalue"`
}

func GetDataStructure(id string) (string, error) {
	var dstr DataStructure
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(dStrt_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&dstr.Id, &dstr.Name, &dstr.MessageType)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	if dstr.Data, err = getData(id, dbConn); err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	if dstr.Lengths, err = getLengthInfo(id, dbConn); err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	b, err := json.Marshal(dstr)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getData(id string, dbConn *sql.DB) ([]Data, error) {
	var data []Data

	stmt, err := dbConn.Prepare(data_sql)
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
		var datum Data
		if err := rows.Scan(&datum.DataIndex, &datum.RecordDelimeter, &datum.MasterFieldGroupId); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		data = append(data, datum)
	}

	for i, d := range data {
		if data[i].Detail, err = getDetail(id, dbConn, d.DataIndex); err != nil {
			panic(err.Error())
		}
	}

	return data, nil
}

func getDetail(id string, dbConn *sql.DB, dIdx int) ([]Detail, error) {
	var details []Detail

	stmt, err := dbConn.Prepare(detail_sql)
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(id, dIdx)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		// get RawBytes from data
		var detail Detail
		if err := rows.Scan(&detail.DetailFieldGroupId, &detail.RepeatCount, &detail.RepeatCountDataIndex,
			&detail.RepeatCountFieldIndex, &detail.GroupType); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		details = append(details, detail)
	}

	return details, nil
}

func getLengthInfo(id string, dbConn *sql.DB) ([]LengthFieldInfo, error) {
	var lengths []LengthFieldInfo

	stmt, err := dbConn.Prepare(length_sql)
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
		var len LengthFieldInfo
		if err := rows.Scan(&len.LengthDataIndex, &len.LengthDetailIndex, &len.LengthFieldIndex,
			&len.DataIndex, &len.DetailIndex, &len.FieldIndex, &len.DiffValue); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		lengths = append(lengths, len)
	}
	return lengths, nil
}
