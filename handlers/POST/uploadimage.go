package POST

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/umitbasakk/RestApi/db"
)

type Picture struct {
	Pictureid  string `json:"Pictureid"`
	pictureurl string `json:"pictureurl"`
}

func (h *PostHandler) UploadImage(g *gin.Context) {
	picture := Picture{}
	err := g.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println(err.Error(), err)
		return
	}
	file, handler, err := g.Request.FormFile("image_file")

	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	uniqueID := uuid.New().String()
	filenameUnique := strings.Replace(uniqueID, "-", "", -1)
	picture.Pictureid = filenameUnique
	fileExt := strings.Split(handler.Filename, ".")[1]
	fileID := fmt.Sprintf("%s.%s", filenameUnique, fileExt)

	f, err := os.OpenFile("./Images/"+fileID, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("error var")
	}

	defer f.Close()
	io.Copy(f, file)

	fPath := "getimage/" + fileID
	picture.pictureurl = fPath
	result := db.GetDb().Exec("INSERT INTO pictures VALUES($1,$2)", picture.Pictureid, picture.pictureurl)

	if result.Error != nil {
		g.JSON(http.StatusBadRequest, result.Error)
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": picture.Pictureid})

}
