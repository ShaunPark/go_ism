package rule

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ism.com/common/db"
)

type Interface struct {
	Id        string `json:"id"`
	Is2pc     string `json:"is2pc"`
	Msgvalid  string `json:"messageValidate"`
	InfType   string `json:"infType"`
	SvcTempId string `json:"svcTempId"`

	RealTimeSvc RealTimeService `json:"rSvc"`
	BatchSvc    BatchService    `json:"bSvc"`
	DeferredSvc DeferredService `json:"dSvc"`

	Systems []SystemEntity `json:"systems"`
}

type RealTimeService struct {
	MsgLog     string `json:"msgLog"`
	ProcLog    string `json:"procLog"`
	AsyncRCnt  int    `json:"asyncRetryCount"`
	AsyncRTOut int    `json:"asyncRetryTimeout"`
}

type BatchService struct {
	FetchCount         int        `json:"fetchCount"`
	RollbackType       string     `json:"rollbackType"`
	WaitTimeout        int        `json:"waitTimeout"`
	FileSeparatorClass NullString `json:"fileSeparatorClass"`
	PostTaskId         string     `json:"postTaskId"`
	IsDirect           string     `json:"isDirect"`
	PreTaskId          string     `json:"preTaskId"`
	CheckSize          NullInt    `json:"checkSize"`
	BinaryMode         string     `json:"binaryMode"`
}

type DeferredService struct {
	FetchInterval         int    `json:"fetchInterval"`
	DataMaxSeq            string `json:"dataMaxSeq"`
	FetchCount            int    `json:"fetchCount"`
	DayCloseFieldName     string `json:"dayCloseFieldName"`
	DayCloseFieldValue    string `json:"dayCloseFieldValue"`
	SequenceInitCondition string `json:"sequenceInitCondition"`
	PostTaskId            string `json:"postTaskId"`
	RollbackType          string `json:"rollbackType"`
	SequenceField         string `json:"sequenceField"`
	DateField             string `json:"dateField"`
	AutoGenerate          string `json:"autoGenerate"`
	GenInterval           int    `json:"genInterval"`
	PreTaskId             string `json:"preTaskId"`
	IncludeEndValue       string `json:"includeEndValue"`
	CloseHandler          string `json:"closeHandler"`
}

type SystemEntity struct {
	SystemId   string     `json:"systemId"`
	ParserName NullString `json:"parserName"`
	Order      int        `json:"order"`
	Timeout    NullInt    `json:"timeout"`
}

func GetInterface(id string) (string, error) {
	var inf Interface
	dbConn := db.GetDatabase()

	stmt, err := dbConn.Prepare(inf_sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&inf.Id, &inf.Is2pc, &inf.Msgvalid, &inf.InfType, &inf.SvcTempId)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
			err = nil
		} else {
			return "", err
		}
	}

	if inf.InfType == "O" {
		println("Online service")
		inf.RealTimeSvc, err = getRealTimeService(id, dbConn)
	} else if inf.InfType == "B" {
		println("Batch service")
		inf.BatchSvc, err = getBatchService(id, dbConn)
	} else if inf.InfType == "D" {
		println("Deferred service")
		inf.DeferredSvc, err = getDeferredService(id, dbConn)
	}

	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	if inf.Systems, err = getSystemEntities(id, dbConn); err != nil {
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

func getRealTimeService(id string, dbConn *sql.DB) (RealTimeService, error) {
	var svc RealTimeService

	stmt, err := dbConn.Prepare(rSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	err = stmt.QueryRow(id).Scan(&svc.MsgLog, &svc.ProcLog, &svc.AsyncRCnt, &svc.AsyncRTOut)
	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
		} else {
			panic(err.Error())
		}
	}
	return svc, nil
}

func getBatchService(id string, dbConn *sql.DB) (BatchService, error) {
	var svc BatchService

	stmt, err := dbConn.Prepare(bSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	err = stmt.QueryRow(id).Scan(&svc.FetchCount, &svc.RollbackType,
		&svc.WaitTimeout, &svc.FileSeparatorClass, &svc.PostTaskId,
		&svc.IsDirect, &svc.PreTaskId, &svc.CheckSize, &svc.BinaryMode)

	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
		} else {
			panic(err.Error())
		}
	}
	return svc, nil
}

func getDeferredService(id string, dbConn *sql.DB) (DeferredService, error) {
	var svc DeferredService

	stmt, err := dbConn.Prepare(dSvc_sql)
	if err != nil {
		panic(err.Error())
	}
	err = stmt.QueryRow(id).Scan(&svc.FetchInterval, &svc.DataMaxSeq,
		&svc.FetchCount, &svc.DayCloseFieldName, &svc.DayCloseFieldValue,
		&svc.SequenceInitCondition, &svc.PostTaskId, &svc.RollbackType, &svc.SequenceField,
		&svc.DateField, &svc.AutoGenerate, &svc.GenInterval, &svc.PreTaskId, &svc.IncludeEndValue, &svc.CloseHandler)

	if err != nil {
		if err == sql.ErrNoRows {
			println("not found")
		} else {
			panic(err.Error())
		}
	}
	return svc, nil
}

func getSystemEntities(id string, dbConn *sql.DB) ([]SystemEntity, error) {
	var sEntities []SystemEntity

	stmt, err := dbConn.Prepare(sysEntity_sql)
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
		var sEntity SystemEntity
		if err := rows.Scan(&sEntity.SystemId, &sEntity.ParserName, &sEntity.Order, &sEntity.Timeout); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		sEntities = append(sEntities, sEntity)
	}
	return sEntities, nil
}
