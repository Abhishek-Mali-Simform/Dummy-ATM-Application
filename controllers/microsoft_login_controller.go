package controllers

import (
	"Application/models"
	"github.com/Abhishek-Mali-Simform/oauth2"
)

type MicrosoftLoginController struct {
	CommonFunction
}

// URLMapping ...
func (controller *MicrosoftLoginController) URLMapping() {
	controller.Mapping("MicrosoftLogin", controller.MicrosoftLogin)
	controller.Mapping("MicrosoftCallback", controller.MicrosoftCallback)
}

// MicrosoftLogin ...
// @Title Strictly Not For Swagger Use
// @Description get User Info from Microsoft
// @Success 200 Success Message
// @Failure 403
// @router /login [get]
func (controller *MicrosoftLoginController) MicrosoftLogin() {
	if controller.Ctx.Input.IsGet() {
		controller.Redirect(
			oauth2.Login(
				controller.GetMicrosoftConfig(),
			),
			307,
		)

	} else {
		controller.InvalidMethod()
	}
}

// MicrosoftCallback ...
// @Title Strictly Not For Swagger Use
// @Description get User Info from Microsoft
// @Success 200 Success Message
// @Failure 403
// @router /callback [get]
func (controller *MicrosoftLoginController) MicrosoftCallback() {
	if controller.Ctx.Input.IsGet() {
		var oauth2UserModel models.OAUTH2User
		userData, userDataError := oauth2.Callback(
			controller.Ctx.Request,
			controller.GetMicrosoftConfig(),
			oauth2.AuthorizationBearer,
		)
		token, claim, err := oauth2UserModel.CreateClaimsAndToken(
			userData,
			userDataError,
			"microsoft_user",
		)
		if !controller.CheckError(err, models.FailParsingMsg, models.FailParsingStatus) {
			if controller.SetCredentials(token, claim) {
				controller.Redirect("/world-bank/dashboard", 307)
			}
		}
	} else {
		controller.InvalidMethod()
	}
}
