package main

import (
	"blog-api/config"
	"blog-api/controller"
	"blog-api/middlewares"
	"blog-api/routes"
	"blog-api/services"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {


	// ** load env file 
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//** make database connection
	mongoUri := os.Getenv("MONGO_URI")
	fmt.Println(mongoUri)
    client,err := config.ConnectDatabase(mongoUri)
	if err != nil {
		log.Fatal("failed to connect to the database", err)
	}
   
	// ** create a user collection and pass it to the user service
	userCollection := client.Database("blog-api").Collection("users")
	userService := &services.UserService{DB: userCollection}

	// ** create a user controller and pass it to the user routes
	userController := &controller.UserController{UserService: userService}

	app := gin.Default()
    

	app.GET("/test", middlewares.AuthMiddleware(), func(c *gin.Context) {
         userId := c.MustGet("userId").(string)
		 username := c.MustGet("username").(string)
		 c.JSON(200, gin.H{"userId": userId, "username": username})
	})



	 userRoutes := app.Group("/api/users")
	 routes.UserRoutes(userRoutes,userController)

	if err := app.Run(":3000"); err != nil {
		log.Fatal("failed to run the app", err)
	}
	

}
