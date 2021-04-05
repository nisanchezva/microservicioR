package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Station struct {
	gorm.Model
	Arrival   time.Time
	Departure time.Time
	RouteID   uint
	Route     Route
}
