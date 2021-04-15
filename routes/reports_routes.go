package routes

import (
	"github.com/gorilla/mux"
	"github.com/nisanchezva/microservicioR/controllers"
)

// SetContactsRoutes agrega las rutas de contactos
func SetContactsRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/reports/{id}", controllers.GetReportById).Methods("GET")
	subRouter.HandleFunc("/reports", controllers.GetReports).Methods("GET")
	subRouter.HandleFunc("/reports", controllers.PostReport).Methods("POST")
	subRouter.HandleFunc("/reports/{id}", controllers.UpdateReport).Methods("PUT")
	subRouter.HandleFunc("/reports/{id}", controllers.DeleteReport).Methods("DELETE")
}
