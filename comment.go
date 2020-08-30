package main

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	WorkId uint
	Comment  string

}