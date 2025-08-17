package model

type Forum struct {
	ID      int    `json:"-"`
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int64  `json:"posts"`
	Threads int32  `json:"threads"`
}

type ForumCreate struct {
	Title string `json:"title"`
	User  string `json:"user"`
	Slug  string `json:"slug"`
}
