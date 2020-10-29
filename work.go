package main

import "github.com/jinzhu/gorm"

type Work struct {
	gorm.Model
	Type   string
	WorkId uint
	Vote   uint
}
