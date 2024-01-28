package client

import (
	"encoding/json"
	"io"
)

type CountryInfo struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type Nationalize struct {
	Count   int           `json:"count"`
	Name    string        `json:"name"`
	Country []CountryInfo `json:"country"`
}

func (c *PersonClient) GetNationality(name string) (string, error) {
	resp, err := c.client.Get(nationilizeURL + name)
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
