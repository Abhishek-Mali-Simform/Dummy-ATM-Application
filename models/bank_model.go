package models

import "fmt"

type Bank struct {
	Account Account `json:"account"`
	Amount  int32   `json:"amount,omitempty"`
}

func (bank *Bank) GetAmountErrorResponse(res *Response, operation string, note1, note2, amountLimit int32) bool {
	if bank.Amount < 0 {

		*res = DataResponse(
			InvalidAmountMsg,
			InvalidDataStatus,
			operation+" "+NegativeAmount,
		)
		return true

	} else if bank.Amount%note1 != 0 || bank.Amount%note2 != 0 {

		*res = DataResponse(
			InvalidAmountMsg,
			InvalidDataStatus,
			operation+" Unsuccessful Amount Entered should be multiple of "+string(note1)+" or "+string(note2),
		)
		return true

	} else if bank.Amount > amountLimit {

		*res = DataResponse(
			InvalidAmountMsg,
			InvalidDataStatus,
			operation+" "+ExceededLimit,
		)
		return true

	}
	return false
}

func (bank *Bank) Withdraw() (res Response) {
	if !bank.GetAmountErrorResponse(&res, "Withdraw", WithdrawNote1, WithdrawNote2, WithdrawLimit) {

		if bank.Account.CheckAccount(&res, "Withdraw") {

			if bank.Account.Balance > float64(bank.Amount) {

				res = bank.Account.SetBalance(
					bank.Account.Balance-float64(bank.Amount),
					"Withdraw",
					float64(bank.Amount),
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

func (bank *Bank) Deposit() (res Response) {
	if !bank.GetAmountErrorResponse(&res, "Deposit", DepositNote1, DepositNote2, DepositLimit) {

		if bank.Account.CheckAccount(&res, "Deposit") {

			res = bank.Account.SetBalance(
				bank.Account.Balance+float64(bank.Amount),
				"Deposit",
				float64(bank.Amount),
			)

		}
	}
	return
}

func (bank *Bank) CheckBalance() (res Response) {
	if bank.Account.CheckAccount(&res, "Checking Balance") {

		res = DataResponse(
			DataWriteMsg,
			DataWriteStatus,
			"Current Balance in Your Account: "+fmt.Sprintf("%v", int(bank.Account.Balance)),
		)

	}
	return
}
