package client

import (
	"net/http"
)

const (
	ageURL         = "https://api.agify.io/?name="
	genderURL      = "https://api.genderize.io/?name="
	nationilizeURL = "https://api.nationalize.io/?name="
)

type Person interface {
	GetAge(name string) (int, error)
	GetGender(name string) (string, error)
	GetNationality(name string) (string, error)
}

type Client struct {
	Person
}

func NewUsersClient(client *http.Client) *Client {
	return &Client{
		Person: NewPersonClient(client),
	}
}
