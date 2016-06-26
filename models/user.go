package models

import (
	"encoding/json"
	// "io"
	// "log"
	// "regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/showntop/tripper/errors"
	"github.com/showntop/tripper/utils"
)

type User struct {
	BaseModel         `bson:",inline" jsonapi:"inline,"`
	Username          string `bson:"username" jsonapi:"attr,username"` // 不区分大小写, 全是小写
	Email             string `bson:"email" jsonapi:"attr,email"`       // 全是小写
	Mobile            string `bson:"mobile" jsonapi:"attr,mobile"`     // 全是小写
	Password          string `bson:"-" json:"-"`
	EncryptedPassword string `bson:"encrypted_password" json:"-"`
}

func NewUser(mobile, password string) (*User, errors.SunError) {
	//校验
	// if username == "" {
	// 	re := regexp.MustCompile("[^a-zA-Z^0-9]+")
	// 	username = re.ReplaceAllString(mobile, "")
	// }
	// ok, msg := utils.Vd("username", username)
	// if !ok {
	// 	return nil, errors.NewValidateError(msg)
	// }
	ok, msg := utils.Vd("mobile", mobile)
	if !ok {
		return nil, errors.NewValidateError(msg)
	}
	ok, msg = utils.Vd("password", password)
	if !ok {
		return nil, errors.NewValidateError(msg)
	}
	//加密
	encryptedPassword, err := utils.Encrypt(password)
	if err != nil {
		return nil, errors.NewServerError(msg)
	}
	return &User{
		// Username:          username,
		Mobile:            mobile,
		Password:          password,
		EncryptedPassword: encryptedPassword,
	}, nil
}

// func ParseUserFromJson(data io.Reader) (*User, errors.SunError) {
// 	decoder := json.NewDecoder(data)
// 	var u User
// 	err := decoder.Decode(&u)
// 	if err != nil {
// 		return nil, errors.NewServerError("parse req json error")
// 	}
// 	//validate allowed params key and value and type
// 	return &u, nil
// }

func (u *User) AuthPassword(password string) bool {
	// log.Println(password)
	// log.Println(u.EncryptedPassword)
	return utils.Compare(password, u.EncryptedPassword)
}

func (u *User) PreSave() {
	if u.Id == "" {
		u.Id = bson.NewObjectId()
	}

	if u.Email != "" {
		u.Email = strings.ToLower(u.Email)

		// 发送验证邮箱
		go func() {
			//emailService.RegisterSendActiveEmail(u, u.Email)
			// 发送给我 life@leanote.com
			// emailService.SendEmail("life@leanote.com", "新增用户", "{header}用户名"+u.Email+"{footer}")
		}()
	}

}

func (u *User) PostSave() {

}

func (u *User) ToJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}
