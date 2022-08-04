package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizsouz404/api-rest-go/config"
	"github.com/luizsouz404/api-rest-go/controller"
	"github.com/luizsouz404/api-rest-go/middleware"
	"github.com/luizsouz404/api-rest-go/repository"
	"github.com/luizsouz404/api-rest-go/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth", middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/edit", userController.Update)
	}

	server.Run(":3000")
}
