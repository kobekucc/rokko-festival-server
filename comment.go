package main

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Type    string
	WorkId  uint
	Comment string
}
