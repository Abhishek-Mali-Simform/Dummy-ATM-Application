package controllers

import (
	"Application/models"
	"github.com/Abhishek-Mali-Simform/oauth2"
)

type GoogleLoginController struct {
	CommonFunction
}

// URLMapping ...
func (controller *GoogleLoginController) URLMapping() {
	controller.Mapping("GoogleLogin", controller.GoogleLogin)
	controller.Mapping("GoogleCallback", controller.GoogleCallback)
}

// GoogleLogin ...
// @Title Strictly Not For Swagger Use
// @Description get User Info from Google
// @Success 200 Success Message
// @Failure 403
// @router /login [get]
func (controller *GoogleLoginController) GoogleLogin() {
	if controller.Ctx.Input.IsGet() {
		controller.Redirect(
			oauth2.Login(
				controller.GetGoogleConfig(),
			),
			307,
		)

	} else {
		controller.InvalidMethod()
	}
}

// GoogleCallback ...
// @Title Strictly Not For Swagger Use
// @Description get User Info from Google
// @Success 200 Success Message
// @Failure 403
// @router /callback [get]
func (controller *GoogleLoginController) GoogleCallback() {
	if controller.Ctx.Input.IsGet() {
		var oauth2UserModel models.OAUTH2User
		userData, userDataError := oauth2.Callback(
			controller.Ctx.Request,
			controller.GetGoogleConfig(),
			oauth2.QueryString,
		)
		token, claim, err := oauth2UserModel.CreateClaimsAndToken(
			userData,
			userDataError,
			"google_user",
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
