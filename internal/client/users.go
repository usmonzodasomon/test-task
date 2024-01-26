package client

import (
	"encoding/json"
	"io"
	"net/http"
)

type UsersClient struct {
	Client *http.Client
}

func NewUsersClient(client *http.Client) *UsersClient {
	return &UsersClient{
		Client: client,
	}
}

const (
	ageURL         = "https://api.agify.io/?name="
	genderURL      = "https://api.genderize.io/?name="
	nationilizeURL = "https://api.nationalize.io/?name="
)

type Agify struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func (c *UsersClient) GetAge(name string) (int, error) {
	resp, err := c.Client.Get(ageURL + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var UserAge Agify
	if err := json.Unmarshal(body, &UserAge); err != nil {
		return 0, err
	}
	return UserAge.Age, nil
}

type Genderize struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func (c *UsersClient) GetGender(name string) (string, error) {
	resp, err := c.Client.Get(genderURL + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var UserGender Genderize
	if err := json.Unmarshal(body, &UserGender); err != nil {
		return "", err
	}
	return UserGender.Gender, nil
}

type CountryInfo struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type Nationalize struct {
	Count   int           `json:"count"`
	Name    string        `json:"name"`
	Country []CountryInfo `json:"country"`
}

func (c *UsersClient) GetNationality(name string) (string, error) {
	resp, err := c.Client.Get(nationilizeURL + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var UserNationalize Nationalize
	if err := json.Unmarshal(body, &UserNationalize); err != nil {
		return "", err
	}
	if len(UserNationalize.Country) == 0 {
		return "", nil
	}
	return UserNationalize.Country[0].CountryID, nil
}
