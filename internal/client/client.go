package client

import "net/http"

type Users interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetNationality(name string) (string, error)
}

type Client struct {
	Users
}

func NewClient(users Users) *Client {
	return &Client{
		Users: NewUsersClient(&http.Client{}),
	}
}
