package gocoinmarketcap

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// type QuotesLatest struct {
// 	Base  string             `json:"base"`
// 	Date  string             `json:"date"`
// 	Rates map[string]float64 `json:"rates"`
// }

func (c *Client) QuotesLatest(params ...map[string]string) {

	u, err := url.Parse(BaseUrl)
	if err != nil {
		log.Println(err)
	}

	u.Path = "/v1/cryptocurrency/quotes/latest"

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
		log.Println(err)
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
}
