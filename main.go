package main

import (
	"ServerRestApi/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

func main() {

	routers := gin.Default()
	godotenv.Load(".env")

	PORT := os.Getenv("PORT")

	user := &models.User{
		ProfileURL:         "asd",
		FullName:           "Umit Basak",
		Email:              "umitbasaak@gmail.com",
		Mobile:             "05348297134",
		Gender:             "E",
		Birthday:           time.Now(),
		Followings:         12,
		Followers:          12,
		ProfileDescription: "hello world",
		WhatsappURL:        "wp",
		FacebookURL:        "fb",
		InstagramURL:       "ins",
		TwitterURL:         "twt",
	}
	comment_1 := &models.Comment{
		User:       *user,
		CommenText: "Comment 1",
	}

	comment_2 := &models.Comment{
		User:       *user,
		CommenText: "Comment 2",
	}

	comment_3 := &models.Comment{
		User:       *user,
		CommenText: "Comment 3",
	}

	comment_4 := &models.Comment{
		User:       *user,
		CommenText: "Comment 4",
	}

	articleTest1 := &models.Article{
		ImageURL:       "asd",
		Title:          "Hello Worldx",
		CreatedTime:    time.Now(),
		Author:         *user,
		CategoryID:     1,
		ArticleContent: "Hello world in context",
		Comments: []models.Comment{
			*comment_1, *comment_2, *comment_3, *comment_4,
		},
	}

	articleTest2 := &models.Article{
		ImageURL:       "asd",
		Title:          "Hello World",
		CreatedTime:    time.Now(),
		Author:         *user,
		CategoryID:     1,
		ArticleContent: "Hello world in context",
		Comments: []models.Comment{
			*comment_1, *comment_2, *comment_3, *comment_4,
		},
	}

	articleTest3 := &models.Article{
		ImageURL:       "asd",
		Title:          "Hello World",
		CreatedTime:    time.Now(),
		Author:         *user,
		CategoryID:     1,
		ArticleContent: "Hello world in context",
		Comments: []models.Comment{
			*comment_1, *comment_2, *comment_3, *comment_4,
		},
	}
	articleArray := []models.Article{*articleTest1, *articleTest2, *articleTest3}

	routers.GET("/articles", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": articleArray})
	})
	routers.Run(":" + PORT)

}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
