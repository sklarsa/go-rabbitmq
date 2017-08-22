package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// Client for RabbitMQ REST API
type Client struct {
	MgmtURL  string
	Username string
	Password string
}

func (c Client) makeRequest(method string, endpoint string) (*http.Response, error) {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s", path.Join(c.MgmtURL, "api", endpoint))
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

func processResponse(response *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	log.Println(string(body))

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf(fmt.Sprintf("Bad Request: %s (Status Code %d). %s", response.Request.URL, response.StatusCode, body))
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return nil

}

// GetOverview GET /api/overview/
func (c Client) GetOverview() (*Overview, error) {
	r, err := c.get("overview")
	if err != nil {
		return nil, err
	}
	o := Overview{}
	err = processResponse(r, &o)
	return &o, err
}

// GetExtensions GET /api/extensions/
func (c Client) GetExtensions() ([]map[string]interface{}, error) {
	r, err := c.get("extensions")
	if err != nil {
		return nil, err
	}
	var m []map[string]interface{}
	err = processResponse(r, &m)
	return m, err
}

// GetConnections GET /api/connections/
func (c Client) GetConnections() ([]Connection, error) {
	r, err := c.get("connections")
	if err != nil {
		return nil, err
	}
	var conn []Connection
	err = processResponse(r, &conn)
	return conn, err

}

// GetConnectionsOnVhost GET /api/vhosts/<vhost>/connections/
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

// GetChannels GET /api/channels/
func (c Client) GetChannels() ([]Channel, error) {
	r, err := c.get("channels")
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = processResponse(r, &channels)
	return channels, err
}

// GetChannelsOnVhost GET /api/vhosts/<vhost>/channels/
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

// GetChannelsOnConnection GET /api/connections/<connection>/channels/
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

// GetClusterName GET /api/cluster-name/
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

// GetNodes GET /api/nodes/
func (c Client) GetNodes() ([]Node, error) {
	r, err := c.get("nodes")
	if err != nil {
		return nil, err
	}
	var n []Node
	err = processResponse(r, &n)
	return n, err
}

// GetNode GET /api/nodes/<node>/
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

// GetExchanges GET /api/exchanges/
func (c Client) GetExchanges() ([]Exchange, error) {
	r, err := c.get("exchanges")
	if err != nil {
		return nil, err
	}
	var e []Exchange
	err = processResponse(r, &e)
	return e, err
}

// GetExchangesOnVhost GET /api/exchanges/<vhost>/
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

// GetBindings GET /api/bindings/
func (c Client) GetBindings() ([]Binding, error) {
	r, err := c.get("bindings")
	if err != nil {
		return nil, err
	}
	var b []Binding
	err = processResponse(r, &b)
	return b, err
}

// GetBindingsOnVhost GET /api/bindings/<vhost>/
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

// GetVhosts GET /api/vhosts/
func (c Client) GetVhosts() ([]Vhost, error) {
	r, err := c.get("vhosts")
	if err != nil {
		return nil, err
	}
	var v []Vhost
	err = processResponse(r, &v)
	return v, err
}

// GetVhostPermissions GET /api/vhosts/<vhost>/permissions/
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

// GetUsers GET /api/users/
func (c Client) GetUsers() ([]User, error) {
	r, err := c.get("users")
	if err != nil {
		return nil, err
	}
	var u []User
	err = processResponse(r, &u)
	return u, err
}

// GetAlivenessTestForVhost GET /api/aliveness-test/<vhost>/
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

// GetHealthcheckForCurrentNode GET /api/healthchecks/node/
func (c Client) GetHealthcheckForCurrentNode() (*Status, error) {
	r, err := c.get("healthchecks/node")
	if err != nil {
		return nil, err
	}
	s := Status{}
	err = processResponse(r, &s)
	return &s, err
}

// GetHealthchecksForNode GET /api/healthchecks/node/<node>/
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

// GetWhoAmI GET /api/whoami/
func (c Client) GetWhoAmI() (*User, error) {
	r, err := c.get("whoami")
	if err != nil {
		return nil, err
	}
	u := User{}
	err = processResponse(r, &u)
	return &u, err
}

// GetQueues GET /api/queues/
func (c Client) GetQueues() ([]Queue, error) {
	r, err := c.get("queues")
	if err != nil {
		return nil, err
	}
	var q []Queue
	err = processResponse(r, &q)
	return q, err
}

// GetQueuesForVhost GET /api/queues/<vhost>/
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

// GetQueue GET /api/<vhost>/<queue>/
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
