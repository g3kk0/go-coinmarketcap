package gocoinmarketcap

import (
	"net/http"
)

const BaseUrl = "https://pro-api.coinmarketcap.com"

type Client struct {
	Conn *http.Client
	Key  string
}

func NewClient(key string) *Client {
	var c Client
	c.Conn = &http.Client{}
	c.Key = key
	return &c
}
