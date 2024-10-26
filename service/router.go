package router

import (
	"github.com/gofiber/fiber/v2"
	"initial-project-go/di/config"
	"initial-project-go/di/database"
	"initial-project-go/repository"
	"initial-project-go/service/api_keys"
	service3 "initial-project-go/service/auth"
	"initial-project-go/service/middleware"
	service2 "initial-project-go/service/user"
	"log"
)

func InitRouter(server *fiber.App) {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	appConfig := config.GetConfig()
	apiKeysService := service.ProvideApiKeysService(db, appConfig)
	userService := service2.ProvideUserService(db, appConfig)
	authService := service3.ProvideAuthService(db, appConfig)
	encryptorRepository := repository.ProvideEncryptorRepository(db, appConfig)
	userRepository := repository.ProvideUserRepository(db, appConfig)

	secret, err := encryptorRepository.GetPassphrase()
	if err != nil {
		panic(err)
	}

	server.Post("/secret/generate", apiKeysService.HandleSecretPostRouter)
	server.Post("/secret/rotate", apiKeysService.HandleRotatePostRouter)
	server.Get("/secret/verify", apiKeysService.HandleVerifyGetRouter)

	// End point User
	server.Get("/user", userService.HandleGetAllUser)
	server.Post("/user", userService.HandelCreatUser)
	server.Delete("user/:id", userService.HandelDeleteUser)
	server.Delete("/user/del/:id", userService.HandleHarddel)
	server.Put("/user/:id", userService.HandleUpdateUser)

	// End point Sign-in
	server.Post("/user/sign-in", authService.HandleSingIn)
	server.Post("/user/login", authService.HandleLogin)

	group := server.Group("/user", middleware.JWTMiddleware(string(secret.Hash), userRepository))
	group.Get("/me", authService.HandlerMe)
}
