package models

import "time"

type ATMSwaggerWithAmount struct {
	DebitCard DebitCardSwagger `json:"debit_card"`
	Amount    int32            `json:"amount,omitempty"`
}

type ATMSwaggerWithoutAmount struct {
	DebitCard DebitCardSwagger `json:"debit_card"`
}

type DebitCardSwagger struct {
	Number    int64  `json:"number,omitempty"`
	Pin       int32  `json:"pin,omitempty"`
	CVV       int32  `json:"cvv,omitempty"`
	ValidUpTo string `json:"valid_up_to,omitempty"`
}

type AccountSwaggerWithUser struct {
	Number int              `json:"number,omitempty"`
	Type   string           `json:"type,omitempty"`
	User   UserSwaggerLogin `json:"user,omitempty"`
}

type AccountSwaggerWithoutUser struct {
	Number int    `json:"number,omitempty"`
	Type   string `json:"type,omitempty"`
}

type UserSwaggerLogin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type BankSwaggerWithAmount struct {
	Account AccountSwaggerWithUser `json:"account"`
	Amount  int32                  `json:"amount,omitempty"`
}

type BankSwaggerWithoutAmount struct {
	Account AccountSwaggerWithoutUser `json:"account"`
}

type AccountInfoSwagger struct {
	Name                string     `json:"name,omitempty"`
	Number              int        `json:"account_number,omitempty"`
	Balance             float64    `json:"balance,omitempty"`
	LastTransactionDate *time.Time `json:"last_transaction_date,omitempty"`
	Type                string     `json:"account_type,omitempty"`
}

type UserSwagger struct {
	UserID   int32  `json:"user_id,omitempty"`
	Username string `json:"username"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Contact  int64  `json:"contact,omitempty"`
	Password string `json:"password"`
}

type AccountSwagger struct {
	Type    string  `json:"type,omitempty"`
	UserID  int32   `json:"user_id,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}
