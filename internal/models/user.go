package models

import "sync"

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int64  `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`

	mu sync.RWMutex
}

func (u *User) SetAge(age int64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Age = age
}

func (u *User) SetGender(gender string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Gender = gender
}

func (u *User) SetNationality(nationality string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Nationality = nationality
}
