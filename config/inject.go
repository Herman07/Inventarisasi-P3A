package config

import (
	"gorm.io/gorm"
	"Inventarisasi-P3A/controllers"
	"Inventarisasi-P3A/repository"
	"Inventarisasi-P3A/services"
)

func InjectAuthController(db *gorm.DB) controllers.FrontAuthController {
	authRepository := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	return authController
}
func InjectDashboardController(db *gorm.DB) controllers.P3Controller {
	dashboardRepository := repository.NewP3Repository(db)
	dashboardService := services.NewP3Service(dashboardRepository)
	dashboardController := controllers.NewP3Controller(dashboardService)
	return dashboardController
}

func InjectMasterController(db *gorm.DB) controllers.ProvDataController {
	masterRepository := repository.NewProvDataRepository(db)
	masterService := services.NewProvDataService(masterRepository)
	masterController := controllers.NewProvDataController(masterService)
	return masterController
}

func InjectKabController(db *gorm.DB) controllers.KabDataController {
	kabRepository := repository.NewKabDataRepository(db)
	kabService := services.NewKabDataService(kabRepository)
	kabController := controllers.NewKabDataController(kabService)
	return kabController
}
func InjectKecController(db *gorm.DB) controllers.KecDataController {
	kecRepository := repository.NewKecDataRepository(db)
	kecService := services.NewKecDataService(kecRepository)
	kecController := controllers.NewKecDataController(kecService)
	return kecController
}
func InjectUserController(db *gorm.DB) controllers.UsersDataController {
	userRepository := repository.NewUsersDataRepository(db)
	userService := services.NewUsersDataService(userRepository)
	userController := controllers.NewUsersDataController(userService)
	return userController
}