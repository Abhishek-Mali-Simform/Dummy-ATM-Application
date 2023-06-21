package models

type PayeeWithAccountNumber struct {
	UserAccountNumber  int
	PayeeAccountNumber int
	PayeeName          string
}

func (payeeWithAccountNumber *PayeeWithAccountNumber) AddPayee() (res Response) {
	var (
		check bool
		payee Payee
	)
	if payeeWithAccountNumber.PayeeAccountNumber == payeeWithAccountNumber.UserAccountNumber {
		res = DataResponse("Idiot You Added ypur Account Number", DataFoundStatus, make([]interface{}, 0))
	} else {
		mapUserAccount := GetAccount(payeeWithAccountNumber.UserAccountNumber)
		if mapUserAccount != nil {
			payee.AccountID, check = mapUserAccount["account_id"].(int32)
			if check {
				mapPayeeAccount := GetAccount(payeeWithAccountNumber.PayeeAccountNumber)
				if mapPayeeAccount != nil {
					if payeeWithAccountNumber.PayeeName == mapPayeeAccount["name"].(string) {
						payee.PayeeAccountID, check = mapPayeeAccount["account_id"].(int32)
						if check {
							res = payee.Create()
						} else {
							res = DataResponse(DataWriteFailMsg, DataWriteFailStatus, make([]interface{}, 0))
						}
					} else {
						res = DataResponse(DataWriteFailMsg, DataWriteFailStatus, make([]interface{}, 0))
					}
				} else {
					res = DataResponse(DataWriteFailMsg, DataWriteFailStatus, make([]interface{}, 0))
				}
			} else {
				res = DataResponse(DataWriteFailMsg, DataWriteFailStatus, make([]interface{}, 0))
			}
		} else {
			res = DataResponse(DataWriteFailMsg, DataWriteFailStatus, make([]interface{}, 0))
		}
	}
	return
}
