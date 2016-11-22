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

func (c Client) get(endpoint string) *http.Response {
	return c.makeRequest("GET", endpoint)
}

func (c Client) delete(endpoint string) *http.Response {
	return c.makeRequest("DELETE", endpoint)
}

func (c Client) put(endpoint string, data interface{}) *http.Response {
	return c.makeRequest("PUT", endpoint)
}

func (c Client) post(endpoint string, data interface{}) *http.Response {
	return c.makeRequest("POST", endpoint)
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
	r := c.get("overview")
	o := Overview{}
	err := processResponse(r, &o)
	return o, err
}

func (c Client) GetClusterName() (string, error) {
	r := c.get("cluster-name")
	type result struct {
		Name string `json:"name"`
	}
	res := result{}
	err := processResponse(r, &res)
	return res.Name, err

}

func (c Client) GetNodes() ([]Node, error) {
	r := c.get("nodes")
	var n []Node
	err := processResponse(r, &n)
	return n, err
}

func (c Client) GetNode(node string) (Node, error) {
	nodeStr := url.QueryEscape(node)
	r := c.get(fmt.Sprintf("nodes/%s", nodeStr))
	n := Node{}
	err := processResponse(r, &n)
	return n, err
}

func (c Client) GetExchanges() ([]Exchange, error) {
	r := c.get("exchanges")
	var e []Exchange
	err := processResponse(r, &e)
	return e, err
}

func (c Client) GetExchangesOnVhost(vhost string) ([]Exchange, error) {
	vhostStr := escapeVhost(vhost)
	r := c.get(fmt.Sprintf("exchanges/%s", vhostStr))
	var e []Exchange
	err := processResponse(r, &e)
	return e, err
}

func (c Client) GetBindings() ([]Binding, error) {
	r := c.get("bindings")
	var b []Binding
	err := processResponse(r, &b)
	return b, err
}

func (c Client) GetBindingsOnVhost(vhost string) ([]Binding, error) {
	vhostStr := escapeVhost(vhost)
	r := c.get(fmt.Sprintf("bindings/%s", vhostStr))
	var b []Binding
	err := processResponse(r, &b)
	return b, err
}

func (c Client) GetVhosts() ([]Vhost, error) {
	r := c.get("vhosts")
	var v []Vhost
	err := processResponse(r, &v)
	return v, err
}

func (c Client) GetVhostPermissions(vhost string) ([]Permission, error) {
	vhostStr := escapeVhost(vhost)
	r := c.get(fmt.Sprintf("vhosts/%s/permissions", vhostStr))
	var p []Permission
	err := processResponse(r, &p)
	return p, err
}

func (c Client) GetUsers() ([]User, error) {
	r := c.get("users")
	var u []User
	err := processResponse(r, &u)
	return u, err
}

func (c Client) GetAlivenessTestForVhost(vhost string) (Status, error) {
	vhostStr := escapeVhost(vhost)
	r := c.get(fmt.Sprintf("aliveness-test/%s", vhostStr))
	s := Status{}
	err := processResponse(r, &s)
	return s, err
}

func (c Client) GetHealthcheckForCurrentNode() (Status, error) {
	r := c.get("healthchecks/node")
	s := Status{}
	err := processResponse(r, &s)
	return s, err
}

func (c Client) GetHealthchecksForNode(node string) (Status, error) {
	nodeStr := url.QueryEscape(node)
	r := c.get(fmt.Sprintf("healthchecks/node/%s", nodeStr))
	s := Status{}
	err := processResponse(r, &s)
	return s, err
}
