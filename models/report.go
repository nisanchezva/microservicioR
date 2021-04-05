package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	Date     string //Lo dejo así por problemas para recibir las fechas, manejar formato dd-mm-aaaa
	RouteID  int    //ID de la Ruta
	Route    Route
	Duration time.Duration //Tiempo total que tomo recorrer la ruta, esta en nano segundos
	Type     bool          // false = Ruta - true = Paradero

}
