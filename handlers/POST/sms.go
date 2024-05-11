// Download the helper library from https://www.twilio.com/docs/go/install
package POST

import (
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
	"strconv"
)

func (h *PostHandler) SendSms(g *gin.Context) {
	smsModel := models.SMS{}
	g.ShouldBindJSON(&smsModel)

	client := twilio.NewRestClient()
	toPhone := "+90" + strconv.Itoa(smsModel.Tophone)
	params := &api.CreateMessageParams{}
	params.SetBody("SMS onay şifreniz:" + smsModel.Body)
	params.SetFrom("+14583022713")
	params.SetTo(toPhone)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {
		if resp.Sid != nil {
			g.JSON(http.StatusBadRequest, *resp.Sid)
		} else {
			g.JSON(http.StatusOK, resp.Sid)
		}
	}

}
