package main

import (
	"fmt"
	sql "database/sql"
	_ "github.com/mattn/go-sqlite3"
	http "net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"errors"
	routes "routes"
)

func log(message string) {
	fmt.Printf("%v: %v\n", time.Now().String(), message)
}

type databaseConfig struct {
	host string
	port int
	username string
	password string
	database string
	maxConn int
	minConn int
}

func makeConfig() (*databaseConfig, error) {
	var config *databaseConfig
	config = new(databaseConfig)
	var rawConfig map[string]interface{}

	content, err := ioutil.ReadFile("./config/database.json")
	if err != nil {
		log("Unable to read config/database.json")
		return config, err
	}

	err = json.Unmarshal(content, &rawConfig)
	if err != nil {
		log("Unable to parse database.json")
		return config, err
	}

	if val, ok := rawConfig["host"]; ok {
		config.host, ok = val.(string)
		if !ok { return config, errors.New("Host is invalid") }
	} else {
		return config, errors.New("No host in config")
	}

	if val, ok := rawConfig["port"]; ok {
		var port float64
		port, ok = val.(float64)
		if !ok { return config, errors.New("Port is invalid") }
		config.port = int(port)
	} else {
		return config, errors.New("No port in config")
	}

	if val, ok := rawConfig["username"]; ok {
		config.username, ok = val.(string)
		if !ok { return config, errors.New("Username is invalid") }
	} else {
		return config, errors.New("No username in config")
	}

	if val, ok := rawConfig["password"]; ok {
		config.password, ok = val.(string)
		if !ok { return config, errors.New("Password is invalid") }
	} else {
		return config, errors.New("No password in config")
	}

	if val, ok := rawConfig["database"]; ok {
		config.database, ok = val.(string)
		if !ok { return config, errors.New("database is invalid") }
	} else {
		return config, errors.New("No database in config")
	}

	if val, ok := rawConfig["maxConnection"]; ok {
		config.maxConn, ok = val.(int)
		if !ok { config.maxConn = 20 }
	} else {
		config.maxConn = 20
	}

	if val, ok := rawConfig["minConnection"]; ok {
		config.minConn, ok = val.(int)
		if !ok { config.minConn = 10 }
	} else {
		config.minConn = 10
	}

	return config, nil

}

func main() {
	
	config, err := makeConfig()
	if err != nil {
		log("Unable to get DB configuration")
		log(err.Error())
		return
	}
	fmt.Println("CONFIG")
	fmt.Println(config)
	dbStr := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		config.username,
		config.password,
		config.host,
		config.port,
		config.database,
	)

	fmt.Println(dbStr)
	db, _ := sql.Open("sqlite3", dbStr)
	err = db.Ping()
	if err != nil {
		log("Unable to connect to database")
		log(err.Error())
		return
	}

	defer db.Close()

	db.SetMaxIdleConns(config.minConn)
	db.SetMaxOpenConns(config.maxConn)
}