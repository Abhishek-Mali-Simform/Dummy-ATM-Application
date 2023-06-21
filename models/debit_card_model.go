package models

import (
	"Application/initializers"
)

type DebitCard struct {
	DebitCardID int32   `gorm:"column(debit_card_id);primaryKey;unique" json:"debit_card_id,omitempty"`
	Number      int64   `gorm:"column(number); not null" json:"number,omitempty"`
	Pin         int32   `gorm:"column(pin); not null" json:"pin,omitempty"`
	CVV         int32   `gorm:"column(cvv); not null" json:"cvv,omitempty"`
	ValidFrom   string  `gorm:"column(valid_from);type:date; not null" json:"valid_from,omitempty"`
	ValidUpTo   string  `gorm:"column(valid_up_to);type:date; not null" json:"valid_up_to,omitempty"`
	AccountID   int32   `gorm:"column(account_id); not null" json:"account_id,omitempty"`
	Account     Account `gorm:"references:AccountID;" json:"-"`
}

func (card *DebitCard) CheckCard(res *Response, operation string) bool {
	db := initializers.DB().Where(&card).Preload("Account").First(&card)
	if db.RowsAffected > 0 {
		return true
	}
	*res = DataResponse(
		InvalidCardMsg,
		InvalidDataStatus,
		operation+" "+InvalidCredentials,
	)
	return false
}
