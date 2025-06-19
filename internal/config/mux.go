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

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
		PKMRDRPController: PKMRDRPController,
		PKMPDPPController: PKMPDPPController,
	}
	routeConfig.Setup()

}
