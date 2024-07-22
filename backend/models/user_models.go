package models

import "time"

type User struct {
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"create_time"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EditReq struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteReq struct {
	Username string `json:"username"`
}
