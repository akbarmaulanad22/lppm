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
	profilVisiMisiRepository := repository.NewProfilVisiMisiRepository(config.Log)
	profilSODTRepository := repository.NewProfilSODTRepository(config.Log)
	pkmRepository := repository.NewPKMRDRPRepository(config.Log)

	// setup use cases
	profilVisiMisiUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, profilVisiMisiRepository)
	profilSODTUseCase := usecase.NewProfilSODTUseCase(config.DB, config.Log, config.Validate, profilSODTRepository)
	pkmUseCase := usecase.NewPKMRDRPUseCase(config.DB, config.Log, config.Validate, pkmRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(profilVisiMisiUseCase, config.Log)
	profilSODTController := route.NewProfilSODTController(profilSODTUseCase, config.Log)
	PKMRDRPController := route.NewPKMRDRPController(pkmUseCase, config.Log)

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
		ProfilSODTController:     profilSODTController,
		PKMRDRPController:        PKMRDRPController,
	}
	routeConfig.Setup()

}
