package controllers

import (
	"Application/models"
	"encoding/json"
	"strconv"
	"strings"
)

// UserController operations for login, signup, etc
type UserController struct {
	CommonFunction
}

// URLMapping ...
func (controller *UserController) URLMapping() {
	controller.Mapping("SignUpUser", controller.SignUpUser)
	controller.Mapping("SignInUser", controller.SignInUser)
	controller.Mapping("RetrieveUserInfo", controller.RetrieveUserInfo)
}

// SignUpUser ...
// @Title Sign Up User
// @Description User registration with unique email and username
// @Param	body	body models.UserSwagger	true		"body for withdrawing amount using debit card"
// @Success 201 {object} models.UserSwagger
// @Failure 403 body is empty
// @router /sign-up [post]
func (controller *UserController) SignUpUser() {
	if controller.Ctx.Input.IsPost() {
		var userModel models.User
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &userModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			controller.Send(userModel.Create())
		}
	} else {
		controller.InvalidMethod()
	}
}

// SignInUser ...
// @Title Sign In User
// @Description User Login to and existing account using username and password
// @Param	body	body models.UserSwaggerLogin	true		"body for withdrawing amount using debit card"
// @Success 201 Successful Message
// @Failure 403 body is empty
// @router /sign-in [post]
func (controller *UserController) SignInUser() {
	if controller.Ctx.Input.IsPost() {
		var userModel models.User
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &userModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			resp, token, claim := userModel.CheckCredentials()
			if controller.SetCredentials(token, claim) {
				controller.Send(resp)
			}
		}
	} else {
		controller.InvalidMethod()
	}
}

// RetrieveUserInfo ...
// @Title Retrieve User(s) Info
// @Description get User Info based on token, id or all users
// @Param	user_id	query	string	false	"user_id to retrieve user info"
// @Param	token	query	string	false	"token to retrieve user info combined with claim_id"
// @Param	claim_id	query	string	false	"claim_id to retrieve user info combined with token"
// @Param	user_email	query	string	false	"user_email to retrieve user info"
// @Success 200 {object} models.UserSwagger
// @Failure 403
// @router / [get]
func (controller *UserController) RetrieveUserInfo() {
	if controller.Ctx.Input.IsGet() {
		var res models.Response
		if strings.TrimSpace(controller.Ctx.Input.Query("user_id")) != "" {
			userID, convertError := strconv.Atoi(strings.TrimSpace(controller.Ctx.Input.Query("user_id")))
			if !controller.CheckError(convertError, models.FailParsingMsg, models.FailParsingStatus) {
				res = models.RetrieveUserByID(int32(userID))
			}
		} else if strings.TrimSpace(controller.Ctx.Input.Query("token")) != "" && strings.TrimSpace(controller.Ctx.Input.Query("claim_id")) != "" {
			res = models.RetrieveUserByToken(strings.TrimSpace(controller.Ctx.Input.Query("token")),
				strings.TrimSpace(controller.Ctx.Input.Query("claim_id")))
		} else if strings.TrimSpace(controller.Ctx.Input.Query("user_email")) != "" {
			res = models.RetrieveUserByEmail(strings.TrimSpace(controller.Ctx.Input.Query("user_email")))
		} else {
			res = models.RetrieveAllUsers()
		}
		controller.Send(res)
	} else {
		controller.InvalidMethod()
	}
}
