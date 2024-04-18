package main

import (
	"ServerRestApi/handlers/GET"
	"ServerRestApi/handlers/POST"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	routers := gin.Default()
	godotenv.Load(".env")

	LISTENPORT := os.Getenv("LISTENPORT")
	getHandler := GET.GetHandler{}
	postHandler := POST.PostHandler{}
	//--- GET ---\\
	routers.GET("/users", getHandler.GetAllUsers)
	routers.GET("/users/:userid", getHandler.GetUser)
	routers.GET("/articles", getHandler.GetAllArticles)
	routers.GET("/articles/:articleid", getHandler.GetArticle)

	//--- POST ---\\
	routers.POST("/users", postHandler.CreateUser)
	routers.POST("/articles", postHandler.CreateArticle)

	routers.Run(":" + LISTENPORT)

}
