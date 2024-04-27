package main

import (
	"os"

	"github.com/umitbasakk/RestApi/handlers/GET"
	"github.com/umitbasakk/RestApi/handlers/POST"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	routers.GET("/categorys", getHandler.GetAllCategorys)
	routers.GET("/categorys/:categoryid", getHandler.GetCategory)

	routers.GET("/subcategorys", getHandler.GetAllSubCategorys)
	routers.GET("/subcategory/:subcategoryid", getHandler.GetSubCategory)
	routers.GET("/follow/:myid", getHandler.GetFollowersData)

	routers.GET("/comments/:articleid", getHandler.GetComments)
	routers.GET("/comments", getHandler.GetAllComments)

	//--- POST ---\\
	routers.POST("/register", postHandler.CreateUser)
	routers.POST("/login", postHandler.LoginUser)
	routers.POST("/articles", postHandler.CreateArticle)
	routers.POST("/category", postHandler.CreateCategory)
	routers.POST("/subcategory", postHandler.CreateSubCategory)
	routers.POST("/subcategoryonuser/:userid", postHandler.AddSubCategoryOnUser)
	routers.POST("/follow/:followedby/:followed", postHandler.FollowUser)
	routers.POST("/comments", postHandler.CreateComment)

	routers.Run(":" + LISTENPORT)

}
