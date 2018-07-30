package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"message"
	"myhttp"
	"net/http"
	"os"

	"github.com/koding/cache"
	"ism.com/online/gid"

	"gopkg.in/yaml.v2"
)

type listeners struct {
	Listeners []listenInfo
}
type listenInfo struct {
	Port     int
	Url      string
	Msgparam string
}

type httpHandler struct {
	http.Handler
}

func main() {
	finish := make(chan bool)

	lInfos := new(listeners)
	lInfos = lInfos.loadListenerInfo()
	myParser = getParser()
	myGidChecker = getGidChecker()
	for _, li := range lInfos.Listeners {
		l := listenInfo(li)
		println(l.Url)
		go func() {
			http.Handle(l.Url, new(httpHandler))
			http.ListenAndServe(fmt.Sprintf(":%d", l.Port), nil)
		}()
	}

	<-finish
}

var myParser message.Parser
var myGidChecker gid.GidChecker

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var returnMsg string
	returnMsg = "Your Request Path is " + req.URL.Path
	println(fmt.Sprint("Requests from ", req.RemoteAddr))

	param1 := req.URL.Query().Get("data")
	if param1 != "" {
		gid := myParser.GetGID(param1)
		infId := myParser.GetInterfaceId(param1)
		rule := getRule(infId, "INF")
		println(rule)
		if myGidChecker.CheckGID(gid) {
			httpCli := new(myhttp.MyhttpClient)
			returnMsg = httpCli.Call()
		} else {
			returnMsg = "GID already exist!"
		}
	}

	println(returnMsg)
	w.Write([]byte(returnMsg))
}

var ruleCache *cache.MemoryNoTS

func getRule(infId string, infType string) string {
	println(infId)

	if ruleCache == nil {
		ruleCache = cache.NewMemoryNoTS()
	}
	var retVal string
	rInfo, err := ruleCache.Get(infId)
	if err != nil {
		if err == cache.ErrNotFound {
			resp, err := http.Get(fmt.Sprint("http://localhost:3000/get?rid=", infId, "&rt=INF"))
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

func (lInfo *listeners) loadListenerInfo() *listeners {
	path := os.Getenv("GO_PROP")
	if path == "" {
		path = "."
	}
	yamlFile, err := ioutil.ReadFile(fmt.Sprint(path + "/listener.yml"))
	if err != nil {
		log.Printf("Read config file error #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, lInfo)
	if err != nil {
		log.Printf("Unmarshal config file error #%v", err)
	}
	return lInfo
}

func getParser() message.Parser {
	return new(message.MyParser)
}

func getGidChecker() gid.GidChecker {
	return new(gid.MySQLGIDChecker)
}
