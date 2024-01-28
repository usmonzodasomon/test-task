package models

type Person struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronomic  string `json:"patronomic"`
	Age         int    `json:"age" `
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type AddPersonInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type GetPersonRequest struct {
	Limit  int
	Offset int

	Age         int
	Gender      string
	Nationality string
}
