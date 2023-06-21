package models

import (
	"Application/initializers"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
)

type OAUTH2User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Contact  string `json:"contact,omitempty"`
}

func (oauth2User *OAUTH2User) CreateClaimsAndToken(userInfo []byte, getUserDataError error, userType string) (tokenStr []byte, claimID int, err error) {
	if getUserDataError != nil {
		err = getUserDataError
	} else {
		var userData map[string]interface{}
		unmarshalError := json.Unmarshal(userInfo, &userData)
		if unmarshalError != nil {
			err = unmarshalError
		} else {
			var claims Claims
			if userType == "google_user" {
				claims = Claims{
					Email: userData["email"].(string),
					Type:  userType,
				}
			} else if userType == "microsoft_user" {
				contactStr, ok := userData["mobilePhone"].(string)
				if ok {
					contact, _ := strconv.Atoi(contactStr)
					claims = Claims{
						Email:   userData["userPrincipalName"].(string),
						Type:    userType,
						Contact: int64(contact),
					}
				} else {
					claims = Claims{
						Email: userData["userPrincipalName"].(string),
						Type:  userType,
					}
				}
			}
			initializers.DB().Model(User{}).
				Where("email = ?", claims.Email).
				Select("username").
				Find(&claims.Username)
			token, tokenError := claims.GenerateToken()
			if tokenError != nil {
				err = tokenError
				logs.Error(tokenError)
			} else {
				encryptedValue, encryptError := claims.Encrypt([]byte(token), []byte(strconv.Itoa(claims.ClaimID)))
				if encryptError != nil {
					err = encryptError
					logs.Error(encryptError)
				} else {
					tokenStr = encryptedValue
					claimID = claims.ClaimID
				}
			}
		}
	}
	return
}
