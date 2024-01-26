package models

type User struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronomic  string `json:"patronomic"`
	Age         int    `json:"age" `
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}