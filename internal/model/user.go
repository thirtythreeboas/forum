package model

type User struct {
	Id       int    `json:"-"`
	Nickname string `json:"nickname"`
	Fullname string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

type NewProfile struct {
	Fullname string
	About    string
	Email    string
	Nickname string
}
