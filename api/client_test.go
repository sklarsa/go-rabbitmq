package api

import (
	"os"
	"reflect"
	"testing"
)

var (
	mgmtUrl  = "localhost:15672/"
	username = "guest"
	password = "guest"
)

func TestMain(m *testing.M) {
	// Set up a test rabbitmq connection

	//conn, _ := amqp.Dial(os.Getenv("AMQP_URL"))
	//defer conn.Close()
	//
	//	//channel, _ := conn.Channel()
	//	//channel.ExchangeDeclare("test_exchange", "topic", true, false, false, false, amqp.Table{})
	//	//channel.QueueDeclare("test_queue", true, false, false, false, amqp.Table{})
	//channel.QueueBind("test_queue", "#", "test_exchange", false, amqp.Table{})

	os.Exit(m.Run())
}

func TestClient(t *testing.T) {

	c := Client{MgmtUrl: mgmtUrl, Username: username, Password: password}
	overview, _ := c.GetOverview()
	if reflect.DeepEqual(overview, Overview{}) {
		t.Errorf("Result from api/overview is nil")
	}

	extensions, _ := c.GetExtensions()
	if len(extensions) == 0 {
		t.Errorf("Result from api/extensions has no values")
	}

	nodes, _ := c.GetNodes()
	if len(nodes) == 0 {
		t.Errorf("Result from api/nodes has no values")
	}

	node, _ := c.GetNode("rabbit@localhost")
	if reflect.DeepEqual(node, Node{}) {
		t.Errorf("Result from api/node/rabbit@localhost is nil")
	}

	connections, _ := c.GetConnections()
	if len(connections) == 0 {
		t.Errorf("Result from api/connections has no values")
	}

	connections, _ = c.GetConnectionsOnVhost("/")
	if len(connections) == 0 {
		t.Errorf("Result from api/vhosts/vhost/connections has no values")
	}

	channels, _ := c.GetChannels()
	if len(channels) == 0 {
		t.Errorf("Result from api/channels has no values")
	}

	channels, _ = c.GetChannelsOnVhost("/")
	if len(channels) == 0 {
		t.Errorf("Result from api/vhosts/vhost/channels has no values")
	}

	channels, _ = c.GetChannelsOnConnection(connections[0].Name)
	if len(channels) == 0 {
		t.Errorf("Result from api/connections/connection/channels has no values")
	}

	exchanges, _ := c.GetExchanges()
	if len(exchanges) == 0 {
		t.Errorf("Result from api/exchanges has no values")
	}

	exchangesOnVhost, _ := c.GetExchangesOnVhost("/")
	if len(exchangesOnVhost) == 0 {
		t.Errorf("Result from api/exchanges/vhost on default vhost has no values")
	}

	bindings, _ := c.GetBindings()
	if len(bindings) == 0 {
		t.Errorf("Result from api/bindings has no values")
	}

	bindings, _ = c.GetBindingsOnVhost("/")
	if len(bindings) == 0 {
		t.Errorf("Result from api/bindings/vhost has no values")
	}

	vhosts, _ := c.GetVhosts()
	if len(vhosts) == 0 {
		t.Error("Result from api/vhosts has no values")
	}

	permissions, _ := c.GetVhostPermissions("/")
	if len(permissions) == 0 {
		t.Error("Result from api/vhost/name/permissions has no values")
	}

	users, _ := c.GetUsers()
	if len(users) == 0 {
		t.Error("Result from api/users has no values")
	}

	status, _ := c.GetAlivenessTestForVhost("/")
	if reflect.DeepEqual(status, Status{}) {
		t.Error("Result from api/aliveness-test/vhost is nil")
	}

	status, _ = c.GetHealthcheckForCurrentNode()
	if reflect.DeepEqual(status, Status{}) {
		t.Error("Result from api/healthchecks/node is nil")
	}

	status, _ = c.GetHealthchecksForNode("rabbit@localhost")
	if reflect.DeepEqual(status, Status{}) {
		t.Error("Result from api/healthchecks/node/node is nil")
	}
}
