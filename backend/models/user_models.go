package models

import "time"

type User struct {
	UserID     int
	Username   string
	Password   string
	CreateTime time.Time
}
