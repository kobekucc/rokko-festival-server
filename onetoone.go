package main

import "github.com/jinzhu/gorm"

type Onetoone struct {
	gorm.Model
	Order uint
	NumberOfTimes uint
	Name string
	SwitchName string
	Done uint
}
