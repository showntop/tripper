package models

import ()

type User struct {
	Id          int64
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Avatar      string `json:"avatar"`
	Age         int    `json:"age"`
	Gender      int    `json:"gender"`
	Description string `json:"description"`
}
