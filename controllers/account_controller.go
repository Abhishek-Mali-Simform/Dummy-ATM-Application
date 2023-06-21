package controllers

import (
	"Application/models"
	"encoding/json"
)

// AccountController operations for performing registration, view balance, etc
type AccountController struct {
	CommonFunction
}

// URLMapping ...
func (controller *AccountController) URLMapping() {
	controller.Mapping("WithdrawMoneyFromATM", controller.RegisterAccount)
}

// RegisterAccount ...
// @Title Sign Up User
// @Description User registration with unique email and username
// @Param	body	body models.AccountSwagger	true		"body for withdrawing amount using debit card"
// @Success 201 {object} models.AccountInfoSwagger
// @Failure 403 body is empty
// @router /register [post]
func (controller *AccountController) RegisterAccount() {
	if controller.Ctx.Input.IsPost() {
		var accountModel models.Account
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &accountModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			controller.Send(accountModel.Create())
		}
	} else {
		controller.InvalidMethod()
	}
}
