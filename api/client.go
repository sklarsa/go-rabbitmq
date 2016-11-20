package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type Client struct {
	MgmtUrl  string
	Username string
	Password string
}

func (c Client) makeRequest(method string, endpoint string) *http.Response {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s", path.Join(c.MgmtUrl, "api", endpoint))
	req, _ := http.NewRequest(method, url, nil)
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err.Error)
	}

	log.Printf("Request: %s (Status: %d)\n", url, resp.StatusCode)

	return resp
}

func processResponse(response *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil

}

func (c Client) GetOverview() Overview {
	r := c.makeRequest("GET", "overview")
	o := Overview{}
	processResponse(r, &o)
	return o
}
