package model

import "time"

type UserAccount struct {
	Id          int64 `json:"id" req:"id"`
	Ban         bool
	Username    string `json:"username" req:"username"`
	Password    string `json:"password" req:"password"`
	Phone       string `json:"phone" req:"phone"`
	PhoneUseful string `json:"phoneUseful" req:"phoneUseful"`
	Email       string `json:"email" req:"email"`
	EmailUseful bool   `json:"emailUseful" req:"emailUseful"`
	HeaderPath  string `json:"headerPath" req:"headerPath"`
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
	Id int64
	UserId int64
	LiveTime int64
	CreateTime time.Time
	Type int
}