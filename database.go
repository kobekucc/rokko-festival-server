package main

import (
  
    "os"
   

    "github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
    var err error
    var datasource string
    
    if os.Getenv("DATABASE_URL") != "" {
        // Heroku用
datasource = "bc70a65d138ae3:14ea8170@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_6a3970f299d4fe5?parseTime=true&charset=utf8&loc=Local"
    } else {
        // ローカル用
        datasource = "bc70a65d138ae3:14ea8170@/localhost?parseTime=true&charset=utf8&loc=Local"
    }
    db, err = gorm.Open("mysql",datasource )
    if err != nil {
      
        panic("failed to connect database")
    }

    db.AutoMigrate(&Product{})
    db.AutoMigrate(&Comment{})
}