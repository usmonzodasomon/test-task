package client

import "net/http"

type PersonClient struct {
	client *http.Client
}

func NewPersonClient(client *http.Client) *PersonClient {
	return &PersonClient{
		client: client,
	}
}
