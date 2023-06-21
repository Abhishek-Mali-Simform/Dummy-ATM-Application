package controllers

import (
	"Application/models"
	"encoding/json"
	"strconv"
	"strings"
)

type PayeeController struct {
	CommonFunction
}

// URLMapping ...
func (controller *PayeeController) URLMapping() {
	controller.Mapping("RegisterPayee", controller.RegisterPayee)
	controller.Mapping("RetrieveAllPayee", controller.RetrieveAllPayee)
}

// RegisterPayee ...
// @Title Register Payee
// @Description Payee registration with account number
// @Param	body	body models.PayeeWithAccountNumber	true		"body for adding payee"
// @Success 201 {object} models.Payee
// @Failure 403 body is empty
// @router / [post]
func (controller *PayeeController) RegisterPayee() {
	if controller.Ctx.Input.IsPost() {
		var payeeModel models.PayeeWithAccountNumber
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &payeeModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			controller.Send(payeeModel.AddPayee())
		}
	} else {
		controller.InvalidMethod()
	}
}

// RetrieveAllPayee ...
// @Title Retrieve All Payee
// @Description get all payee with account number
// @Param	account_number	query	string	false	"account_number to retrieve payee info"
// @Success 201 {object} models.Payee
// @Failure 403 body is empty
// @router / [get]
func (controller *PayeeController) RetrieveAllPayee() {
	if controller.Ctx.Input.IsGet() {
		if strings.TrimSpace(controller.Ctx.Input.Query("account_number")) != "" {
			accountNumber, convertError := strconv.Atoi(strings.TrimSpace(controller.Ctx.Input.Query("account_number")))
			if !controller.CheckError(convertError, models.FailParsingMsg, models.FailParsingStatus) {
				controller.Send(models.RetrieveAllPayeeByUserAccountNumber(accountNumber))
			}
		} else {
			controller.InvalidMethod()
		}
	} else {
		controller.InvalidMethod()
	}
}
