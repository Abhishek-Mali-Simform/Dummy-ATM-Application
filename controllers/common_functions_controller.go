package controllers

import (
	"Application/models"
	"github.com/Abhishek-Mali-Simform/oauth2"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	conifg "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type CommonFunction struct {
	beego.Controller
}

func (controller *CommonFunction) Send(resp models.Response) {
	controller.Data["json"] = resp
	serveJsonError := controller.ServeJSON()
	if serveJsonError != nil {
		logs.Error("controller/common_functions_controller.go : line-15 : ", serveJsonError)
		controller.Abort("500")
		os.Exit(-1)
	}
	controller.Finish()
}

func (controller *CommonFunction) InvalidMethod() {
	resp := models.DataResponse(models.InvalidRequestMsg, models.InvalidRequestStatus, make([]interface{}, 0))
	controller.Send(resp)
}

func (controller *CommonFunction) CheckError(err error, msg, status string) bool {
	if err != nil {
		logs.Error(err)
		resp := models.DataResponse(msg, status, make([]interface{}, 0))
		controller.Send(resp)
		return true
	}
	return false
}

func (controller *CommonFunction) GetID(paramName string) int32 {
	paramValue := controller.Ctx.Input.Query(paramName)
	ID, strConError := strconv.Atoi(paramValue)
	if !controller.CheckError(strConError, models.FailParsingMsg, models.FailParsingStatus) {
		return int32(ID)
	}
	return 0
}

func (controller *CommonFunction) GetGoogleConfig() *oauth2.ConfigureOAuth2 {
	return &oauth2.ConfigureOAuth2{
		Config: conifg.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes:       strings.Split(os.Getenv("GOOGLE_SCOPES"), ","),
		},
		State:         os.Getenv("GOOGLE_STATE"),
		RequestMethod: http.MethodGet,
		Body:          http.NoBody,
		GetInfoURL:    os.Getenv("GOOGLE_GET_DETAIL_URL") + "?access_token=",
	}
}

func (controller *CommonFunction) GetMicrosoftConfig() *oauth2.ConfigureOAuth2 {
	return &oauth2.ConfigureOAuth2{
		Config: conifg.Config{
			ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
			ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
			Endpoint:     microsoft.AzureADEndpoint(os.Getenv("MICROSOFT_TENANT_ID")),
			RedirectURL:  os.Getenv("MICROSOFT_REDIRECT_URL"),
			Scopes:       strings.Split(os.Getenv("MICROSOFT_SCOPES"), ","),
		},
		State:         os.Getenv("MICROSOFT_STATE"),
		RequestMethod: http.MethodGet,
		Body:          http.NoBody,
		GetInfoURL:    os.Getenv("MICROSOFT_GET_DETAIL_URL"),
	}
}

func (controller *CommonFunction) SetCredentials(token []byte, claim int) bool {
	tokenError := controller.SetSession("token", token)
	claimError := controller.SetSession("claim", claim)
	if controller.CheckError(
		tokenError,
		models.FailSettingSessionMsg,
		models.FailSettingSessionStatus,
	) && controller.CheckError(
		claimError,
		models.FailSettingSessionMsg,
		models.FailSettingSessionStatus,
	) {
		return false
	}
	return true
}

func (controller *CommonFunction) GetCredentials() (token []byte, claim int, check bool) {
	token, check = controller.GetSession("token").([]byte)
	if check {
		claim, check = controller.GetSession("claim").(int)
	}
	return
}

func (controller *CommonFunction) RemoveCredentials() bool {
	deleteTokenError := controller.DelSession("token")
	deleteClaimError := controller.DelSession("claim")
	if controller.CheckError(
		deleteTokenError,
		models.FailSettingSessionMsg,
		models.FailSettingSessionStatus,
	) && controller.CheckError(
		deleteClaimError,
		models.FailSettingSessionMsg,
		models.FailSettingSessionStatus,
	) {
		return false
	}
	return true
}

func (controller *CommonFunction) GetPageNumber() int {
	if controller.Ctx.Input.Query("page") != "" {
		pageNumber, pageNumberConvertError := strconv.Atoi(controller.Ctx.Input.Query("page"))
		controller.CheckError(pageNumberConvertError, models.FailParsingMsg, models.FailParsingStatus)
		return pageNumber
	}
	return 0
}
