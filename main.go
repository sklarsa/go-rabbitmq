package main

import (
	"bitbucket.org/sklarsa/go-rabbitmq/api"
	"fmt"
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

}
