package service

import (
	"context"
	"effective-test/internal/db"
	"effective-test/internal/models"
	"effective-test/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Service interface {
	GetUsers(ctx context.Context, params models.Params) ([]models.User, int64, error)
	CreateUser(ctx context.Context, user *models.User) (int64, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error

	getAge(ctx context.Context, wg *sync.WaitGroup, user *models.User)
	getGender(ctx context.Context, wg *sync.WaitGroup, user *models.User)
	getNationality(ctx context.Context, wg *sync.WaitGroup, user *models.User)
}

type service struct {
	db db.Repository
}

func NewService(db db.Repository) Service {
	return &service{
		db: db,
	}
}

func (s *service) GetUsers(ctx context.Context, params models.Params) ([]models.User, int64, error) {
	users, total, err := s.db.GetUsers(ctx, params)
	if err != nil {
		return []models.User{}, 0, err
	}
	return users, total, nil
}

func (s *service) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	// var err error
	var wg sync.WaitGroup
	wg.Add(3)
	logger.Debug("Enriching...")
	go s.getAge(ctx, &wg, user)
	go s.getGender(ctx, &wg, user)
	go s.getNationality(ctx, &wg, user)
	wg.Wait()
	logger.Debug("The enrichment is complete")

	id, err := s.db.InsertUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) UpdateUser(ctx context.Context, user *models.User) error {
	err := s.db.UpdateUser(ctx, user)
	return err
}

func (s *service) DeleteUser(ctx context.Context, id int64) error {
	err := s.db.DeleteUser(ctx, id)
	return err
}

func (s *service) getAge(ctx context.Context, wg *sync.WaitGroup, user *models.User) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.agify.io/?name=%s", user.Name)
	res, err := http.Get(url)
	if err != nil {
		logger.Errorf("getAge err %v", err)
		return
	}
	var response struct {
		Age int64 `json:"age"`
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Errorf("getAge io.ReadAll err %v", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Errorf("getAge json.Unmarshal err %v", err)
		return
	}

	user.SetAge(response.Age)
}

func (s *service) getGender(ctx context.Context, wg *sync.WaitGroup, user *models.User) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.genderize.io/?name=%s", user.Name)
	res, err := http.Get(url)
	if err != nil {
		logger.Errorf("getAge err %v", err)
		return
	}
	var response struct {
		Gender string `json:"gender"`
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Errorf("getAge io.ReadAll err %v", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Errorf("getAge json.Unmarshal err %v", err)
		return
	}
	user.SetGender(response.Gender)
}

func (s *service) getNationality(ctx context.Context, wg *sync.WaitGroup, user *models.User) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", user.Name)
	res, err := http.Get(url)
	if err != nil {
		logger.Errorf("getAge err %v", err)
		return
	}
	type country struct {
		ID          string  `json:"country_id"`
		Probability float64 `json:"probability"`
	}

	var response struct {
		Country []country `json:"country"`
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Errorf("getAge io.ReadAll err %v", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Errorf("getAge json.Unmarshal err %v", err)
		return
	}
	user.SetNationality(response.Country[0].ID)
}
