package controllers

import (
	"Application/models"
	"strconv"
	"strings"
)

// TransactionController operations for getting transactions
type TransactionController struct {
	CommonFunction
}

// URLMapping ...
func (controller *TransactionController) URLMapping() {
	controller.Mapping("RetrieveTransactionHistory", controller.RetrieveTransactionHistory)
	controller.Mapping("RetrieveTransactionSummary", controller.RetrieveTransactionSummary)
}

// RetrieveTransactionHistory ...
// @Title Retrieve Transaction(s) History Info
// @Description get Transaction(s) History based on account number
// @Param	account_number	query	string	false	"account_number to retrieve transaction info"
// @Param	page	query	string	false	"Page number fo pagination. Must be an integer"
// @Success 200 {object} models.UserSwagger
// @Failure 403
// @router / [get]
func (controller *TransactionController) RetrieveTransactionHistory() {
	if controller.Ctx.Input.IsGet() {
		if strings.TrimSpace(controller.Ctx.Input.Query("account_number")) != "" {
			accountNumber, convertError := strconv.Atoi(strings.TrimSpace(controller.Ctx.Input.Query("account_number")))
			if !controller.CheckError(convertError, models.FailParsingMsg, models.FailParsingStatus) {
				controller.Send(models.RetrieveTransactionByAccountNumber(accountNumber, controller.GetPageNumber()))
			}
		} else {
			controller.InvalidMethod()
		}
	} else {
		controller.InvalidMethod()
	}
}

// RetrieveTransactionSummary ...
// @Title Retrieve Transaction(s) Summary Info
// @Description get Transaction(s) Summary based on account number
// @Param	account_number	query	string	false	"account_number to retrieve transaction info"
// @Success 200 {object} models.UserSwagger
// @Failure 403
// @router /report [get]
func (controller *TransactionController) RetrieveTransactionSummary() {
	if controller.Ctx.Input.IsGet() {
		if strings.TrimSpace(controller.Ctx.Input.Query("account_number")) != "" {
			accountNumber, convertError := strconv.Atoi(strings.TrimSpace(controller.Ctx.Input.Query("account_number")))
			if !controller.CheckError(convertError, models.FailParsingMsg, models.FailParsingStatus) {
				controller.Send(models.RetrieveTransactionReportByAccountNumber(accountNumber))
			}
		} else {
			controller.InvalidMethod()
		}
	} else {
		controller.InvalidMethod()
	}
}
