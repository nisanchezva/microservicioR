package utils

import (
	"fmt"
	"time"

	"github.com/nisanchezva/microservicioR/models"
)

// MigrateDB migra la base de datos
func MigrateDB() {
	db := GetConnection()
	defer db.Close()
	fmt.Println("Migrating models....")
	//Elimino las tablas anteriores si ya estaban
	db.DropTableIfExists(&models.Route{}, &models.Station{})
	// Automigrate se encarga de migrar la base de datos sí no se ha migrado, y lo hace a partir del modelo
	db.AutoMigrate(&models.Route{}, &models.Station{})

	//Agregando datos
	//Con solo agregar el de ruta, grom agrega los datos de estación, así se mantiene la integridad, pero si llegara a ocurrir un error, alguna id deberia estar en 0
	//De momento dejo estos datos quemados pero en cuanto se implemente el recorrido del bus pues se actualizan las horas cuando se llegue a cada estación

	route := models.Route{}
	stations := []models.Station{
		models.Station{
			Arrival:   time.Date(2021, time.April, 5, 14, 10, 30, 0, time.Local),
			Departure: time.Date(2021, time.April, 5, 13, 1, 13, 0, time.Local),
			Route:     route,
		},
		models.Station{
			Arrival:   time.Date(2021, time.April, 5, 13, 20, 0, 0, time.Local),
			Departure: time.Date(2021, time.April, 5, 13, 21, 0, 0, time.Local),
			Route:     route,
		},
		models.Station{
			Arrival:   time.Date(2021, time.April, 5, 13, 53, 0, 0, time.Local),
			Departure: time.Date(2021, time.April, 5, 13, 54, 0, 0, time.Local),
			Route:     route,
		},
		/*La interpretación es que para la primera estación ya se completo una ruta, por eso en arrival
		es mas tarde que en departure (se genera el reporte antes de actuailizar departure), pero para el resto de estaciones esto aun no pasa
		*/
	}

	route.Stations = stations
	db.Create(&route)

}
