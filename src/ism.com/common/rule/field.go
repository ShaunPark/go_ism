package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type FieldGroup struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	FieldDelimeter string `json:"fDelimeter"`

	Fields []FieldMap `json:"fields"`
}

type FieldMap struct {
	Fieldindex      int        `json:"fIndex"`
	Fieldid         string     `json:"fldId"`
	Fieldoffset     int        `json:"fldOffset"`
	Diffvalue       int        `json:"diffValue"`
	Iskey           string     `json:"isKey"`
	Isnull          string     `json:"isNull"`
	Issqlfunction   string     `json:"isSqlFunction"`
	LengthFieldType string     `json:"lendthFldType"`
	InOutType       NullString `json:"inoutType"`
	FilterType      NullString `json:"filterType"`
}

type Field struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	FieldType   string `json:"fldType"`
	FieldLength int    `json:"fldLength"`
	Fieldformat string `json:"fldFormat"`
	Fillchar    string `json:"fillChar"`
	Aligntype   string `json:"alignType"`
}

func GetField(id string) (string, error) {
	var fld Field
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(fld_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&fld.Id, &fld.Name, &fld.FieldType, &fld.FieldLength,
		&fld.Fieldformat, &fld.Fillchar, &fld.Aligntype)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	b, err := json.Marshal(fld)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func GetFieldGroup(id string) (string, error) {
	var flg FieldGroup
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(flg_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&flg.Id, &flg.Name, &flg.FieldDelimeter)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	if flg.Fields, err = getFieldGroupMap(id, dbConn); err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	b, err := json.Marshal(flg)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getFieldGroupMap(id string, dbConn *sql.DB) ([]FieldMap, error) {
	var fMaps []FieldMap

	stmt, err := dbConn.Prepare(flgMap_sql)
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
		var fMap FieldMap
		if err := rows.Scan(&fMap.Fieldindex, &fMap.Fieldid, &fMap.Fieldoffset,
			&fMap.Diffvalue, &fMap.Iskey, &fMap.Isnull, &fMap.Issqlfunction,
			&fMap.LengthFieldType, &fMap.InOutType, &fMap.FilterType); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fMaps = append(fMaps, fMap)
	}
	return fMaps, nil
}
