package service

import (
	"context"
	"time"

	"github.com/usmonzodasomon/test-task/internal/models"
	"github.com/usmonzodasomon/test-task/internal/repository"
	"golang.org/x/sync/errgroup"
)

type Client interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetNationality(name string) (string, error)
}

type UsersService struct {
	repo   repository.Repository
	client Client
}

func NewUsersService(repo repository.Repository, client Client) *UsersService {
	return &UsersService{
		repo:   repo,
		client: client,
	}
}

func (s *UsersService) CreateUser(user models.CreateUserInput) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, ctx := errgroup.WithContext(ctx)

	var result struct {
		age         int
		gender      string
		nationality string
	}

	r.Go(func() error {
		age, err := s.client.GetAge(user.Name)
		if err != nil {
			return err
		}
		result.age = age
		return nil
	})

	r.Go(func() error {
		gender, err := s.client.GetGender(user.Name)
		if err != nil {
			return err
		}
		result.gender = gender
		return nil
	})

	r.Go(func() error {
		nationality, err := s.client.GetNationality(user.Name)
		if err != nil {
			return err
		}
		result.nationality = nationality
		return nil
	})

	if err := r.Wait(); err != nil {
		return 0, err
	}

	var UserDB models.User
	UserDB.Name = user.Name
	UserDB.Surname = user.Surname
	UserDB.Patronomic = user.Patronomic
	UserDB.Age = result.age
	UserDB.Gender = result.gender
	UserDB.Nationality = result.nationality

	return s.repo.CreateUser(UserDB)
}

func (s *UsersService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
