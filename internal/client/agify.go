package client

import (
	"encoding/json"
	"io"
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
