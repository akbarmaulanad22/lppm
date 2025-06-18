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
	userRepository := repository.NewProfilVisiMisiRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, userRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(userUseCase, config.Log)

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
	}
	routeConfig.Setup()

}
