package models

import (
  "errors"
  
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
  
  db, err := gorm.Open("postgres", "host=localhost port=5432 user=javlon dbname=messaging sslmode=disable password=therichestman8")

  if err != nil {
    panic("Failed to connect to database!")
  }

  db.AutoMigrate(&User{})

  DB = db
}

func IsNotFound(row *gorm.DB) bool {

  return errors.Is(row.Error, gorm.ErrRecordNotFound)
}