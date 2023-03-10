package handlers

import (
	"net/http"
	"practice/cors"

	"github.com/gorilla/mux"
)

/*GetRouter funcion encargada de retornar todas las rutas del api rest*/
func GetRoutes() http.Handler {
	routes := mux.NewRouter()
	cors.EnableCORS(routes)
	routes.HandleFunc("/", welcome).Methods(http.MethodGet)
	routes.HandleFunc("/persons", getPerson).Methods(http.MethodGet)
	routes.HandleFunc("/persons", postPerson).Methods(http.MethodPost)
	routes.HandleFunc("/api/report", rephandler).Methods(http.MethodGet)

	return routes
}
