package model

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Pages     int    `json:"pages"`
	Country   string `json:"country"`
	ImageLink string `json:"imageLink"`
	Year      string `json:"year"`
	Language  string `json:"language"`
	Link      string `json:"link"`
}
