package main

import (
	"Application/initializers"
	"Application/models"
	"Application/services"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	services.CheckError(
		"migrations/migrations.go",
		14,
		initializers.DB().Migrator().DropTable(&models.User{}, &models.Account{}, &models.DebitCard{}, &models.Transaction{}, &models.Payee{}),
	)
	services.CheckError(
		"migrations/migrations.go",
		19,
		initializers.DB().AutoMigrate(&models.User{}, &models.Account{}, &models.DebitCard{}, &models.Transaction{}),
	)
	services.CheckError(
		"migrations/migrations.go",
		25,
		initializers.DB().AutoMigrate(&models.Payee{}),
	)
	t := time.Now()
	initializers.DB().Create(&models.User{
		Username:    "abhi",
		Name:        "Abhishek Mali",
		Email:       "abhishek.m@simformsolutions.com",
		Contact:     9429865212,
		Password:    "1234",
		CreatedDate: &t,
		UpdatedDate: &t,
	})
	var user models.User
	initializers.DB().First(&user)
	initializers.DB().Create(&models.Account{
		Number:      services.GetRandomNumber(10),
		Type:        "Savings",
		UserID:      user.UserID,
		Balance:     5000000000000000.00,
		CreatedDate: &t,
		UpdatedDate: &t,
	})
	var acc models.Account
	//d := t.AddDate(5, 0, 0)
	initializers.DB().First(&acc)
	initializers.DB().Create(&models.DebitCard{
		Number:    4361522182630242,
		Pin:       1234,
		CVV:       221,
		ValidFrom: "2023-05-04",
		ValidUpTo: "2028-05-04",
		AccountID: acc.AccountID,
	})
	initializers.DB().Create(&models.Transaction{
		AccountID: acc.AccountID,
		Operation: "Initial Balance",
		Amount:    5000000000000000.00,
		DateTime:  &t,
	})
}
