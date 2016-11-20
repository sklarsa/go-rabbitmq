package api

import (
	"reflect"
	"testing"
)

var (
	mgmtUrl  = "localhost:15672/"
	username = "guest"
	password = "guest"
)

func TestClient(t *testing.T) {

	c := Client{MgmtUrl: mgmtUrl, Username: username, Password: password}
	overview, _ := c.GetOverview()
	if reflect.DeepEqual(overview, Overview{}) {
		t.Errorf("Result from api/overview is nil")
	}

	nodes, _ := c.GetNodes()
	if len(nodes) == 0 {
		t.Errorf("Result from api/nodes has no values")
	}

	node, _ := c.GetNode("rabbit@localhost")
	if reflect.DeepEqual(node, Node{}) {
		t.Errorf("Result from api/node/rabbit@localhost is nil")
	}

	exchanges, _ := c.GetExchanges()
	if len(exchanges) == 0 {
		t.Errorf("Result from api/exchanges has no values")
	}

	exchangesOnVhost, _ := c.GetExchangesOnVhost("/")
	if len(exchangesOnVhost) == 0 {
		t.Errorf("Result from api/exchanges/vhost on default vhost has no values")
	}

}
