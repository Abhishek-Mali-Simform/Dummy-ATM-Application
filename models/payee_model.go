package models

import "Application/initializers"

type Payee struct {
	PayeeID        int32 `gorm:"column(payee_id);primaryKey;" json:"payee_id,omitempty"`
	AccountID      int32 `gorm:"column(account_id); not null;" json:"account_id,omitempty"`
	PayeeAccountID int32 `gorm:"column(payee_account_id); not null;" json:"payee_account_id,omitempty"`
}

func (payee *Payee) Create() (res Response) {
	initializers.DB().Where(&payee).Find(&payee)
	if payee.PayeeID == 0 {
		initializers.DB().Create(&payee)
		if payee.PayeeID != 0 {
			res = DataResponse(DataWriteMsg, DataWriteStatus, payee)
		} else {
			res = DataResponse(DataWriteFailMsg, DataWriteFailStatus, make([]interface{}, 0))
		}
	} else {
		res = DataResponse(DataFoundMsg, DataFoundStatus, &payee)
	}
	return
}

func RetrieveAllPayeeByUserAccountNumber(accountNumber int) (res Response) {
	var (
		accountID int32
		result    []map[string]interface{}
	)
	initializers.DB().Model(Account{}).
		Where("accounts.number = ?", accountNumber).
		Select("account_id").
		Find(&accountID)
	rows := initializers.DB().Model(Payee{}).
		Joins("inner join accounts on payees.payee_account_id = accounts.account_id").
		Joins("inner join users on users.user_id = accounts.user_id").
		Where("payees.account_id = ?", accountID).
		Select("accounts.number, users.name, users.user_id, users.password,accounts.type").
		Find(&result)
	if rows.RowsAffected > 0 {
		res = DataResponse(DataFoundMsg, DataFoundStatus, result)
	} else {
		res = DataResponse(DataNotFoundMsg, DataNotFoundStatus, make([]interface{}, 0))
	}
	return
}
