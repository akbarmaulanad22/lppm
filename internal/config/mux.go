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

	penelitianRdrpRepository := repository.NewPenelitianRDRPRepository(config.Log)
	penelitianPdppRepository := repository.NewPenelitianPDPPRepository(config.Log)
	penelitianTcrRepository := repository.NewPenelitianTCRRepository(config.Log)
	penelitianSkrRepository := repository.NewPenelitianSKRRepository(config.Log)
	penelitianHppRepository := repository.NewPenelitianHPPRepository(config.Log)
	penelitianStpRepository := repository.NewPenelitianSTPRepository(config.Log)
	penelitianBadmeRepository := repository.NewPenelitianBADMERepository(config.Log)
	penelitianLpRepository := repository.NewPenelitianLPRepository(config.Log)

	jurnalTeknoisRepository := repository.NewJurnalTeknoisRepository(config.Log)
	jurnalTAJBRepository := repository.NewJurnalTAJBRepository(config.Log)
	jurnalTMJBRepository := repository.NewJurnalTMJBRepository(config.Log)
	jurnalJKRepository := repository.NewJurnalJKRepository(config.Log)
	jurnalJSRepository := repository.NewJurnalJSRepository(config.Log)
	jurnalKIATRepository := repository.NewJurnalKIATRepository(config.Log)

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

	penelitianRdrpUseCase := usecase.NewPenelitianRDRPUseCase(config.DB, config.Log, config.Validate, penelitianRdrpRepository)
	penelitianPdppUseCase := usecase.NewPenelitianPDPPUseCase(config.DB, config.Log, config.Validate, penelitianPdppRepository)
	penelitianTcrUseCase := usecase.NewPenelitianTCRUseCase(config.DB, config.Log, config.Validate, penelitianTcrRepository)
	penelitianSkrUseCase := usecase.NewPenelitianSKRUseCase(config.DB, config.Log, config.Validate, penelitianSkrRepository)
	penelitianHppUseCase := usecase.NewPenelitianHPPUseCase(config.DB, config.Log, config.Validate, penelitianHppRepository)
	penelitianStpUseCase := usecase.NewPenelitianSTPUseCase(config.DB, config.Log, config.Validate, penelitianStpRepository)
	penelitianBadmeUseCase := usecase.NewPenelitianBADMEUseCase(config.DB, config.Log, config.Validate, penelitianBadmeRepository)
	penelitianLpUseCase := usecase.NewPenelitianLPUseCase(config.DB, config.Log, config.Validate, penelitianLpRepository)

	jurnalTeknoisUseCase := usecase.NewJurnalTeknoisUseCase(config.DB, config.Log, config.Validate, jurnalTeknoisRepository)
	jurnalTAJBUseCase := usecase.NewJurnalTAJBUseCase(config.DB, config.Log, config.Validate, jurnalTAJBRepository)
	jurnalTMJBUseCase := usecase.NewJurnalTMJBUseCase(config.DB, config.Log, config.Validate, jurnalTMJBRepository)
	jurnalJKUseCase := usecase.NewJurnalJKUseCase(config.DB, config.Log, config.Validate, jurnalJKRepository)
	jurnalJSUseCase := usecase.NewJurnalJSUseCase(config.DB, config.Log, config.Validate, jurnalJSRepository)
	jurnalKIATUseCase := usecase.NewJurnalKIATUseCase(config.DB, config.Log, config.Validate, jurnalKIATRepository)

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

	penelitianRDRPController := route.NewPenelitianRDRPController(penelitianRdrpUseCase, config.Log)
	penelitianPDPPController := route.NewPenelitianPDPPController(penelitianPdppUseCase, config.Log)
	penelitianTCRController := route.NewPenelitianTCRController(penelitianTcrUseCase, config.Log)
	penelitianSKRController := route.NewPenelitianSKRController(penelitianSkrUseCase, config.Log)
	penelitianHPPController := route.NewPenelitianHPPController(penelitianHppUseCase, config.Log)
	penelitianSTPController := route.NewPenelitianSTPController(penelitianStpUseCase, config.Log)
	penelitianBADMEController := route.NewPenelitianBADMEController(penelitianBadmeUseCase, config.Log)
	penelitianLPController := route.NewPenelitianLPController(penelitianLpUseCase, config.Log)

	jurnalTeknoisController := route.NewJurnalTeknoisController(jurnalTeknoisUseCase, config.Log)
	jurnalTAJBController := route.NewJurnalTAJBController(jurnalTAJBUseCase, config.Log)
	jurnalTMJBController := route.NewJurnalTMJBController(jurnalTMJBUseCase, config.Log)
	jurnalJKController := route.NewJurnalJKController(jurnalJKUseCase, config.Log)
	jurnalJSController := route.NewJurnalJSController(jurnalJSUseCase, config.Log)
	jurnalKIATController := route.NewJurnalKIATController(jurnalKIATUseCase, config.Log)

	routeConfig := route.RouteConfig{
		Router:                    config.Router,
		ProfilVisiMisiController:  profilVisiMisiController,
		ProfilSODTController:      profilSODTController,
		PKMRDRPController:         PKMRDRPController,
		PKMPDPPController:         PKMPDPPController,
		PKMTCRController:          PKMTCRController,
		PKMSKRController:          PKMSKRController,
		PKMHPPController:          PKMHPPController,
		PKMSTPController:          PKMSTPController,
		PKMBADMEController:        PKMBADMEController,
		PKMLPController:           PKMLPController,
		HKIMHSController:          hkiMhsController,
		HKIDosenController:        hkiDosenController,
		PenelitianRDRPController:  penelitianRDRPController,
		PenelitianPDPPController:  penelitianPDPPController,
		PenelitianTCRController:   penelitianTCRController,
		PenelitianSKRController:   penelitianSKRController,
		PenelitianHPPController:   penelitianHPPController,
		PenelitianSTPController:   penelitianSTPController,
		PenelitianBADMEController: penelitianBADMEController,
		PenelitianLPController:    penelitianLPController,
		JurnalTeknoisController:   jurnalTeknoisController,
		JurnalTAJBController:      jurnalTAJBController,
		JurnalTMJBController:      jurnalTMJBController,
		JurnalJKController:        jurnalJKController,
		JurnalJSController:        jurnalJSController,
		JurnalKIATController:      jurnalKIATController,
	}
	routeConfig.Setup()

}
