package helper

import (
	"os"
	"log"
	"io"
	"database/sql"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	"errors"
)

// SetOutput sets the output destination for the logger to Stdout and "File".
// File will be created or read in project working directory
func AddLogFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	return nil
}

// Error handling function wrapped because of its
// frequent usage. os.Exit(1) on error!
func HandleErrorFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Error handling function wrapped because of its
// frequent usage. Continue working after Log error
func HandleErrorPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

// HandleError check if error not nil and print error on Log
// Return true or false.
func HandleError(err error) bool {
	if err != nil {
		log.Println(err)
	}
	return err != nil
}

// Open MySql database connection using "github.com/jinzhu/gorm"
// Using build in error handling: Exit on error.
func OpenDatabaseMysqlGorm(username, password, dbName string) *gorm.DB {
	db, err := gorm.Open("mysql", username + ":" + password + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	HandleErrorFatal(err)
	return db
}

// Open MySql database connection using guild in sql library
// Using build in error handling: Exit on error.
func OpenDatabaseMysql(username, password, dbName string) *sql.DB {
	db, err := sql.Open("mysql", username + ":" + password + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	HandleErrorFatal(err)
	return db
}

// Convert sting formatted as DD-MM-YYYY to time
// When i used time.Parse it is forced to use Month as first parameter.
func StringToTimeDDMMYYYY(str string) (time.Time, error) {
	i := strings.Split(str, string(str[2]))
	if len(i) != 3 {
		return time.Time{}, errors.New("Wrong parameter at helper.StringToTimeDDMMYYYY(str string): " + str)
	}
	return time.Parse("01-02-2006", i[1] + "-" + i[0] + "-" + i[2])
}
