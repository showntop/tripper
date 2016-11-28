package models

import ()

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Avatar      string `json:"avatar"`
	Age         int    `json:"age"`
	Gender      int    `json:"gender"`
	Description string `json:"description"`
}
