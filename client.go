package gocoinmarketcap

import (
	"net/http"
	"os"
)

const BaseUrl = "https://pro-api.coinmarketcap.com"

type Client struct {
	Conn *http.Client
	Key  string
}

func NewClient() *Client {
	var c Client
	c.Conn = &http.Client{}
	c.Key = os.Getenv("COINMARKETCAP_KEY")
	return &c
}
