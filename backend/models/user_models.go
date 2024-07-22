package models

import "time"

type User struct {
	UserID     int
	Username   string
	Password   string
	CreateTime time.Time
}

type LoginReq struct {
	Username string
	Password string
}

type RegisterReq struct {
	Username string
	Password string
}

type EditReq struct {
	UserID   int
	Username string
	Password string
}

type DeleteReq struct {
	Username string
}
