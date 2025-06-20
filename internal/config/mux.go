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
	pkmRdrpRepository := repository.NewPKMRDRPRepository(config.Log)
	pkmPdppRepository := repository.NewPKMPDPPRepository(config.Log)
	pkmTcrRepository := repository.NewPKMTCRRepository(config.Log)
	pkmSkrRepository := repository.NewPKMSKRRepository(config.Log)
	hkiMhsRepository := repository.NewHKIMHSRepository(config.Log)
	hkiDosenRepository := repository.NewHKIDosenRepository(config.Log)

	// setup use cases
	profilVisiMisiUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, profilVisiMisiRepository)
	profilSODTUseCase := usecase.NewProfilSODTUseCase(config.DB, config.Log, config.Validate, profilSODTRepository)
	pkmRdrpUseCase := usecase.NewPKMRDRPUseCase(config.DB, config.Log, config.Validate, pkmRdrpRepository)
	pkmPdppUseCase := usecase.NewPKMPDPPUseCase(config.DB, config.Log, config.Validate, pkmPdppRepository)
	pkmTcrUseCase := usecase.NewPKMTCRUseCase(config.DB, config.Log, config.Validate, pkmTcrRepository)
	pkmSkrUseCase := usecase.NewPKMSKRUseCase(config.DB, config.Log, config.Validate, pkmSkrRepository)
	hkiMhsUseCase := usecase.NewHKIMHSUseCase(config.DB, config.Log, config.Validate, hkiMhsRepository)
	hkiDosenUseCase := usecase.NewHKIDosenUseCase(config.DB, config.Log, config.Validate, hkiDosenRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(profilVisiMisiUseCase, config.Log)
	profilSODTController := route.NewProfilSODTController(profilSODTUseCase, config.Log)
	PKMRDRPController := route.NewPKMRDRPController(pkmRdrpUseCase, config.Log)
	PKMPDPPController := route.NewPKMPDPPController(pkmPdppUseCase, config.Log)
	PKMTCRController := route.NewPKMTCRController(pkmTcrUseCase, config.Log)
	PKMSKRController := route.NewPKMSKRController(pkmSkrUseCase, config.Log)
	hkiMhsController := route.NewHKIMHSController(hkiMhsUseCase, config.Log)
	hkiDosenController := route.NewHKIDosenController(hkiDosenUseCase, config.Log)

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
		ProfilSODTController:     profilSODTController,
		PKMRDRPController:        PKMRDRPController,
		PKMPDPPController:        PKMPDPPController,
		PKMTCRController:         PKMTCRController,
		PKMSKRController:         PKMSKRController,
		HKIMHSController:         hkiMhsController,
		HKIDosenController:       hkiDosenController,
	}
	routeConfig.Setup()

}
