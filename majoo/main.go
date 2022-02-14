package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"majoo/config"
	"majoo/docs"
	"majoo/controllers"
	"majoo/middleware"
	"majoo/repo"
	"majoo/service"

	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                       = config.SetupDatabaseConnection()
	userRepo           repo.UserRepository            = repo.NewUserRepo(db)
	transactionRepo    repo.TransactionRepository     = repo.NewTransactionRepo(db)
	authService        service.AuthService            = service.NewAuthService(userRepo)
	jwtService         service.JWTService             = service.NewJWTService()
	userService        service.UserService            = service.NewUserService(userRepo)
	transactionService service.TransactionService     = service.NewTransactionService(transactionRepo)
	authHandler        controllers.AuthHandler        = controllers.NewAuthHandler(authService, jwtService, userService)
	transactionHandler controllers.TransactionHandler = controllers.NewTransactionHandler(authService, jwtService, userService, transactionService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
		
	}))
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "Majoo"
	docs.SwaggerInfo.Description = "Test Majoo"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:8080"

	authRoutes := server.Group("auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	checkRoutes := server.Group("check")
	{
		checkRoutes.GET("health", controllers.Health)
	}

	transactionRoutes := server.Group("transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.GET("/:dateFrom/:dateTo/:page/:limit", transactionHandler.TransactionReport)
	}

	server.Run()
}
