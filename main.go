package main

import (
	"context"
	"fmt"
	"work_in_que/auth"
	"work_in_que/cars"
	"work_in_que/health"
	"work_in_que/logging"
	"work_in_que/middleware"
	"work_in_que/user"

	//"strings"
	// "time"

	"work_in_que/db"

	"github.com/gin-gonic/gin"

	// "github.com/gin-contrib/cors"
	// "database/sql"
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
	carsRepository := cars.NewInstanceOfCarsRepository(db)
	forgotPasswordRepository := user.NewInstanceOfForgotPasswordRepository(db)

	// Services
	userServices := user.NewInstanceOfUserServices(logger, userRepository, forgotPasswordRepository)
	carsServices := cars.NewInstanceOfCarsServices(logger, userRepository, carsRepository)

	// Handlers
	userHandlers := user.NewInstanceOfUserHandlers(logger, userServices)
	carsHandlers := cars.NewInstanceOfCarsHandlers(logger, carsServices)

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

	carsAPI := router.Group("/cars")
	{
		carsAPI.GET("/", auth.ValidateAuth(userRepository), carsHandlers.GetAll)
		carsAPI.GET("/:id", auth.ValidateAuth(userRepository), carsHandlers.GetByID)
		carsAPI.POST("/", auth.ValidateAuth(userRepository), carsHandlers.Create)
		carsAPI.PUT("/:id", auth.ValidateAuth(userRepository), carsHandlers.Update)
		carsAPI.DELETE("/:id", auth.ValidateAuth(userRepository), carsHandlers.Delete)
	}

	router.Run(":8080")
}
