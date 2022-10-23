package configs

import (
	"CICD_GolangtoHeroku/controllers"
	"CICD_GolangtoHeroku/models"
	"CICD_GolangtoHeroku/repositories"
	"CICD_GolangtoHeroku/services"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var userRepository repositories.UserRepo
var userService services.UserService
var UserController controllers.UserController

func Init() {
	initDatabase()
	initUserRepository()
	initUserService()
	initUserController()
}

func initDatabase() {

	dsn := "postgres://lnomsnssfzcugw:55e9ac0e9fb6e0482e09463da92d81e19050cd30c59891da9d96d182d2716d84@ec2-54-173-237-110.compute-1.amazonaws.com:5432/dfneakpp90h2no"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Koneksi DB Gagal")
	}

	log.Println("Koneksi DB Berhasil")

	initMigrate()

}

func initMigrate() {
	DB.AutoMigrate(&models.User{})
}

func initUserRepository() {
	userRepository = repositories.NewUserRepo(DB)
}

func initUserService() {
	userService = services.NewUserService(userRepository)
}

func initUserController() {
	UserController = controllers.NewUserController(userService)
}
