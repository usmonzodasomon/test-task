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

type PersonService struct {
	repo   repository.Repository
	client Client
}

func NewPersonService(repo repository.Repository, client Client) *PersonService {
	return &PersonService{
		repo:   repo,
		client: client,
	}
}

func (s *PersonService) AddPerson(person models.AddPersonInput) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, ctx := errgroup.WithContext(ctx)

	var result struct {
		age         int
		gender      string
		nationality string
	}

	r.Go(func() error {
		age, err := s.client.GetAge(person.Name)
		if err != nil {
			return err
		}
		result.age = age
		return nil
	})

	r.Go(func() error {
		gender, err := s.client.GetGender(person.Name)
		if err != nil {
			return err
		}
		result.gender = gender
		return nil
	})

	r.Go(func() error {
		nationality, err := s.client.GetNationality(person.Name)
		if err != nil {
			return err
		}
		result.nationality = nationality
		return nil
	})

	if err := r.Wait(); err != nil {
		return 0, err
	}

	var PersonDB models.Person
	PersonDB.Name = person.Name
	PersonDB.Surname = person.Surname
	PersonDB.Patronomic = person.Patronomic
	PersonDB.Age = result.age
	PersonDB.Gender = result.gender
	PersonDB.Nationality = result.nationality

	return s.repo.AddPerson(PersonDB)
}

func (s *PersonService) DeletePerson(id int64) error {
	return s.repo.DeletePerson(id)
}
