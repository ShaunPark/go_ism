package rmgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/koding/cache"
	"ism.com/common/rule"
)

type RuleType int

const (
	Interface     RuleType = 0
	DataStructure RuleType = 1
	ServiceMap    RuleType = 2
	SerivceModel  RuleType = 3
	FieldGroup    RuleType = 4
	Field         RuleType = 5
	System        RuleType = 6
	Server        RuleType = 7
	Application   RuleType = 8
	Service       RuleType = 9
)

var ruleCache *cache.MemoryNoTS

func GetInterface(id string) rule.Interface {
	retStr := getRule(id, Interface)
	var rule rule.Interface
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetSerivceModel(id string) rule.SvcModel {
	retStr := getRule(id, SerivceModel)
	var rule rule.SvcModel
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetFieldGroup(id string) rule.FieldGroup {
	retStr := getRule(id, FieldGroup)
	var rule rule.FieldGroup
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetField(id string) rule.Field {
	retStr := getRule(id, Field)
	var rule rule.Field
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetDataStructure(id string) rule.DataStructure {
	retStr := getRule(id, DataStructure)
	var rule rule.DataStructure
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetServiceMap(id string) rule.ServiceMap {
	retStr := getRule(id, ServiceMap)
	var rule rule.ServiceMap
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetService(id string) rule.Service {
	retStr := getRule(id, Service)
	var rule rule.Service
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetSystem(id string) rule.System {
	retStr := getRule(id, System)
	var rule rule.System
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetServer(id string) rule.Server {
	retStr := getRule(id, Server)
	var rule rule.Server
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

func GetApplication(id string) rule.Application {
	retStr := getRule(id, Application)
	var rule rule.Application
	if err := json.Unmarshal([]byte(retStr), &rule); err != nil {
		panic(err)
	}
	return rule
}

// 로컬 cache에서 Json형태의 룰을 받아오는 함수
// 로컬 cache에 룰이 없으면 http로 룰 프로세스에서 받아옴
func getRule(infId string, infType RuleType) string {
	println(infId)

	if ruleCache == nil {
		ruleCache = cache.NewMemoryNoTS()
	}
	var retVal string
	rInfo, err := ruleCache.Get(infId)
	if err != nil {
		if err == cache.ErrNotFound {
			resp, err := http.Get(fmt.Sprint("http://localhost:3000/get?rid=", infId, "&rt=", infType))
			if err != nil {
				panic(err.Error())
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			retVal = string(body)
			ruleCache.Set(infId, retVal)
		} else {
			panic(err.Error())
		}
	} else {
		retVal = rInfo.(string)
	}

	return retVal
}
