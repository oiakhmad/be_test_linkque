package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Database Database
type Database struct {
	Host      string
	User      string
	Password  string
	DBName    string
	DBNumber  int
	Port      int
	DebugMode bool
	SSL       bool
}

// LoadDBConfig load database configuration
func LoadDBConfig(name string) Database {
	dbName := fmt.Sprintf("DATABASE_%s_", strings.ToUpper(name))
	port, err := strconv.Atoi(os.Getenv(dbName + "PORT"))
	if err != nil {
		log.Println(err)
		log.Fatal(fmt.Sprintf("failed to convert %sPORT from string to int", dbName))
	}

	conf := Database{
		Host:      os.Getenv(dbName + "HOST"),
		User:      os.Getenv(dbName + "USER"),
		Password:  os.Getenv(dbName + "PASSWORD"),
		DBName:    os.Getenv(dbName + "NAME"),
		Port:      port,
		SSL:       os.Getenv(dbName+"SSL") == "true",
		DebugMode: os.Getenv(dbName+"DEBUG") == "true",
	}
	return conf
}
