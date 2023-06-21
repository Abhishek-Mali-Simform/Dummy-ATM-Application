package controllers

import (
	"Application/models"
	"encoding/json"
	"strconv"
	"strings"
)

// BankController operations for performing withdraw, deposit, check balance
type BankController struct {
	CommonFunction
}

// URLMapping ...
func (controller *BankController) URLMapping() {
	controller.Mapping("WithdrawMoneyFromBank", controller.WithdrawMoneyFromBank)
	controller.Mapping("DepositMoneyFromBank", controller.DepositMoneyFromBank)
	controller.Mapping("CheckBalanceFromBank", controller.CheckBalanceFromBank)
	controller.Mapping("RetrieveAccountInfo", controller.RetrieveAccountInfo)
	controller.Mapping("RetrieveCardInfo", controller.RetrieveCardInfo)
}

// WithdrawMoneyFromBank ...
// @Title Withdraw Money
// @Description Withdraw money from an account using Bank Account
// @Param	body	body models.BankSwaggerWithAmount	true		"body for withdrawing amount using bank account"
// @Param	user_id	query	string	false	"get user id. Should only be integer"
// @Success 201 Success Message
// @Failure 403 body is empty
// @router /withdraw [post]
func (controller *BankController) WithdrawMoneyFromBank() {
	if controller.Ctx.Input.IsPost() {
		if userID := controller.GetID("user_id"); userID != 0 {
			bankModel := models.Bank{
				Account: models.Account{UserID: userID},
			}
			if !controller.CheckError(
				json.Unmarshal(controller.Ctx.Input.RequestBody, &bankModel),
				models.DataNotFoundMsg,
				models.FailParsingStatus,
			) {
				controller.Send(bankModel.Withdraw())
			}
		}
	} else {
		controller.InvalidMethod()
	}
}

// DepositMoneyFromBank ...
// @Title Deposit Money
// @Description Deposit money to an account  using Bank Account
// @Param	body	body models.BankSwaggerWithAmount	true		"body for depositing amount using bank account"
// @Param	user_id	query	string	false	"get user id. Should only be integer"
// @Success 201 Success Message
// @Failure 403 body is empty
// @router /deposit [post]
func (controller *BankController) DepositMoneyFromBank() {
	if controller.Ctx.Input.IsPost() {
		if userID := controller.GetID("user_id"); userID != 0 {
			bankModel := models.Bank{
				Account: models.Account{UserID: userID},
			}
			if !controller.CheckError(
				json.Unmarshal(controller.Ctx.Input.RequestBody, &bankModel),
				models.DataNotFoundMsg,
				models.FailParsingStatus,
			) {
				controller.Send(bankModel.Deposit())
			}
		}
	} else {
		controller.InvalidMethod()
	}
}

// CheckBalanceFromBank ...
// @Title check balance
// @Description check money from an account via bank account
// @Param	body	body models.BankSwaggerWithoutAmount	true		"body for check remaining balance using bank account"
// @Param	user_id	query	string	false	"get user id. Should only be integer"
// @Success 201 Success Message
// @Failure 403 body is empty
// @router /check-balance [post]
func (controller *BankController) CheckBalanceFromBank() {
	if controller.Ctx.Input.IsPost() {
		if userID := controller.GetID("user_id"); userID != 0 {
			bankModel := models.Bank{
				Account: models.Account{UserID: userID},
			}
			if !controller.CheckError(
				json.Unmarshal(controller.Ctx.Input.RequestBody, &bankModel),
				models.DataNotFoundMsg,
				models.FailParsingStatus,
			) {
				controller.Send(bankModel.CheckBalance())
			}
		}
	} else {
		controller.InvalidMethod()
	}
}

// RetrieveAccountInfo ...
// @Title Retrieve All Account Info
// @Description get Account Info based on role
// @Param	account_number	query	string	false	"account_number to retrieve account info"
// @Success 200 {object} models.AccountInfoSwagger
// @Failure 403
// @router /accounts-info [get]
func (controller *BankController) RetrieveAccountInfo() {
	if controller.Ctx.Input.IsGet() {
		var res models.Response
		if strings.TrimSpace(controller.Ctx.Input.Query("account_number")) != "" {
			accountNumber, convertError := strconv.Atoi(strings.TrimSpace(controller.Ctx.Input.Query("account_number")))
			if !controller.CheckError(convertError, models.FailParsingMsg, models.FailParsingStatus) {
				res = models.GetAccountInfo(accountNumber)
			}
		} else {
			res = models.AllAccountInfo()
		}
		controller.Send(res)
	} else {
		controller.InvalidMethod()
	}
}

// RetrieveCardInfo ...
// @Title Retrieve All Card Info
// @Description get Card Info based on role
// @Param	account_number	query	string	false	"account_number to retrieve card info"
// @Success 200 {object} models.AccountInfoSwagger
// @Failure 403
// @router /cards-info [get]
func (controller *BankController) RetrieveCardInfo() {
	if controller.Ctx.Input.IsGet() {
		var res models.Response
		if strings.TrimSpace(controller.Ctx.Input.Query("account_number")) != "" {
			accountNumber, convertError := strconv.Atoi(strings.TrimSpace(controller.Ctx.Input.Query("account_number")))
			if !controller.CheckError(convertError, models.FailParsingMsg, models.FailParsingStatus) {
				res = models.GetCardInfo(accountNumber)
			}
		} else {
			res = models.GetAllCardInfo()
		}
		controller.Send(res)
	} else {
		controller.InvalidMethod()
	}
}
