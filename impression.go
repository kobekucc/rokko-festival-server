package main

import "github.com/jinzhu/gorm"

type Impression struct {
	gorm.Model
	Type string
	Comment string
}
