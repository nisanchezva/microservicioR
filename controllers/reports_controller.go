package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nisanchezva/microservicioR/models"
	"github.com/nisanchezva/microservicioR/utils"
)

//No pongo los controladores de ruta y estación ya que se supone que son de otra base de datos

//GetReportById obtiene un reporte por su id (cualquier tipo)
//de momento solo dejo la busqueda por id

func GetReportById(w http.ResponseWriter, r *http.Request) {
	// Estructura vacia donde se gurdarán los datos
	report := models.Report{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM contacts WHERE ID = ?
	db.Find(&report, id)
	// Se comprueba que exista el registro
	if report.ID > 0 {
		// Se codifican los datos a formato JSON
		j, _ := json.Marshal(report)
		// Se envian los datos
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		// Si no existe se envia un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

// GetReports obtiene todos los reportes
func GetReports(w http.ResponseWriter, r *http.Request) {
	// Slice (array) donde se guardaran los datos
	reports := []models.Report{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Consulta a la DB - SELECT * FROM contacts
	db.Find(&reports)
	// Se codifican los datos a formato JSON
	j, _ := json.Marshal(reports)
	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
}

// PostReport guarda un nuevo reporte
func PostReport(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se gurdaran los datos del body
	report := models.Report{}
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se decodifican los datos del body a la estructura contact
	err := json.NewDecoder(r.Body).Decode(&report)

	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	//Obtengo los datos de la ruta
	route := models.Route{}
	db.Find(&route, report.RouteID)
	//Comprobamos que exista
	if route.ID > 0 {

		//No estoy pudiendo recuperar la info de []Station, así que asumo que el orden es el de la base de datos
		stations := []models.Station{}
		db.Find(&stations, "route_id = ?", route.ID)

		if !report.Type { //Este es el reporte tipo 0 (de ruta)

			//Y ahora calculamos la demora de la ruta
			difference := stations[0].Arrival.Sub(stations[0].Departure)
			report.Duration = difference
			// Se guardan los datos en la DB
			err = db.Create(&report).Error
			if err != nil {
				// Sí hay algun error al guardar los datos se devolvera un error 500
				fmt.Println(err)
				utils.SendErr(w, http.StatusInternalServerError)
				return
			}
			// Se codifica el nuevo registro y se devuelve
			j, _ := json.Marshal(report)
			utils.SendResponse(w, http.StatusCreated, j)

		} else { //Este es el reporte tipo 1 (de estacion) (Por ahora lo dejo como la demora entra la estación 1 y 2)

			//Calculamos la demora
			difference := stations[1].Arrival.Sub(stations[0].Departure)
			report.Duration = difference
			// Se guardan los datos en la DB
			err = db.Create(&report).Error
			if err != nil {
				// Sí hay algun error al guardar los datos se devolvera un error 500
				fmt.Println(err)
				utils.SendErr(w, http.StatusInternalServerError)
				return
			}
			// Se codifica el nuevo registro y se devuelve
			j, _ := json.Marshal(report)
			utils.SendResponse(w, http.StatusCreated, j)

		}
	} else {
		// Si la id no existe se envia un error 404
		utils.SendErr(w, http.StatusNotFound)
	}

}

// UpdateReport modifica los datos de un reporte por su ID
func UpdateReport(w http.ResponseWriter, r *http.Request) {
	// Estructuras donde se almacenaran los datos
	reportFind := models.Report{}
	reportData := models.Report{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se buscan los datos
	db.Find(&reportFind, id)
	if reportFind.ID > 0 {
		// Si existe el registro se decodifican los datos del body
		err := json.NewDecoder(r.Body).Decode(&reportData)
		if err != nil {
			// Sí hay algun error en los datos se devolvera un error 400
			utils.SendErr(w, http.StatusBadRequest)
			return
		}
		// Se modifican los datos
		db.Model(&reportFind).Updates(reportData)
		// Se codifica el registro modificado y se devuelve
		j, _ := json.Marshal(reportFind)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

// DeleteReport elimina un contacto por ID
func DeleteReport(w http.ResponseWriter, r *http.Request) {
	// Estructura donde se guardara el registo buscado
	report := models.Report{}
	// Se obtiene el parametro id de la URL
	id := mux.Vars(r)["id"]
	// Conexión a la DB
	db := utils.GetConnection()
	defer db.Close()
	// Se busca el contacto
	db.Find(&report, id)
	if report.ID > 0 {
		// Sí existe, se borra y se envia contenido vacio
		db.Delete(report)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}
