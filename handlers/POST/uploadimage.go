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
	"github.com/umitbasakk/RestApi/handlers/GET"
)

func (h *PostHandler) UploadImage(g *gin.Context) {
	token := g.Param("token")
	userId := GET.TokentoUserID(token)
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
			return
		}

		uniqueID, _ := exec.Command("uuidgen").Output()
		filenameUnique := strings.Replace(string(uniqueID), "-", "", -1)
		filenameUnique = filenameUnique[0 : len(filenameUnique)-1]
		fileExt := strings.Split(files[i].Filename, ".")[1]
		fileID := fmt.Sprintf("%s.%s", filenameUnique, fileExt)

		out, err := os.Create("Images/" + fileID)
		defer out.Close()
		if err != nil {
			fmt.Println("unable to create the file for writing")
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fPath := "getimage/" + fileID
		result := db.GetDb().Exec("UPDATE users SET profileimageurl=$1 WHERE userid=$2", fPath, userId)

		if result.Error != nil {
			g.JSON(http.StatusBadRequest, result.Error)
			return
		}
		g.JSON(http.StatusOK, "Görsel Başarıyla Yüklendi")

	}

}
