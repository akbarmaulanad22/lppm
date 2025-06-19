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
	PKMPDPPController *PKMPDPPController
=======
	ProfilSODTController     *ProfilSODTController
	PKMRDRPController        *PKMRDRPController
	HKIMHSController         *HKIMHSController
>>>>>>> 8cb079c3c0f397c989d3ad4c7e9cbf477793c8cc
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

    pkmRouter.HandleFunc("/pdpp", route.PKMPDPPController.Create).Methods("POST")
    pkmRouter.HandleFunc("/pdpp", route.PKMPDPPController.List).Methods("GET")
    pkmRouter.HandleFunc("/pdpp/{id}", route.PKMPDPPController.Update).Methods("PUT")
    pkmRouter.HandleFunc("/pdpp/{id}", route.PKMPDPPController.Delete).Methods("DELETE")
}





=======
	profilRouter.HandleFunc("/sodt", route.ProfilSODTController.Create).Methods("POST")
	profilRouter.HandleFunc("/sodt", route.ProfilSODTController.List).Methods("GET")
	profilRouter.HandleFunc("/sodt/{id}", route.ProfilSODTController.Update).Methods("PUT")
	profilRouter.HandleFunc("/sodt/{id}", route.ProfilSODTController.Delete).Methods("DELETE")

	pkmRouter := route.Router.PathPrefix("/pkm").Subrouter()
	pkmRouter.HandleFunc("/rdrp", route.PKMRDRPController.Create).Methods("POST")
	pkmRouter.HandleFunc("/rdrp", route.PKMRDRPController.List).Methods("GET")
	pkmRouter.HandleFunc("/rdrp/{id}", route.PKMRDRPController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/rdrp/{id}", route.PKMRDRPController.Delete).Methods("DELETE")

	hkiRouter := route.Router.PathPrefix("/hki").Subrouter()
	hkiRouter.HandleFunc("/mhs", route.HKIMHSController.Create).Methods("POST")
	hkiRouter.HandleFunc("/mhs", route.HKIMHSController.List).Methods("GET")
	hkiRouter.HandleFunc("/mhs/{id}", route.HKIMHSController.Update).Methods("PUT")
	hkiRouter.HandleFunc("/mhs/{id}", route.HKIMHSController.Delete).Methods("DELETE")

}

>>>>>>> 8cb079c3c0f397c989d3ad4c7e9cbf477793c8cc
func (route *RouteConfig) SetupAuthRoute() {
}
