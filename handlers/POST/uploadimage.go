package POST

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
)

type Picture struct {
	Pictureid  string `json:"Pictureid"`
	pictureurl string `json:"pictureurl"`
}

func (h *PostHandler) UploadImage(g *gin.Context) {
	picture := Picture{}
	err := g.Request.ParseMultipartForm(200000)
	if err != nil {
		fmt.Println(err.Error(), err)
		return
	}
	formdata := g.Request.MultipartForm
	files := formdata.File["image_file"]

	for i, _ := range files {
		file, err := files[i].Open()
		bytes, _ := ioutil.ReadAll(file)
		mimtype := http.DetectContentType(bytes)
		if !strings.Contains(mimtype, "image") {
			g.JSON(http.StatusBadRequest, "this is not an image")
			return
		}
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}

		uniqueID, _ := exec.Command("uuidgen").Output()
		filenameUnique := strings.Replace(string(uniqueID), "-", "", -1)
		//filenameUnique = filenameUnique[0 : len(filenameUnique)-1]
		picture.Pictureid = filenameUnique
		fileExt := strings.Split(files[i].Filename, ".")[1]
		fileID := fmt.Sprintf("%s.%s", filenameUnique, fileExt)

		out, err := os.Create("Images/" + fileID)
		defer out.Close()
		if err != nil {
			fmt.Println("unable to create the file for writing")
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println(err.Error())
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fPath := "getimage/" + fileID
		picture.pictureurl = fPath
		result := db.GetDb().Exec("INSERT INTO pictures VALUES($1,$2)", picture.Pictureid, picture.pictureurl)

		if result.Error != nil {
			g.JSON(http.StatusBadRequest, result.Error)
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": picture.Pictureid})

	}

}
