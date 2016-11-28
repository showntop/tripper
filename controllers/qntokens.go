package controllers

import (
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

type Qntokens struct {
	application
}

func (q *Qntokens) Create() ([]byte, *HttpError) {
	//初始化AK，SK
	conf.ACCESS_KEY = "FKj0IPQwM3gquMCFBSQjYI_1EGxb-8sASTh30--U"
	conf.SECRET_KEY = "t938V-cIGD6GFZWhS2VxsVWmZ1S06lWVrFKgKtsS"

	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: "tripper-test",
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)
	return WrapResp(map[string]string{
		"uptoken":    token,
		"expired_at": "10",
	})
}
