package models

type Users struct {
	UserId      int64  `json:"userId"`
	UserCode    string `json:"userCode"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Email       string `json:"email"`
}
