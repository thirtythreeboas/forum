package model

import "time"

type Thread struct {
	Id      int       `json:"-"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Message string    `json:"message"`
	Votes   int       `json:"votes"`
	Slug    string    `json:"slug"`
	Created time.Time `json:"created"`
}

type NewThread struct {
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Slug    string    `json:"slug"`
}
