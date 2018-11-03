package gocoinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type QuotesLatest struct {
	Status Status                    `json:"status"`
	Data   map[string]Cryptocurrency `json:"data"`
}

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int64  `json:"elapsed"`
	CreditCount  int64  `json:"credit_count"`
}

type Cryptocurrency struct {
	Id                int64            `json:"id"`
	Name              string           `json:"name"`
	Symbol            string           `json:"symbol"`
	Slug              string           `json:"slug"`
	CirculatingSupply float64          `json:"circulating_supply"`
	TotalSupply       float64          `json:"total_supply"`
	MaxSupply         float64          `json:"max_supply"`
	DateAdded         string           `json:"date_added"`
	NumMarketPairs    int64            `json:"num_market_pairs"`
	CmcRank           int64            `json:"cmc_rank"`
	LastUpdated       string           `json:"last_updated"`
	Quote             map[string]Quote `json:"quote"`
}

type Quote struct {
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume_24h"`
	PercentChange1h  float64 `json:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	PercentChange7d  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}

func (c *Client) QuotesLatest(params ...map[string]string) (QuotesLatest, error) {

	var quotes QuotesLatest
	var err error

	if params[0]["id"] == "" && params[0]["symbol"] == "" {
		return quotes, errors.New("missing 'id' or 'symbol' parameter")
	}

	u, err := url.Parse(BaseUrl)
	if err != nil {
		return quotes, err
	}

	u.Path = "/v1/cryptocurrency/quotes/latest"

	q := u.Query()
	if params[0]["id"] != "" {
		q.Set("id", params[0]["id"])
	} else {
		q.Set("symbol", strings.ToUpper(params[0]["symbol"]))
	}
	if params[0]["convert"] != "" {
		q.Set("convert", strings.ToUpper(params[0]["convert"]))
	}
	u.RawQuery = q.Encode()

	fmt.Printf("u = %+v\n", u)
	fmt.Printf("uStr = %+v\n", u.String())
	fmt.Printf("c.Key = %+v\n", c.Key)

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("X-CMC_PRO_API_KEY", c.Key)
	resp, err := c.Conn.Do(req)
	defer resp.Body.Close()

	// resp, err := c.Conn.Get(u.String())
	// if err != nil {
	// 	return latest, err
	// }
	// defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return quotes, err
	}

	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return quotes, err
	}

	fmt.Printf("body = %+v\n", string(body))

	//var latest Latest

	// u, err := url.Parse(BaseUrl)
	// if err != nil {
	// 	return latest, err
	// }

	// u.Path = "latest"
	// if len(params) > 0 {
	// 	q := u.Query()
	// 	if params[0]["base"] != "" {
	// 		q.Set("base", strings.ToUpper(params[0]["base"]))
	// 	}
	// 	if params[0]["symbols"] != "" {
	// 		q.Set("symbols", strings.ToUpper(params[0]["symbols"]))
	// 	}
	// 	u.RawQuery = q.Encode()
	// }

	// resp, err := c.Conn.Get(u.String())
	// if err != nil {
	// 	return latest, err
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return latest, err
	// }

	// err = json.Unmarshal(body, &latest)
	// if err != nil {
	// 	return latest, err
	// }

	// return latest, nil

	return quotes, nil
}
