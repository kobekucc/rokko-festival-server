package main

import "github.com/jinzhu/gorm"

type Onetoone struct {
	gorm.Model
	NumberOfTimes uint
	Name string
	SwitchName string
}
