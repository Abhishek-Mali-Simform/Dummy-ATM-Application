package models

import "fmt"

var (
	WithdrawLimit int32 = 25000
	WithdrawNote1 int32 = 100
	WithdrawNote2 int32 = 50
	DepositLimit  int32 = 50000
	DepositNote1  int32 = 50
	DepositNote2  int32 = 10
)

type ATM struct {
	DebitCard DebitCard `json:"debit_card"`
	Amount    int32     `json:"amount,omitempty"`
}

func (atm *ATM) Withdraw() (res Response) {
	if !atm.GetAmountErrorResponse(&res, "Withdraw", WithdrawNote1, WithdrawNote2, WithdrawLimit) {

		if atm.DebitCard.CheckCard(&res, "Withdraw") {

			if atm.DebitCard.Account.Balance > float64(atm.Amount) {

				res = atm.DebitCard.Account.SetBalance(
					atm.DebitCard.Account.Balance-float64(atm.Amount),
					"Withdraw",
					float64(atm.Amount),
				)

			} else {

				res = DataResponse(
					InvalidAmountMsg,
					InvalidDataStatus,
					"Withdraw "+InsufficientFunds,
				)

			}

		}
	}
	return
}

func (atm *ATM) GetAmountErrorResponse(res *Response, operation string, note1, note2, amountLimit int32) bool {
	if atm.Amount < 0 {

		*res = DataResponse(
			InvalidAmountMsg,
			InvalidDataStatus,
			operation+" "+NegativeAmount,
		)
		return true

	} else if atm.Amount%note1 != 0 || atm.Amount%note2 != 0 {

		*res = DataResponse(
			InvalidAmountMsg,
			InvalidDataStatus,
			operation+" Unsuccessful Amount Entered should be multiple of "+string(note1)+" or "+string(note2),
		)
		return true

	} else if atm.Amount > amountLimit {

		*res = DataResponse(
			InvalidAmountMsg,
			InvalidDataStatus,
			operation+" "+ExceededLimit,
		)
		return true

	}
	return false
}

func (atm *ATM) Deposit() (res Response) {
	if !atm.GetAmountErrorResponse(&res, "Deposit", DepositNote1, DepositNote2, DepositLimit) {

		if atm.DebitCard.CheckCard(&res, "Deposit") {

			res = atm.DebitCard.Account.SetBalance(
				atm.DebitCard.Account.Balance+float64(atm.Amount),
				"Deposit",
				float64(atm.Amount),
			)

		}
	}
	return
}

func (atm *ATM) CheckBalance() (res Response) {
	if atm.DebitCard.CheckCard(&res, "Checking Balance") {

		res = DataResponse(
			DataWriteMsg,
			DataWriteStatus,
			"Current Balance in Your Account: "+fmt.Sprintf("%v", int(atm.DebitCard.Account.Balance)),
		)

	}
	return
}
