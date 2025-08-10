package model

type Forum struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int64  `json:"posts,omitempty"`
	Threads int32  `json:"threads,omitempty"`
}

type ForumCreate struct {
	Title string `json:"title"`
	User  string `json:"user"`
	Slug  string `json:"slug"`
}
