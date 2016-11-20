package models

import ()

type Post struct {
	Base
	ProjectId  int            `json:"project_id"`
	Content    string         `json:"content"`
	UserId     int            `json:"user_id"`
	Comments   []*PostComment `json:"comments"`
	LikeNum    int            `json:"like_num"`
	CommentNum int            `json:"comment_num"`
}

type PostLike struct {
	Base
	UserId int `json:"user_id`
	PostId int `json:"post_id`
}

type PostComment struct {
	Base
	PostId  int    `json:"post_id"`
	Content string `json:"content"`
}
