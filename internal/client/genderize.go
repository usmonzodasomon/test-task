package client

import (
	"encoding/json"
	"io"
)

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