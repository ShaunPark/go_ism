package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"ism.com/common/db"
	"ism.com/online/ismredis"
)

type httpGetHandler struct {
	http.Handler
}
type httpReloadHandler struct {
	http.Handler
}

func main() {
	http.Handle("/get", new(httpGetHandler))
	http.Handle("/reload", new(httpReloadHandler))
	http.ListenAndServe(":3000", nil)
}

func (h *httpReloadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rType := req.URL.Query().Get("rt")
	if rType == "" {
		println("empty rule type")
	}
	rId := req.URL.Query().Get("rid")
	if rId == "" {
		println("empty rule id")
	}
	rInfo := fmt.Sprint(rType, ":", rId)

	w.Write([]byte(reload(rInfo)))
}

func (h *httpGetHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rType := req.URL.Query().Get("rt")
	if rType == "" {
		println("empty rule type")
	}
	rId := req.URL.Query().Get("rid")
	if rId == "" {
		println("empty rule id")
	}
	rInfo := fmt.Sprint(rType, ":", rId)

	w.Write([]byte(getRule(rInfo)))
}

func reload(key string) string {
	strs := strings.Split(key, ":")
	val := getRuleFromDB(strs[1])
	if val == "" {
		fmt.Println("interface ", key, " does not exist!!")
	} else {
		ismredis.Set(key, val)
	}

	return val
}

func getRule(key string) string {
	return getFromRedis(key)
}

func getFromRedis(key string) string {

	val, err := ismredis.Get(key)

	if err == redis.Nil {
		fmt.Println(key, " does not exist")

		// not exist. get from db and set to redis
		strs := strings.Split(key, ":")
		val = getRuleFromDB(strs[1])
		if val == "" {
			fmt.Println("interface ", key, " does not exist!!")
		} else {
			ismredis.Set(key, val)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(key)
	}

	return val
}

type RuleInf struct {
	ID   string
	Name string
	Data string
}

func getRuleFromDB(key string) string {
	var rule RuleInf

	dbConn := db.GetDatabase()

	// defer the close till after the main function has finished
	// executing

	stmt, err := dbConn.Prepare("SELECT ID, NAME, DATA FROM ISM_INF where ID = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRow(key).Scan(&rule.ID, &rule.Name, &rule.Data)
	if err != nil {
		if err == sql.ErrNoRows {
			return ""
		} else {
			panic(err.Error())
		}
	}
	return rule.Data
}
