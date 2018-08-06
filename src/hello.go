package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"message"
	"myhttp"
	"net/http"
	"os"

	"ism.com/common/rule/rmgr"
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
		rule := rmgr.GetInterface(infId)
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
