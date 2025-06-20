package route

import (
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	// router
	Router *mux.Router

	// all field controller
	ProfilVisiMisiController *ProfilVisiMisiController
	ProfilSODTController     *ProfilSODTController
	PKMRDRPController        *PKMRDRPController
	PKMPDPPController        *PKMPDPPController
	PKMTCRController         *PKMTCRController
	HKIMHSController         *HKIMHSController
	HKIDosenController       *HKIDosenController
	PenelitianRDRPController *PenelitianRDRPController
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

	profilRouter.HandleFunc("/sodt", route.ProfilSODTController.Create).Methods("POST")
	profilRouter.HandleFunc("/sodt", route.ProfilSODTController.List).Methods("GET")
	profilRouter.HandleFunc("/sodt/{id}", route.ProfilSODTController.Update).Methods("PUT")
	profilRouter.HandleFunc("/sodt/{id}", route.ProfilSODTController.Delete).Methods("DELETE")

	pkmRouter := route.Router.PathPrefix("/pkm").Subrouter()
	pkmRouter.HandleFunc("/rdrp", route.PKMRDRPController.Create).Methods("POST")
	pkmRouter.HandleFunc("/rdrp", route.PKMRDRPController.List).Methods("GET")
	pkmRouter.HandleFunc("/rdrp/{id}", route.PKMRDRPController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/rdrp/{id}", route.PKMRDRPController.Delete).Methods("DELETE")

	pkmRouter.HandleFunc("/pdpp", route.PKMPDPPController.Create).Methods("POST")
	pkmRouter.HandleFunc("/pdpp", route.PKMPDPPController.List).Methods("GET")
	pkmRouter.HandleFunc("/pdpp/{id}", route.PKMPDPPController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/pdpp/{id}", route.PKMPDPPController.Delete).Methods("DELETE")

	pkmRouter.HandleFunc("/tcr", route.PKMTCRController.Create).Methods("POST")
	pkmRouter.HandleFunc("/tcr", route.PKMTCRController.List).Methods("GET")
	pkmRouter.HandleFunc("/tcr/{id}", route.PKMTCRController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/tcr/{id}", route.PKMTCRController.Delete).Methods("DELETE")

	hkiRouter := route.Router.PathPrefix("/hki").Subrouter()
	hkiRouter.HandleFunc("/mhs", route.HKIMHSController.Create).Methods("POST")
	hkiRouter.HandleFunc("/mhs", route.HKIMHSController.List).Methods("GET")
	hkiRouter.HandleFunc("/mhs/{id}", route.HKIMHSController.Update).Methods("PUT")
	hkiRouter.HandleFunc("/mhs/{id}", route.HKIMHSController.Delete).Methods("DELETE")

	hkiRouter.HandleFunc("/dosen", route.HKIDosenController.Create).Methods("POST")
	hkiRouter.HandleFunc("/dosen", route.HKIDosenController.List).Methods("GET")
	hkiRouter.HandleFunc("/dosen/{id}", route.HKIDosenController.Update).Methods("PUT")
	hkiRouter.HandleFunc("/dosen/{id}", route.HKIDosenController.Delete).Methods("DELETE")

	penelitianRouter := route.Router.PathPrefix("/penelitian").Subrouter()
	penelitianRouter.HandleFunc("/rdrp", route.PenelitianRDRPController.Create).Methods("POST")
	penelitianRouter.HandleFunc("/rdrp", route.PenelitianRDRPController.List).Methods("GET")
	penelitianRouter.HandleFunc("/rdrp/{id}", route.PenelitianRDRPController.Update).Methods("PUT")
	penelitianRouter.HandleFunc("/rdrp/{id}", route.PenelitianRDRPController.Delete).Methods("DELETE")

}

func (route *RouteConfig) SetupAuthRoute() {
}
