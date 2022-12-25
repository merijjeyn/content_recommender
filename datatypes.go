package main

type Content struct {
	id    int
	Title string
	Url   string
}

type User struct {
	id       int
	NotionId string
	Username string
}

type Edge struct {
	id        int
	UserId    int
	ContentId int
	Weight    float64
}
