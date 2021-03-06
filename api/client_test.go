package api

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"testing"
)

var (
	mgmtURL  = "localhost:15672/"
	username = "guest"
	password = "guest"
)

func TestMain(m *testing.M) {
	os.Exit(RunTests(m))

}

func RunTests(m *testing.M) int {

	log.Println("Creating test exchange and queue...")
	url := os.Getenv("AMQP_URL")
	if url == "" {
		panic("AMQP_URL environment variable is empty")
	}
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	channel.ExchangeDeclare("test_exchange", "topic", true, false, false, false, amqp.Table{})
	channel.QueueDeclare("test_queue", true, false, false, false, amqp.Table{})
	channel.QueueBind("test_queue", "#", "test_exchange", false, amqp.Table{})

	log.Println("Exchange and queue created.  Running unit tests...")

	return m.Run()

}

func TestClient(t *testing.T) {

	c := Client{MgmtURL: mgmtURL, Username: username, Password: password}
	overview, err := c.GetOverview()
	if overview == nil {
		t.Errorf("Result from api/overview is nil")
	}
	if err != nil {
		t.Error(err)
	}

	extensions, err := c.GetExtensions()
	if len(extensions) == 0 {
		t.Errorf("Result from api/extensions has no values")
	}
	if err != nil {
		t.Error(err)
	}

	nodes, err := c.GetNodes()
	if len(nodes) == 0 {
		t.Errorf("Result from api/nodes has no values")
	}
	if err != nil {
		t.Error(err)
	}
	nodeName := nodes[0].Name

	node, err := c.GetNode(nodeName)
	if node == nil {
		t.Errorf("Result from api/node/rabbit@localhost is nil")
	}
	if err != nil {
		t.Error(err)
	}

	connections, err := c.GetConnections()
	if len(connections) == 0 {
		t.Errorf("Result from api/connections has no values")
	}
	if err != nil {
		t.Error(err)
	}

	connections, err = c.GetConnectionsOnVhost("/")
	if len(connections) == 0 {
		t.Errorf("Result from api/vhosts/vhost/connections has no values")
	}
	if err != nil {
		t.Error(err)
	}

	channels, err := c.GetChannels()
	if len(channels) == 0 {
		t.Errorf("Result from api/channels has no values")
	}
	if err != nil {
		t.Error(err)
	}

	channels, err = c.GetChannelsOnVhost("/")
	if len(channels) == 0 {
		t.Errorf("Result from api/vhosts/vhost/channels has no values")
	}
	if err != nil {
		t.Error(err)
	}

	channels, err = c.GetChannelsOnConnection(connections[0].Name)
	if len(channels) == 0 {
		t.Errorf("Result from api/connections/connection/channels has no values")
	}
	if err != nil {
		t.Error(err)
	}

	exchanges, err := c.GetExchanges()
	if len(exchanges) == 0 {
		t.Errorf("Result from api/exchanges has no values")
	}
	if err != nil {
		t.Error(err)
	}

	exchangesOnVhost, err := c.GetExchangesOnVhost("/")
	if len(exchangesOnVhost) == 0 {
		t.Errorf("Result from api/exchanges/vhost on default vhost has no values")
	}
	if err != nil {
		t.Error(err)
	}

	bindings, err := c.GetBindings()
	if len(bindings) == 0 {
		t.Errorf("Result from api/bindings has no values")
	}
	if err != nil {
		t.Error(err)
	}

	bindings, err = c.GetBindingsOnVhost("/")
	if len(bindings) == 0 {
		t.Errorf("Result from api/bindings/vhost has no values")
	}
	if err != nil {
		t.Error(err)
	}

	vhosts, err := c.GetVhosts()
	if len(vhosts) == 0 {
		t.Error("Result from api/vhosts has no values")
	}
	if err != nil {
		t.Error(err)
	}

	permissions, err := c.GetVhostPermissions("/")
	if len(permissions) == 0 {
		t.Error("Result from api/vhost/name/permissions has no values")
	}
	if err != nil {
		t.Error(err)
	}

	users, err := c.GetUsers()
	if len(users) == 0 {
		t.Error("Result from api/users has no values")
	}
	if err != nil {
		t.Error(err)
	}

	status, err := c.GetAlivenessTestForVhost("/")
	if status == nil {
		t.Error("Result from api/aliveness-test/vhost is nil")
	}
	if err != nil {
		t.Error(err)
	}

	status, err = c.GetHealthcheckForCurrentNode()
	if status == nil {
		t.Error("Result from api/healthchecks/node is nil")
	}
	if err != nil {
		t.Error(err)
	}

	status, err = c.GetHealthchecksForNode(nodeName)
	if status == nil {
		t.Error("Result from api/healthchecks/node/node is nil")
	}
	if err != nil {
		t.Error(err)
	}

	user, err := c.GetWhoAmI()
	if user == nil {
		t.Error("Results from api/whoami is nil")
	}
	if err != nil {
		t.Error(err)
	}

	queues, err := c.GetQueues()
	if len(queues) == 0 {
		t.Error("Result from api/queues has no values")
	}
	if err != nil {
		t.Error(err)
	}

	queues, err = c.GetQueuesForVhost("/")
	if len(queues) == 0 {
		t.Error("Result from api/queues/vhost has no values")
	}
	if err != nil {
		t.Error(err)
	}

	queue, err := c.GetQueue("/", "test_queue")
	if queue == nil {
		t.Error("Result from api/queues/vhost/name is nil")
	}
	if err != nil {
		t.Error(err)
	}
}
