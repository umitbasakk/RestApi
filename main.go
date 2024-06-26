package main

import (
	"net/http"
	"os"

	"github.com/umitbasakk/RestApi/handlers/DELETE"

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
	deleteHandler := DELETE.DeleteHandler{}
	//--- GET ---\\
	routers.GET("/users", getHandler.GetAllUsers)
	routers.GET("/users/:userid", getHandler.GetUser)

	routers.GET("/articles", getHandler.GetAllArticles)
	routers.GET("/articles/get/:userid", getHandler.GetAllArticlesToUser)
	routers.GET("/articles/:articleid", getHandler.GetArticle)
	routers.GET("/articles/subcategory/:subcategoryid", getHandler.GetArticlesofSubcategory)

	routers.GET("/categorys", getHandler.GetAllCategorys)
	routers.GET("/categorys/:categoryid", getHandler.GetCategory)

	routers.GET("/subcategorys", getHandler.GetAllSubCategorys)
	routers.GET("/subcategory/:subcategoryid", getHandler.GetSubCategory)
	routers.GET("/follow/:username", getHandler.GetFollowersData)

	routers.GET("/comments/:articleid", getHandler.GetComments)
	routers.GET("/comments", getHandler.GetAllComments)
	routers.GET("/bookmark/:username", getHandler.GetBookmark)
	routers.StaticFS("/getimage", http.Dir("data"))

	//--- POST ---\\
	routers.POST("/register", postHandler.CreateUser)
	routers.POST("/login", postHandler.LoginUser)
	routers.POST("/articles/:userid", postHandler.CreateArticle)
	routers.POST("/category", postHandler.CreateCategory)
	routers.POST("/subcategory", postHandler.CreateSubCategory)
	routers.POST("/subcategoryonuser/:userid", postHandler.AddSubCategoryOnUser)
	routers.POST("/follow", postHandler.FollowUser)
	routers.POST("/comments", postHandler.CreateComment)
	routers.POST("/uploadimage", postHandler.UploadImage)
	routers.POST("/bookmark", postHandler.CreateBookmark)
	routers.POST("/sms/send", postHandler.SendSms)
	routers.POST("/verify", postHandler.VerifyUser)

	routers.DELETE("/delete/bookmark", deleteHandler.DeleteBookmark)
	routers.DELETE("/delete/follow", deleteHandler.UnFollowUser)

	routers.Run(":" + LISTENPORT)

}
