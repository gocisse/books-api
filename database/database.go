package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//declare a var to point to our DB connection

var (
	DBConn *gorm.DB
)
