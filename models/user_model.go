package models

import (
	"Application/initializers"
	"github.com/beego/beego/v2/core/logs"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"strings"
	"time"
)

type User struct {
	UserID      int32      `gorm:"column(user_id);primaryKey;" json:"user_id,omitempty"`
	Username    string     `gorm:"column(username); type:varchar(255); unique; not null" json:"username,omitempty"`
	Name        string     `gorm:"column(name); type:varchar(255); not null" json:"name,omitempty"`
	Email       string     `gorm:"column(email); type:varchar(255); unique; not null" json:"email,omitempty"`
	Contact     int64      `gorm:"column(contact); not null" json:"contact,omitempty"`
	Password    string     `gorm:"column(password); type:varchar(255); not null" json:"password,omitempty"`
	CreatedDate *time.Time `gorm:"column(create_date); type:timestamp without time zone" json:"created_date,omitempty"`
	UpdatedDate *time.Time `gorm:"column(update_date); type:timestamp without time zone" json:"updated_date,omitempty"`
}

func (user *User) CheckUser(res *Response, operation string) bool {
	db := initializers.DB().
		Where(&user).
		First(&user)
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

func (user *User) Create() (res Response) {
	timeDate := time.Now()
	user.CreatedDate, user.UpdatedDate = &timeDate, &timeDate
	rowCreated := initializers.DB().Create(&user)
	user.Password, user.CreatedDate, user.UpdatedDate = "", nil, nil
	if rowCreated.Error != nil {
		res = DataResponse(
			DataWriteFailMsg,
			DataWriteFailStatus,
			UserSwagger{},
		)
	} else {
		res = DataResponse(
			DataWriteMsg,
			DataWriteStatus,
			user,
		)
	}
	return
}

func (user *User) CheckCredentials() (res Response, tokenStr []byte, claimID int) {
	var (
		claim map[string]interface{}
		resp  = map[string]interface{}{}
	)
	rowMatched := initializers.DB().Model(User{}).
		Where(&user).
		Select("username,contact,email").
		First(&claim)
	if rowMatched.RowsAffected == 1 {
		claims := Claims{
			Username: claim["username"].(string),
			Contact:  claim["contact"].(int64),
			Email:    claim["email"].(string),
			Type:     "user",
		}
		token, tokenError := claims.GenerateToken()
		if tokenError != nil {
			res = DataResponse(
				DataFoundMsg,
				DataFoundStatus,
				tokenError,
			)
			logs.Error(tokenError)
		} else {
			encryptedValue, encryptError := claims.Encrypt([]byte(token), []byte(strconv.Itoa(claims.ClaimID)))
			if encryptError != nil {
				res = DataResponse(
					DataFoundMsg,
					DataFoundStatus,
					strconv.Itoa(claims.ClaimID)+" is too small to proceed",
				)
				logs.Error(encryptError)
			} else {
				resp["token"] = encryptedValue
				resp["claim_id"] = claims.ClaimID
				res = DataResponse(
					DataFoundMsg,
					DataFoundStatus,
					resp,
				)
				tokenStr = encryptedValue
				claimID = claims.ClaimID
			}
		}
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			InvalidCredentials,
		)
	}
	return
}

func RetrieveAllUsers() (res Response) {
	var users []map[string]interface{}
	rowsFound := initializers.DB().Model(User{}).
		Joins("inner join accounts on accounts.user_id = users.user_id").
		Select("username,email,name,contact,users.user_id,number,'user' as type").
		Find(&users)
	if rowsFound.RowsAffected > 0 {
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			users,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
	}
	return
}

func RetrieveUserByID(userID int32) (res Response) {
	var user map[string]interface{}
	rowsFound := initializers.DB().Model(User{}).
		Where("users.user_id = ?", userID).
		Joins("inner join accounts on accounts.user_id = users.user_id").
		Select("username,email,name,contact,users.user_id,number").
		Find(&user)
	if rowsFound.RowsAffected > 0 {
		user["type"] = "user"
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			user,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
	}
	return
}

func RetrieveUserByEmail(userEmail string) (res Response) {
	var user map[string]interface{}
	rowsFound := initializers.DB().Model(User{}).
		Where("email = ?", userEmail).
		Select("user_id").
		Find(&user)
	if rowsFound.RowsAffected > 0 {
		res = DataResponse(
			DataFoundMsg,
			DataFoundStatus,
			user,
		)
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
	}
	return
}

func RetrieveUserByToken(tokenStr, claimID string) (res Response) {
	claims := &Claims{}
	var (
		exist    bool
		byteTemp []byte
	)
	temp := strings.Split(tokenStr[1:len(tokenStr)-1], " ")
	for index, _ := range temp {
		tempInt, _ := strconv.Atoi(temp[index])
		byteTemp = append(byteTemp, uint8(tempInt))

	}
	plaintext, err := claims.Decrypt(byteTemp, []byte(claimID))
	if err == nil {
		token, getTokenError := jwt.ParseWithClaims(string(plaintext), claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		})
		if getTokenError == nil && token.Valid {
			user := make(map[string]interface{}, 0)
			user["email"] = claims.Email
			user["username"] = claims.Username
			user["contact"] = claims.Contact
			user["user_id"] = claims.ClaimID
			user["type"] = claims.Type
			resp := RetrieveUserByEmail(claims.Email)
			mapTemp := resp.Data.(map[string]interface{})
			user["account_number"], exist = GetAccountNumber(mapTemp["user_id"].(int32))
			if exist {
				res = DataResponse(
					DataFoundMsg,
					DataFoundStatus,
					user,
				)
			} else {
				res = DataResponse(
					DataNotFoundMsg,
					DataNotFoundStatus,
					make([]interface{}, 0),
				)
			}
		} else {
			res = DataResponse(
				DataNotFoundMsg,
				DataNotFoundStatus,
				make([]interface{}, 0),
			)
		}
	} else {
		res = DataResponse(
			DataNotFoundMsg,
			DataNotFoundStatus,
			make([]interface{}, 0),
		)
	}
	return
}
