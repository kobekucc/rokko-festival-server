package main

import "github.com/jinzhu/gorm"

type Questionnaire struct {
	gorm.Model
	Age string
	Gender string
	Rate  uint
	Opinion string
}
