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

func (c Client) makeRequest(method string, endpoint string) (*http.Response, error) {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s", path.Join(c.MgmtUrl, "api", endpoint))
	req, _ := http.NewRequest(method, url, nil)
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("Request: %s (Status: %d)\n", url, resp.StatusCode)

	return resp, nil
}

func (c Client) get(endpoint string) (*http.Response, error) {
	return c.makeRequest("GET", endpoint)
}

func (c Client) delete(endpoint string) (*http.Response, error) {
	return c.makeRequest("DELETE", endpoint)
}

func (c Client) put(endpoint string, data interface{}) (*http.Response, error) {
	return c.makeRequest("PUT", endpoint)
}

func (c Client) post(endpoint string, data interface{}) (*http.Response, error) {
	return c.makeRequest("POST", endpoint)
}

func processResponse(response *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Println(string(body))

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Bad Request: %s (Status Code %d). %s", response.Request.URL, response.StatusCode, body))
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return nil

}

func (c Client) GetOverview() (*Overview, error) {
	r, err := c.get("overview")
	if err != nil {
		return nil, err
	}
	o := Overview{}
	err = processResponse(r, &o)
	return &o, err
}

func (c Client) GetExtensions() ([]map[string]interface{}, error) {
	r, err := c.get("extensions")
	if err != nil {
		return nil, err
	}
	var m []map[string]interface{}
	err = processResponse(r, &m)
	return m, err
}

func (c Client) GetConnections() ([]Connection, error) {
	r, err := c.get("connections")
	if err != nil {
		return nil, err
	}
	var conn []Connection
	err = processResponse(r, &conn)
	return conn, err

}

func (c Client) GetConnectionsOnVhost(vhost string) ([]Connection, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("vhosts/%s/connections", vhostStr))
	if err != nil {
		return nil, err
	}
	var conn []Connection
	err = processResponse(r, &conn)
	return conn, err
}

func (c Client) GetChannels() ([]Channel, error) {
	r, err := c.get("channels")
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = processResponse(r, &channels)
	return channels, err
}

func (c Client) GetChannelsOnVhost(vhost string) ([]Channel, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("vhosts/%s/channels", vhostStr))
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = processResponse(r, &channels)
	return channels, err
}

func (c Client) GetChannelsOnConnection(connection string) ([]Channel, error) {
	connStr := url.QueryEscape(connection)
	r, err := c.get(fmt.Sprintf("connections/%s/channels", connStr))
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = processResponse(r, &channels)
	return channels, err
}

func (c Client) GetClusterName() (string, error) {
	r, err := c.get("cluster-name")
	if err != nil {
		return "", err
	}
	type result struct {
		Name string `json:"name"`
	}
	res := result{}
	err = processResponse(r, &res)
	return res.Name, err

}

func (c Client) GetNodes() ([]Node, error) {
	r, err := c.get("nodes")
	if err != nil {
		return nil, err
	}
	var n []Node
	err = processResponse(r, &n)
	return n, err
}

func (c Client) GetNode(node string) (*Node, error) {
	nodeStr := url.QueryEscape(node)
	r, err := c.get(fmt.Sprintf("nodes/%s", nodeStr))
	if err != nil {
		return nil, err
	}
	n := Node{}
	err = processResponse(r, &n)
	return &n, err
}

func (c Client) GetExchanges() ([]Exchange, error) {
	r, err := c.get("exchanges")
	if err != nil {
		return nil, err
	}
	var e []Exchange
	err = processResponse(r, &e)
	return e, err
}

func (c Client) GetExchangesOnVhost(vhost string) ([]Exchange, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("exchanges/%s", vhostStr))
	if err != nil {
		return nil, err
	}
	var e []Exchange
	err = processResponse(r, &e)
	return e, err
}

func (c Client) GetBindings() ([]Binding, error) {
	r, err := c.get("bindings")
	if err != nil {
		return nil, err
	}
	var b []Binding
	err = processResponse(r, &b)
	return b, err
}

func (c Client) GetBindingsOnVhost(vhost string) ([]Binding, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("bindings/%s", vhostStr))
	if err != nil {
		return nil, err
	}
	var b []Binding
	err = processResponse(r, &b)
	return b, err
}

func (c Client) GetVhosts() ([]Vhost, error) {
	r, err := c.get("vhosts")
	if err != nil {
		return nil, err
	}
	var v []Vhost
	err = processResponse(r, &v)
	return v, err
}

func (c Client) GetVhostPermissions(vhost string) ([]Permission, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("vhosts/%s/permissions", vhostStr))
	if err != nil {
		return nil, err
	}
	var p []Permission
	err = processResponse(r, &p)
	return p, err
}

func (c Client) GetUsers() ([]User, error) {
	r, err := c.get("users")
	if err != nil {
		return nil, err
	}
	var u []User
	err = processResponse(r, &u)
	return u, err
}

func (c Client) GetAlivenessTestForVhost(vhost string) (*Status, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("aliveness-test/%s", vhostStr))
	if err != nil {
		return nil, err
	}
	s := Status{}
	err = processResponse(r, &s)
	return &s, err
}

func (c Client) GetHealthcheckForCurrentNode() (*Status, error) {
	r, err := c.get("healthchecks/node")
	if err != nil {
		return nil, err
	}
	s := Status{}
	err = processResponse(r, &s)
	return &s, err
}

func (c Client) GetHealthchecksForNode(node string) (*Status, error) {
	nodeStr := url.QueryEscape(node)
	r, err := c.get(fmt.Sprintf("healthchecks/node/%s", nodeStr))
	if err != nil {
		return nil, err
	}
	s := Status{}
	err = processResponse(r, &s)
	return &s, err
}

func (c Client) GetWhoAmI() (*User, error) {
	r, err := c.get("whoami")
	if err != nil {
		return nil, err
	}
	u := User{}
	err = processResponse(r, &u)
	return &u, err
}

func (c Client) GetQueues() ([]Queue, error) {
	r, err := c.get("queues")
	if err != nil {
		return nil, err
	}
	var q []Queue
	err = processResponse(r, &q)
	return q, err
}

func (c Client) GetQueuesForVhost(vhost string) ([]Queue, error) {
	vhostStr := escapeVhost(vhost)
	r, err := c.get(fmt.Sprintf("queues/%s/", vhostStr))
	if err != nil {
		return nil, err
	}
	var q []Queue
	err = processResponse(r, &q)
	return q, err
}

func (c Client) GetQueue(vhost string, queue string) (*Queue, error) {
	vhostStr := escapeVhost(vhost)
	queueStr := url.QueryEscape(queue)
	r, err := c.get(fmt.Sprintf("queues/%s/%s", vhostStr, queueStr))
	if err != nil {
		return nil, err
	}
	q := Queue{}
	err = processResponse(r, &q)
	return &q, err
}
