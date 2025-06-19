package config

import (
	"tugasakhir/internal/delivery/http/route"
	"tugasakhir/internal/repository"
	"tugasakhir/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type MuxConfig struct {
	Router   *mux.Router
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func NewMux(config *MuxConfig) {

	// setup repositories
<<<<<<< HEAD
	profilRepository := repository.NewProfilVisiMisiRepository(config.Log)
	pkmrdrpRepository := repository.NewPKMRDRPRepository(config.Log)
	pkmpdppRepository := repository.NewPKMPDPPRepository(config.Log)

	// setup use cases
	profilUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, profilRepository)
	pkmrdrpUseCase := usecase.NewPKMRDRPUseCase(config.DB, config.Log, config.Validate, pkmrdrpRepository)
	pkmpdppUseCase := usecase.NewPKMPDPPUseCase(config.DB, config.Log, config.Validate, pkmpdppRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(profilUseCase, config.Log)
	PKMRDRPController := route.NewPKMRDRPController(pkmrdrpUseCase, config.Log)
	PKMPDPPController := route.NewPKMPDPPController(pkmpdppUseCase, config.Log)
=======
	profilVisiMisiRepository := repository.NewProfilVisiMisiRepository(config.Log)
	profilSODTRepository := repository.NewProfilSODTRepository(config.Log)
	pkmRepository := repository.NewPKMRDRPRepository(config.Log)
	hkiMhsRepository := repository.NewHKIMHSRepository(config.Log)

	// setup use cases
	profilVisiMisiUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, profilVisiMisiRepository)
	profilSODTUseCase := usecase.NewProfilSODTUseCase(config.DB, config.Log, config.Validate, profilSODTRepository)
	pkmUseCase := usecase.NewPKMRDRPUseCase(config.DB, config.Log, config.Validate, pkmRepository)
	hkiMhsUseCase := usecase.NewHKIMHSUseCase(config.DB, config.Log, config.Validate, hkiMhsRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(profilVisiMisiUseCase, config.Log)
	profilSODTController := route.NewProfilSODTController(profilSODTUseCase, config.Log)
	PKMRDRPController := route.NewPKMRDRPController(pkmUseCase, config.Log)
	hkiMhsController := route.NewHKIMHSController(hkiMhsUseCase, config.Log)
>>>>>>> 8cb079c3c0f397c989d3ad4c7e9cbf477793c8cc

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
<<<<<<< HEAD
		PKMRDRPController: PKMRDRPController,
		PKMPDPPController: PKMPDPPController,
=======
		ProfilSODTController:     profilSODTController,
		PKMRDRPController:        PKMRDRPController,
		HKIMHSController:         hkiMhsController,
>>>>>>> 8cb079c3c0f397c989d3ad4c7e9cbf477793c8cc
	}
	routeConfig.Setup()

}
