package route

import (
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	// router
	Router *mux.Router

	// all field controller
	ProfilVisiMisiController *ProfilVisiMisiController
<<<<<<< HEAD
	PKMRDRPController *PKMRDRPController
=======
>>>>>>> 006ad76c5e76579d26dc36a60676ed24113c225b
}

func (route *RouteConfig) Setup() {
	route.SetupGuestRoute()
	route.SetupAuthRoute()
}

func (route *RouteConfig) SetupGuestRoute() {
	profilRouter := route.Router.PathPrefix("/profil").Subrouter()
	profilRouter.HandleFunc("/visi-misi", route.ProfilVisiMisiController.Create).Methods("POST")
	profilRouter.HandleFunc("/visi-misi", route.ProfilVisiMisiController.List).Methods("GET")
	profilRouter.HandleFunc("/visi-misi/{id}", route.ProfilVisiMisiController.Update).Methods("PUT")
	profilRouter.HandleFunc("/visi-misi/{id}", route.ProfilVisiMisiController.Delete).Methods("DELETE")

<<<<<<< HEAD
	pkmRouter := route.Router.PathPrefix("/pkm").Subrouter()
    pkmRouter.HandleFunc("/rdrp", route.PKMRDRPController.Create).Methods("POST")
    pkmRouter.HandleFunc("/rdrp", route.PKMRDRPController.List).Methods("GET")
    pkmRouter.HandleFunc("/rdrp/{id}", route.PKMRDRPController.Update).Methods("PUT")
    pkmRouter.HandleFunc("/rdrp/{id}", route.PKMRDRPController.Delete).Methods("DELETE")
}





=======
}

>>>>>>> 006ad76c5e76579d26dc36a60676ed24113c225b
func (route *RouteConfig) SetupAuthRoute() {
}
