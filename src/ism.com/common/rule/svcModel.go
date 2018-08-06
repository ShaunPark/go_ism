package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type SvcModel struct {
	Id           string `json:"id"`
	RtnType      string `json:"rtnType"`
	RtnIndex     int    `json:"rtnIdx"`
	IsConcurrent string `json:"isConcurrent"`

	Services []ServiceEntity `json:"svcs"`
}

type ServiceEntity struct {
	SvcId              string `json:"svcId"`
	SyncType           string `json:"syncType"`
	RoutingMatchMethod string `json:"rMatchMtd"`
	EntityOrder        int    `json:"eOrder"`
	ErrMapId           string `json:"errMap"`
	SvcType            string `json:"svcType"`

	Routes []ServiceRoute `json:"routes"`
	Inputs []RouteInput   `json:"rInputs"`
}

type ServiceRoute struct {
	TgtSvcIdx    int    `json:"tgtSvcIdx"`
	MappingId    string `json:"mapId"`
	IsDefault    string `json:"isDefault"`
	RoutePattern string `json:"routPtn"`
}

type RouteInput struct {
	SrcDataIndex   int    `json:"srcDataIdx"`
	SrcDetailIndex int    `json:"srcDetailIdx"`
	SrcFieldIndex  int    `json:"srcFldIdx"`
	Value          string `json:"value"`
	FieldId        string `json:"fldId"`
	Path           string `json:"path"`
}

func GetServiceModel(id string) (string, error) {
	var svc SvcModel
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(sModel_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&svc.Id, &svc.RtnType, &svc.RtnIndex, &svc.IsConcurrent)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	if svc.Services, err = getServiceEntities(id, dbConn); err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	b, err := json.Marshal(svc)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getServiceEntities(id string, dbConn *sql.DB) ([]ServiceEntity, error) {
	var svcs []ServiceEntity

	stmt, err := dbConn.Prepare(svcE_sql)
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
		var svc ServiceEntity
		if err := rows.Scan(&svc.SvcId, &svc.SyncType, &svc.RoutingMatchMethod, &svc.EntityOrder, &svc.ErrMapId, &svc.SvcType); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		if svc.Routes, err = getServiceRoutes(id, dbConn); err != nil {
			fmt.Printf("Error: %s", err)
			return nil, err
		}
		if svc.Inputs, err = getRouteInputs(id, dbConn); err != nil {
			fmt.Printf("Error: %s", err)
			return nil, err
		}

		svcs = append(svcs, svc)
	}
	return svcs, nil
}

func getServiceRoutes(id string, dbConn *sql.DB) ([]ServiceRoute, error) {
	var svcs []ServiceRoute

	stmt, err := dbConn.Prepare(svcR_sql)
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
		var svc ServiceRoute
		if err := rows.Scan(&svc.TgtSvcIdx, &svc.MappingId, &svc.IsDefault, &svc.RoutePattern); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		svcs = append(svcs, svc)
	}
	return svcs, nil
}

func getRouteInputs(id string, dbConn *sql.DB) ([]RouteInput, error) {
	var inputs []RouteInput

	stmt, err := dbConn.Prepare(rInput_sql)
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
		var input RouteInput
		if err := rows.Scan(&input.SrcDataIndex, &input.SrcDetailIndex, &input.SrcFieldIndex,
			&input.Value, &input.FieldId, &input.Path); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		inputs = append(inputs, input)
	}
	return inputs, nil
}
