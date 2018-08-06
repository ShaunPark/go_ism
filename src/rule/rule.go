package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"ism.com/common/rediscache"
	"ism.com/common/rule"
	"ism.com/common/rule/rmgr"
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
	rTypeStr := req.URL.Query().Get("rt")
	if rTypeStr == "" {
		println("empty rule type")
	}
	rId := req.URL.Query().Get("rid")
	if rId == "" {
		println("empty rule id")
	}
	var rType int
	var err error
	if rType, err = strconv.Atoi(rTypeStr); err != nil {
		panic(err)
	}

	w.Write([]byte(reload(rmgr.RuleType(rType), rId)))
}

func (h *httpGetHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	rTypeStr := req.URL.Query().Get("rt")
	if rTypeStr == "" {
		println("empty rule type")
	}
	rId := req.URL.Query().Get("rid")
	if rId == "" {
		println("empty rule id")
	}
	var rType int
	var err error
	if rType, err = strconv.Atoi(rTypeStr); err != nil {
		panic(err)
	}

	w.Write([]byte(getRule(rmgr.RuleType(rType), rId)))
}

func reload(rType rmgr.RuleType, rId string) string {
	key := fmt.Sprint(rType, ":", rId)
	val := getRuleFromDB(rType, rId)
	if val == "" {
		fmt.Println("interface ", key, " does not exist!!")
	} else {
		rediscache.Set(key, val)
	}

	return val
}

func getRule(rType rmgr.RuleType, rId string) string {
	return getFromRedis(rType, rId)
}

func getFromRedis(rType rmgr.RuleType, rId string) string {

	key := fmt.Sprint(rType, ":", rId)

	val, err := rediscache.Get(key)

	if err == redis.Nil {
		fmt.Println(key, " does not exist")

		// not exist. get from db and set to redis
		val = getRuleFromDB(rType, rId)
		if val == "" {
			fmt.Println("interface ", key, " does not exist!!")
		} else {
			rediscache.Set(key, val)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(key)
	}

	return val
}

func getRuleFromDB(rType rmgr.RuleType, rId string) string {
	var retVal string
	var err error
	switch rType {
	case rmgr.Interface:
		retVal, err = rule.GetInterface(rId)
	case rmgr.SerivceModel:
		retVal, err = rule.GetServiceModel(rId)
	case rmgr.DataStructure:
		retVal, err = rule.GetDataStructure(rId)
	case rmgr.FieldGroup:
		retVal, err = rule.GetFieldGroup(rId)
	case rmgr.Field:
		retVal, err = rule.GetField(rId)
	}
	if err != nil {
		panic(err.Error())
	}
	return retVal
}
