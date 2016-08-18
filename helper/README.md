## Go helper

Go useful functions. This library uses github.com/jinzhu/gorm library

## Install:
```
go get github.com/kroksys/go/helper

import (
    "github.com/kroksys/go/helper"
)
```

## Functions:

AddLogFile(fileName string) error
> Set "log" output to console and file

HandleErrorFatal(err error)
> Error handling function using log.Fatalln

HandleErrorPanic(err error)
> Error handling function using log.Panicln

HandleError(err error) bool
> Error handling function using log.Println and returning boolean (true if error != nil)

OpenDatabaseMysqlGorm(username, password, dbName string) *gorm.DB
> Open MySql database connection using "github.com/jinzhu/gorm"

OpenDatabaseMysql(username, password, dbName string) *sql.DB
> Open MySql database connection using build in sql library

StringToTimeDDMMYYYY(str string) (time.Time, error)
> Convert sting formatted as DD-MM-YYYY to time

ToString(object interface{}) string
> 
