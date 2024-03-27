package main

import (
	"context"
	"fmt"
	"work_in_que/auth"
	chats "work_in_que/chats"
	"work_in_que/health"
	"work_in_que/logging"
	"work_in_que/middleware"
	"work_in_que/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"work_in_que/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Starting...")
	logger := logging.NewLogger()

	// env.VerifyRequiredEnvVarsSet()

	dbName := "mydb"
	client, err := db.CreateDatabaseConnection(dbName)
	if err != nil {
		fmt.Println("Failed to connect to DB")
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(dbName)

	// Repositories
	userRepository := user.NewInstanceOfUserRepository(db)
	chatsRepository := chats.NewInstanceOfchatsRepository(db)
	forgotPasswordRepository := user.NewInstanceOfForgotPasswordRepository(db)

	// Services
	userServices := user.NewInstanceOfUserServices(logger, userRepository, forgotPasswordRepository)
	chatsServices := chats.NewInstanceOfchatsServices(logger, userRepository, chatsRepository)

	// Handlers
	userHandlers := user.NewInstanceOfUserHandlers(logger, userServices)
	chatsHandlers := chats.NewInstanceOfchatsHandlers(logger, chatsServices)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	healthAPI := router.Group("/")
	{
		healthAPI.GET("", health.Check)
		healthAPI.GET("health", health.Check)
	}

	userAPI := router.Group("/user")
	{
		userAPI.POST("/signin", userHandlers.SignIn)
		userAPI.GET("/signup", userHandlers.SignUp)
		userAPI.POST("/signout", auth.ValidateAuth(userRepository), userHandlers.LogOut)
		userAPI.POST("/session/unlock", userHandlers.UnlockSession)
		userAPI.POST("/forgot-password/", userHandlers.SendForgotPassword)
		userAPI.POST("/forgot-password/reset", userHandlers.ForgotPassword)
	}

	chatsAPI := router.Group("/chats")
	{
		chatsAPI.GET("/", auth.ValidateAuth(userRepository), chatsHandlers.GetAll)
		chatsAPI.GET("/:id", auth.ValidateAuth(userRepository), chatsHandlers.GetByID)
		chatsAPI.POST("/", auth.ValidateAuth(userRepository), chatsHandlers.Create)
		chatsAPI.PUT("/:id", auth.ValidateAuth(userRepository), chatsHandlers.Update)
		chatsAPI.DELETE("/:id", auth.ValidateAuth(userRepository), chatsHandlers.Delete)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
