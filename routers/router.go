package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Links
	router.NewRoute()

	return router
}
