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

	r, _ := c.GetNodes()
	fmt.Println(r)

}
