package models

import (
	"Application/initializers"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	TransactionID int32      `gorm:"column(transaction_id);primaryKey;" json:"transaction_id,omitempty"`
	AccountID     int32      `gorm:"column(account_id); not null" json:"account_id,omitempty"`
	Operation     string     `gorm:"column(operation); type:varchar(255); not null" json:"operation,omitempty"`
	Amount        float64    `gorm:"column(amount); not null" json:"amount,omitempty"`
	DateTime      *time.Time `gorm:"column(date_time); type:timestamp without time zone" json:"date_time,omitempty"`
	Account       Account    `gorm:"references:AccountID;" json:"account,omitempty"`
}

func (transaction *Transaction) Create() bool {
	dateTime := time.Now()
	transaction.DateTime = &dateTime
	createdRow := initializers.DB().Create(transaction)
	if createdRow.Error == nil {
		return true
	}
	logs.Error(createdRow.Error)
	return false
}

func RetrieveTransactionByAccountNumber(accountNumber, pageNumber int) (res Response) {
	var (
		result     []map[string]interface{}
		pagination Pagination
	)
	pagination.Page = pageNumber
	rowsFound := initializers.DB().Model(Transaction{}).
		Joins("inner join accounts on accounts.account_id = transactions.account_id").
		Where("accounts.number = ?", accountNumber).
		Select("accounts.number,operation,amount,date_time").
		Count(&pagination.TotalRows)
	pagination.Limit = 10
	pagination.GetTotalPages()
	pagination.Sort = "date_time desc"
	rowsFound.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}).Find(&result)
	if rowsFound.RowsAffected < 1 {
		res = DataResponse(DataNotFoundMsg, DataNotFoundStatus, make([]interface{}, 0))
	} else {
		pagination.Rows = result
		res = DataResponse(DataFoundMsg, DataFoundStatus, pagination)
	}
	return
}

func RetrieveTransactionReportByAccountNumber(accountNumber int) (res Response) {
	var result = map[string]interface{}{}
	result["Withdraw"] = GetTransactionReport(accountNumber, "Withdraw")
	result["Deposit"] = GetTransactionReport(accountNumber, "Deposit")
	res = DataResponse(DataFoundMsg, DataFoundStatus, result)
	return
}
func GetTransactionReport(accountNumber int, operation string) (result []map[string]interface{}) {
	initializers.DB().Model(Transaction{}).
		Joins("inner join accounts on accounts.account_id = transactions.account_id").
		Where("accounts.number = ? AND transactions.operation = ? AND date_part('year', transactions.date_time) = ?", accountNumber, operation, time.Now().Year()).
		Select("to_char(transactions.date_time,'month'), avg(transactions.amount)").
		Group("to_char(transactions.date_time,'month')").
		Find(&result)
	return
}
