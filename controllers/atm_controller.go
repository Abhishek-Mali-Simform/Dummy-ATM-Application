package controllers

import (
	"Application/models"
	"encoding/json"
)

// ATMController operations for performing withdraw, deposit, check balance
type ATMController struct {
	CommonFunction
}

// URLMapping ...
func (controller *ATMController) URLMapping() {
	controller.Mapping("WithdrawMoneyFromATM", controller.WithdrawMoneyFromATM)
	controller.Mapping("DepositMoneyFromATM", controller.DepositMoneyFromATM)
	controller.Mapping("CheckBalanceFromATM", controller.CheckBalanceFromATM)
}

// WithdrawMoneyFromATM ...
// @Title Withdraw Money
// @Description Withdraw money from an account using ATM and Debit Card
// @Param	body	body models.ATMSwaggerWithAmount	true		"body for withdrawing amount using debit card"
// @Success 201 Success Message
// @Failure 403 body is empty
// @router /withdraw [post]
func (controller *ATMController) WithdrawMoneyFromATM() {
	if controller.Ctx.Input.IsPost() {
		var atmModel models.ATM
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &atmModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			controller.Send(atmModel.Withdraw())
		}
	} else {
		controller.InvalidMethod()
	}
}

// DepositMoneyFromATM ...
// @Title Deposit Money
// @Description Deposit money to an account  using ATM and Debit Card
// @Param	body	body models.ATMSwaggerWithAmount	true		"body for depositing amount using debit card"
// @Success 201 Success Message
// @Failure 403 body is empty
// @router /deposit [post]
func (controller *ATMController) DepositMoneyFromATM() {
	if controller.Ctx.Input.IsPost() {
		var atmModel models.ATM
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &atmModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			controller.Send(atmModel.Deposit())
		}
	} else {
		controller.InvalidMethod()
	}
}

// CheckBalanceFromATM ...
// @Title check balance
// @Description check money from an account via debit card on ATM
// @Param	body	body models.ATMSwaggerWithoutAmount	true		"body for check remaining balance using debit card"
// @Success 201 Success Message
// @Failure 403 body is empty
// @router /check-balance [post]
func (controller *ATMController) CheckBalanceFromATM() {
	if controller.Ctx.Input.IsPost() {
		var atmModel models.ATM
		if !controller.CheckError(
			json.Unmarshal(controller.Ctx.Input.RequestBody, &atmModel),
			models.DataNotFoundMsg,
			models.FailParsingStatus,
		) {
			controller.Send(atmModel.CheckBalance())
		}
	} else {
		controller.InvalidMethod()
	}
}
