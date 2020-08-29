package main

import (
   
    "os"
    "log"

    "github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
    var err error
    var datasource string

    // DBMS := "mysql"
	// USER := "b6269a7b123b6c"
	// PASS := "c2129dd2"
	// PROTOCOL := "tcp(localhost:3306)"
    // DBNAME := "heroku_4cadff22e8acc17"

    // CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

    
    if os.Getenv("DATABASE_URL") != "" {
        // Heroku用
datasource = "bc70a65d138ae3:14ea8170@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_6a3970f299d4fe5?parseTime=true&charset=utf8&loc=Local"
    } else {
        // ローカル用
        datasource = "bc70a65d138ae3:14ea8170@/localhost?parseTime=true&charset=utf8&loc=Local"
    }
    db, err = gorm.Open("mysql",datasource )
    if err != nil {
      
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }

    db.AutoMigrate(&Product{})
}