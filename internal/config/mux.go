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
	pkmHppRepository := repository.NewPKMHPPRepository(config.Log)
	pkmStpRepository := repository.NewPKMSTPRepository(config.Log)
	pkmBadmeRepository := repository.NewPKMBADMERepository(config.Log)
	pkmLpRepository := repository.NewPKMLPRepository(config.Log)
	hkiMhsRepository := repository.NewHKIMHSRepository(config.Log)
	hkiDosenRepository := repository.NewHKIDosenRepository(config.Log)
	penelitianRDRPRepository := repository.NewPenelitianRDRPRepository(config.Log)
	penelitianPDPPRepository := repository.NewPenelitianPDPPRepository(config.Log)

	// setup use cases
	profilVisiMisiUseCase := usecase.NewProfilVisiMisiUseCase(config.DB, config.Log, config.Validate, profilVisiMisiRepository)
	profilSODTUseCase := usecase.NewProfilSODTUseCase(config.DB, config.Log, config.Validate, profilSODTRepository)
	pkmRdrpUseCase := usecase.NewPKMRDRPUseCase(config.DB, config.Log, config.Validate, pkmRdrpRepository)
	pkmPdppUseCase := usecase.NewPKMPDPPUseCase(config.DB, config.Log, config.Validate, pkmPdppRepository)
	pkmTcrUseCase := usecase.NewPKMTCRUseCase(config.DB, config.Log, config.Validate, pkmTcrRepository)
	pkmSkrUseCase := usecase.NewPKMSKRUseCase(config.DB, config.Log, config.Validate, pkmSkrRepository)
	pkmHppUseCase := usecase.NewPKMHPPUseCase(config.DB, config.Log, config.Validate, pkmHppRepository)
	pkmStpUseCase := usecase.NewPKMSTPUseCase(config.DB, config.Log, config.Validate, pkmStpRepository)
	pkmBadmeUseCase := usecase.NewPKMBADMEUseCase(config.DB, config.Log, config.Validate, pkmBadmeRepository)
	pkmLpUseCase := usecase.NewPKMLPUseCase(config.DB, config.Log, config.Validate, pkmLpRepository)
	hkiMhsUseCase := usecase.NewHKIMHSUseCase(config.DB, config.Log, config.Validate, hkiMhsRepository)
	hkiDosenUseCase := usecase.NewHKIDosenUseCase(config.DB, config.Log, config.Validate, hkiDosenRepository)
	penelitianRDRPUseCase := usecase.NewPenelitianRDRPUseCase(config.DB, config.Log, config.Validate, penelitianRDRPRepository)
	penelitianPDPPUseCase := usecase.NewPenelitianPDPPUseCase(config.DB, config.Log, config.Validate, penelitianPDPPRepository)

	// setup controller
	profilVisiMisiController := route.NewProfilVisiMisiController(profilVisiMisiUseCase, config.Log)
	profilSODTController := route.NewProfilSODTController(profilSODTUseCase, config.Log)
	PKMRDRPController := route.NewPKMRDRPController(pkmRdrpUseCase, config.Log)
	PKMPDPPController := route.NewPKMPDPPController(pkmPdppUseCase, config.Log)
	PKMTCRController := route.NewPKMTCRController(pkmTcrUseCase, config.Log)
	PKMSKRController := route.NewPKMSKRController(pkmSkrUseCase, config.Log)
	PKMHPPController := route.NewPKMHPPController(pkmHppUseCase, config.Log)
	PKMSTPController := route.NewPKMSTPController(pkmStpUseCase, config.Log)
	PKMBADMEController := route.NewPKMBADMEController(pkmBadmeUseCase, config.Log)
	PKMLPController := route.NewPKMLPController(pkmLpUseCase, config.Log)
	hkiMhsController := route.NewHKIMHSController(hkiMhsUseCase, config.Log)
	hkiDosenController := route.NewHKIDosenController(hkiDosenUseCase, config.Log)
	penelitianRDRPController := route.NewPenelitianRDRPController(penelitianRDRPUseCase, config.Log)
	penelitianPDPPController := route.NewPenelitianPDPPController(penelitianPDPPUseCase, config.Log)

	routeConfig := route.RouteConfig{
		Router:                   config.Router,
		ProfilVisiMisiController: profilVisiMisiController,
		ProfilSODTController:     profilSODTController,
		PKMRDRPController:        PKMRDRPController,
		PKMPDPPController:        PKMPDPPController,
		PKMTCRController:         PKMTCRController,
		PKMSKRController:         PKMSKRController,
		PKMHPPController:         PKMHPPController,
		PKMSTPController:         PKMSTPController,
		PKMBADMEController:         PKMBADMEController,
		PKMLPController:         PKMLPController,
		HKIMHSController:         hkiMhsController,
		HKIDosenController:       hkiDosenController,
		PenelitianRDRPController: penelitianRDRPController,
		PenelitianPDPPController: penelitianPDPPController,
	}
	routeConfig.Setup()

}
