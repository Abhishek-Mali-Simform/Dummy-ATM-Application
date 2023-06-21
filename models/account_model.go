package models

import (
	"Application/initializers"
	"Application/services"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type Account struct {
	AccountID   int32      `gorm:"column(account_id);primaryKey;" json:"account_id,omitempty"`
	Number      int        `gorm:"column(number) not null" json:"number,omitempty"`
	Type        string     `gorm:"column(type); type:varchar(255); not null" json:"type,omitempty"`
	UserID      int32      `gorm:"column(user_id); not null" json:"user_id,omitempty"`
	Balance     float64    `gorm:"column(balance); not null" json:"balance,omitempty"`
	CreatedDate *time.Time `gorm:"column(create_date); type:timestamp without time zone" json:"created_date,omitempty"`
	UpdatedDate *time.Time `gorm:"column(update_date); type:timestamp without time zone" json:"updated_date,omitempty"`
	User        User       `gorm:"references:UserID;" json:"user,omitempty"`
	Payer       Payee      `gorm:"foreignKey:AccountID;" json:"account,omitempty"`
	Payee       Payee      `gorm:"foreignKey:PayeeAccountID;" json:"payee_account,omitempty"`
}

func (account *Account) SetBalance(balance float64, operation string, amount float64) (res Response) {
	temporaryAssignBalance := account.Balance
	update := initializers.DB().
		Model(&Account{}).
		Where("account_id = ?", account.AccountID).
		Update("balance", balance)
	services.CheckError("models/account_model.go", 20, update.Error)
	if update.RowsAffected > 0 {
		transaction := Transaction{
			AccountID: account.AccountID,
			Operation: operation,
			Amount:    amount,
		}
		if transaction.Create() {

			res = DataResponse(
				DataWriteMsg,
				DataWriteStatus,
				operation+" "+Success,
			)

		} else {
			initializers.DB().
				Model(&Account{}).
				Where("account_id = ?", account.AccountID).
				Update("balance", temporaryAssignBalance)
			res = DataResponse(
				DataWriteMsg,
				DataWriteStatus,
				operation+" "+UnsuccessfulTransaction,
			)
		}

	} else {

		res = DataResponse(
			DataWriteMsg,
			DataWriteStatus,
			operation+" "+Unsuccessful,
		)

	}
	return
}

func (account *Account) CheckAccount(res *Response, operation string) bool {
	if account.User.CheckUser(res, operation) {
		db := initializers.DB().Where(&account).First(&account)
		if db.RowsAffected > 0 {
			return true
		}
		*res = DataResponse(
			InvalidCardMsg,
			InvalidDataStatus,
			operation+" "+InvalidCredentials,
		)
	}
	return false
}

func AllAccountInfo() (res Response) {
	accounts := make([]map[string]interface{}, 0)
	result := initializers.DB().Model(Account{}).
		Joins("inner join users on users.user_id = accounts.user_id").
		Joins("inner join (select account_id,max(date_time) as last_transaction_date from transactions group by account_id) as transaction on transaction.account_id = accounts.account_id").
		Select("name,number as account_number,balance,type as account_type,last_transaction_date").
		Find(&accounts)
	if result.RowsAffected > 0 {
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			accounts,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
		logs.Error(result.Error)
	}
	return
}

func (account *Account) Create() (res Response) {
	dateTime := time.Now()
	account.CreatedDate, account.UpdatedDate = &dateTime, &dateTime
	account.Number = services.GetRandomNumber(10)
	rowCreated := initializers.DB().Create(&account)
	account.CreatedDate, account.UpdatedDate = nil, nil
	if rowCreated.Error != nil {
		res = DataResponse(
			DataWriteFailMsg,
			DataWriteFailStatus,
			AccountInfoSwagger{},
		)

	} else {
		transaction := Transaction{
			AccountID: account.AccountID,
			Operation: "Initial Balance",
			Amount:    account.Balance,
		}
		if transaction.Create() {

			res = DataResponse(
				DataWriteMsg,
				DataWriteStatus,
				account,
			)

		} else {
			initializers.DB().
				Where("account_id = ?", account.AccountID).
				Delete(&account)
			res = DataResponse(
				DataWriteMsg,
				DataWriteStatus,
				"Creating Account "+UnsuccessfulTransaction,
			)
		}
		res = DataResponse(
			DataWriteMsg,
			DataWriteStatus,
			account,
		)
	}
	return
}

func GetAccountInfo(accountNumber int) (res Response) {
	accounts := make([]map[string]interface{}, 0)
	result := initializers.DB().Model(Account{}).
		Joins("inner join users on users.user_id = accounts.user_id").
		Joins("inner join (select account_id,max(date_time) as last_transaction_date from transactions group by account_id) as transaction on transaction.account_id = accounts.account_id").
		Select("name,number as account_number,balance,type as account_type,last_transaction_date").
		Where("accounts.number = ?", accountNumber).
		Find(&accounts)
	if result.RowsAffected > 0 {
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			accounts,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
		logs.Error(result.Error)
	}
	return
}

func GetAccountNumber(userID int32) (number int, exist bool) {
	initializers.DB().Model(Account{}).Where("user_id = ?", userID).Select("number").Find(&number)
	if number != 0 {
		exist = true
	}
	return
}

func GetAccount(accountNumber int) (accountDetails map[string]interface{}) {
	initializers.DB().Model(Account{}).
		Joins("inner join users on users.user_id = accounts.user_id").
		Where("accounts.number = ?", accountNumber).
		Select("accounts.account_id,users.name").
		Find(&accountDetails)
	return
}

func GetCardInfo(accountNumber int) (res Response) {
	accounts := make([]map[string]interface{}, 0)
	result := initializers.DB().Model(Account{}).
		Joins("inner join debit_cards on debit_cards.account_id = debit_cards.account_id").
		Select("debit_cards.number as debit_card_number,accounts.number as account_number,cvv ,valid_from,valid_up_to").
		Where("accounts.number = ?", accountNumber).
		Find(&accounts)
	if result.RowsAffected > 0 {
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			accounts,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
		logs.Error(result.Error)
	}
	return
}

func GetAllCardInfo() (res Response) {
	accounts := make([]map[string]interface{}, 0)
	result := initializers.DB().Model(Account{}).
		Joins("inner join debit_cards on debit_cards.account_id = debit_cards.account_id").
		Select("debit_cards.number as debit_card_number,accounts.number as account_number,cvv ,valid_from,valid_up_to").
		Find(&accounts)
	if result.RowsAffected > 0 {
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			accounts,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
		logs.Error(result.Error)
	}
	return
}
