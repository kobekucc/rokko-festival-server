package main

import "github.com/jinzhu/gorm"

type Work struct {
	gorm.Model
	WorkId uint
	Name  string
	Vote uint
}