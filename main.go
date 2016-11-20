package main

import (
	"fmt"
	"github.com/sklarsa/go-rabbitmq/api"
)

var (
	mgmtUrl  = "localhost:15672/"
	username = "guest"
	password = "guest"
)

func main() {

	c := api.Client{MgmtUrl: mgmtUrl, Username: username, Password: password}
	overview, _ := c.GetOverview()
	fmt.Println(overview)

	nodes, _ := c.GetNodes()
	fmt.Println(nodes)

	node, _ := c.GetNode("rabbit@localhost")
	fmt.Println(node)

	exchanges, _ := c.GetExchanges()
	fmt.Println(exchanges)

	exchangesOnVhost, _ := c.GetExchangesOnVhost("/")
	fmt.Println(exchangesOnVhost)

}
