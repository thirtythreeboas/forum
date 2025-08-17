package model

type NewPost struct {
	Parent  int    `josn:"parent"`
	Author  string `josn:"author"`
	Message string `josn:"message"`
}
