package client

import (
	"net/http"
)

const (
	ageURL         = "https://api.agify.io/?name="
	genderURL      = "https://api.genderize.io/?name="
	nationilizeURL = "https://api.nationalize.io/?name="
)

type UsersClient struct {
	Client *http.Client
}

func NewUsersClient(client *http.Client) *UsersClient {
	return &UsersClient{
		Client: client,
	}
}
