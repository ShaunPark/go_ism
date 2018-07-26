package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
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
	if client == nil {
		createRedisClient()
	}

	strs := strings.Split(key, ":")
	val := getRuleFromDB(strs[1])
	if val == "" {
		fmt.Println("interface ", key, " does not exist!!")
	} else {
		client.Set(key, val, 0)
	}

	return val
}
func getRule(key string) string {
	return getFromRedis(key)
}

var client *redis.Client

func createRedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func getFromRedis(key string) string {
	if client == nil {
		createRedisClient()
	}

	val, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key, " does not exist")

		// not exist. get from db and set to redis
		strs := strings.Split(key, ":")
		val = getRuleFromDB(strs[1])
		if val == "" {
			fmt.Println("interface ", key, " does not exist!!")
		} else {
			client.Set(key, val, 0)
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

var dbConn *sql.DB

func getRuleFromDB(key string) string {
	var rule RuleInf
	var err error
	if dbConn == nil {
		dbConn, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/go_ism")
		// if there is an error opening the connection, handle it
		if err != nil {
			panic(err.Error())
		}

		dbConn.Ping()
		dbConn.SetMaxIdleConns(5)
		dbConn.SetMaxOpenConns(100)
	}

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
