package models

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int64  `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}
