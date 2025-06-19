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
	pkmRepository := repository.NewPKMRDRPRepository(config.Log)

	// setup use cases
	profilUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, profilRepository)
	pkmUseCase := usecase.NewPKMRDRPUseCase(config.DB, config.Log, config.Validate, pkmRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(profilUseCase, config.Log)
	PKMRDRPController := route.NewPKMRDRPController(pkmUseCase, config.Log)
=======
	userRepository := repository.NewProfilVisiMisiRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, userRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(userUseCase, config.Log)
>>>>>>> 006ad76c5e76579d26dc36a60676ed24113c225b

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
<<<<<<< HEAD
		PKMRDRPController: PKMRDRPController,
=======
>>>>>>> 006ad76c5e76579d26dc36a60676ed24113c225b
	}
	routeConfig.Setup()

}
