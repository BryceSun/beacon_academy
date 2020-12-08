package model

import "time"

type UserAccount struct {
	Id          int64 `json:"id" form:"id"`
	Ban         bool
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Phone       string `json:"phone" form:"phone"`
	PhoneUseful bool   `json:"phoneUseful" form:"phoneUseful"`
	Email       string `json:"email" form:"email"`
	EmailUseful bool   `json:"emailUseful" form:"emailUseful"`
	HeaderPath  string `json:"headerPath" form:"headerPath"`
	CreateTime  time.Time
	Level       int
}

type UserInfo struct {
	UserId       int64
	Type         int8
	Gender       string
	Name         string
	Birthday     string
	Country      string
	Province     string
	City         string
	Region       string
	detail       string
	SerialNumber string
	WorkIn       string
	Email        string
	Phone        string
}

type UserRole struct {
	Id         int64
	Role       int
	Bind       int
	Module     string
	Desc       string
	ValidTime  time.Time
	CreateTime time.Time
}

type UserOnlineInfo struct {
	Id         int64
	UserId     int64
	LiveTime   int64
	CreateTime time.Time
	Type       int
}
