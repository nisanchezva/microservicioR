package models

import (
	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	Stations []Station
}
