package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
		log.Panic(err)
	}

	log.Printf("Request: %s (Status: %d)\n", url, resp.StatusCode)

	return resp
}

func processResponse(response *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Bad Request: %s (Status Code %d). %s", response.Request.URL, response.StatusCode, body))
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return nil

}

func (c Client) GetOverview() (Overview, error) {
	r := c.makeRequest("GET", "overview")
	o := Overview{}
	err := processResponse(r, &o)
	return o, err
}

func (c Client) GetClusterName() (string, error) {
	r := c.makeRequest("GET", "cluster-name")
	type result struct {
		Name string `json:"name"`
	}
	res := result{}
	err := processResponse(r, &res)
	return res.Name, err

}

func (c Client) GetNodes() ([]Node, error) {
	r := c.makeRequest("GET", "nodes")
	var n []Node
	err := processResponse(r, &n)
	return n, err
}

func (c Client) GetNode(node string) (Node, error) {
	nodeStr := url.QueryEscape(node)
	r := c.makeRequest("GET", fmt.Sprintf("nodes/%s", nodeStr))
	n := Node{}
	err := processResponse(r, &n)
	return n, err
}

func (c Client) GetExchanges() ([]Exchange, error) {
	r := c.makeRequest("GET", "exchanges")
	var e []Exchange
	err := processResponse(r, &e)
	return e, err
}

func (c Client) GetExchangesOnVhost(vhost string) ([]Exchange, error) {
	vhostStr := escapeVhost(vhost)
	r := c.makeRequest("GET", fmt.Sprintf("exchanges/%s", vhostStr))
	var e []Exchange
	err := processResponse(r, &e)
	return e, err
}
