package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB pointer gorm db
var (
	MySQL *gorm.DB
)

// Database set database vars
type DatabaseMysql struct {
	Host              string
	User              string
	Password          string
	DBName            string
	Port              string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

// LoadDBConfig load database configuration
func LoadMySQLConfig() DatabaseMysql {
	dbDebug, _ := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	conf := DatabaseMysql{
		Host:      os.Getenv("DB_HOST"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		Port:      os.Getenv("DB_PORT"),
		DebugMode: dbDebug,
	}

	return conf
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MySQLConnect() *gorm.DB {
	dbConf := LoadMySQLConfig()
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=latin1&parseTime=True&loc=Local`, dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Enable Color
		},
	)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	if dbConf.DebugMode {
		return connection.Debug()
	}

	return connection
}

// OpenDbPool open connection
func OpenMySQLPool() {
	MySQL = MySQLConnect()
	sqlDB, err := MySQL.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
