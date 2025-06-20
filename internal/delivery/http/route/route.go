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
	PKMSKRController         *PKMSKRController
	PKMHPPController         *PKMHPPController
	PKMSTPController         *PKMSTPController
	PKMBADMEController       *PKMBADMEController
	PKMLPController          *PKMLPController
	HKIMHSController         *HKIMHSController
	HKIDosenController       *HKIDosenController
	PenelitianRDRPController *PenelitianRDRPController
	PenelitianPDPPController *PenelitianPDPPController
	PenelitianTCRController  *PenelitianTCRController
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

	pkmRouter.HandleFunc("/skr", route.PKMSKRController.Create).Methods("POST")
	pkmRouter.HandleFunc("/skr", route.PKMSKRController.List).Methods("GET")
	pkmRouter.HandleFunc("/skr/{id}", route.PKMSKRController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/skr/{id}", route.PKMSKRController.Delete).Methods("DELETE")

	pkmRouter.HandleFunc("/hpp", route.PKMHPPController.Create).Methods("POST")
	pkmRouter.HandleFunc("/hpp", route.PKMHPPController.List).Methods("GET")
	pkmRouter.HandleFunc("/hpp/{id}", route.PKMHPPController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/hpp/{id}", route.PKMHPPController.Delete).Methods("DELETE")

	pkmRouter.HandleFunc("/stp", route.PKMSTPController.Create).Methods("POST")
	pkmRouter.HandleFunc("/stp", route.PKMSTPController.List).Methods("GET")
	pkmRouter.HandleFunc("/stp/{id}", route.PKMSTPController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/stp/{id}", route.PKMSTPController.Delete).Methods("DELETE")

	pkmRouter.HandleFunc("/badme", route.PKMBADMEController.Create).Methods("POST")
	pkmRouter.HandleFunc("/badme", route.PKMBADMEController.List).Methods("GET")
	pkmRouter.HandleFunc("/badme/{id}", route.PKMBADMEController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/badme/{id}", route.PKMBADMEController.Delete).Methods("DELETE")

	pkmRouter.HandleFunc("/lp", route.PKMLPController.Create).Methods("POST")
	pkmRouter.HandleFunc("/lp", route.PKMLPController.List).Methods("GET")
	pkmRouter.HandleFunc("/lp/{id}", route.PKMLPController.Update).Methods("PUT")
	pkmRouter.HandleFunc("/lp/{id}", route.PKMLPController.Delete).Methods("DELETE")

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

	penelitianRouter.HandleFunc("/pdpp", route.PenelitianPDPPController.Create).Methods("POST")
	penelitianRouter.HandleFunc("/pdpp", route.PenelitianPDPPController.List).Methods("GET")
	penelitianRouter.HandleFunc("/pdpp/{id}", route.PenelitianPDPPController.Update).Methods("PUT")
	penelitianRouter.HandleFunc("/pdpp/{id}", route.PenelitianPDPPController.Delete).Methods("DELETE")

	penelitianRouter.HandleFunc("/tcr", route.PenelitianTCRController.Create).Methods("POST")
	penelitianRouter.HandleFunc("/tcr", route.PenelitianTCRController.List).Methods("GET")
	penelitianRouter.HandleFunc("/tcr/{id}", route.PenelitianTCRController.Update).Methods("PUT")
	penelitianRouter.HandleFunc("/tcr/{id}", route.PenelitianTCRController.Delete).Methods("DELETE")

}

func (route *RouteConfig) SetupAuthRoute() {
}
