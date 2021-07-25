package db

import (
	"gin-scaffold/util"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConn
var Conn *gorm.DB

// init
func init() {
	if runMode := os.Getenv("RUN_MODE"); runMode == "testing" {
		// TO-DO
	} else {
		DBInit()
	}
}

// Run mode init
func DBConfigInit() (string, string) {
	util.InitConfig()
	// get db config
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = viper.GetString("DB_HOST")
	}

	// get db port
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = viper.GetString("DB_PORT")
	}

	// get db name
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = viper.GetString("DB_NAME")
	}

	// get db user
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = viper.GetString("DB_USER")
	}

	// get db passwd
	dbPasswd := os.Getenv("DB_PASSWD")
	if dbPasswd == "" {
		dbPasswd = viper.GetString("DB_PASSWD")
	}

	log.Println("db connecting ...", dbName, dbHost)

	// db dsn
	dsn := dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn, dbName
}

// Run mode init
func DBInit() {
	//db config init
	dsn, dbName := DBConfigInit()
	// db dsn
	var err error
	// set db log level
	logLevel := logger.Info
	dbLogMode := os.Getenv("DB_LOGMODE")
	if dbLogMode == "" {
		dbLogMode = viper.GetString("DB_LOGMODE")
		if dbLogMode == "Warn" {
			logLevel = logger.Warn
		}
	}
	//db logger
	DBLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             3 * time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,        // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,           // Disable color
		},
	)

	//get conn
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: DBLogger,
	})

	//connect err
	if err != nil {
		panic(err.Error())
	}

	//check db
	status := CheckDB(Conn, dbName)
	if !status {
		return
	}

	//sql db
	sqlDB, errSql := Conn.DB()
	//err sql
	if errSql != nil {
		panic(" seting database error !")
	}
	// db setting
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

//check db
func CheckDB(Conn *gorm.DB, dbName string) bool {
	checkDB := "select * from information_schema.SCHEMATA where SCHEMA_NAME = '" + dbName + "'; "
	tx := Conn.Raw(checkDB)
	rows, _ := tx.Rows()
	defer rows.Close()
	checkDBError := tx.Error
	if checkDBError != nil {
		log.Println(checkDBError.Error())
		return false
	}

	if !rows.Next() {
		Conn.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;")
	}

	//use db
	useError := Conn.Exec("use " + dbName).Error
	if useError != nil {
		log.Println(useError.Error())
		return false
	}

	return true
}
