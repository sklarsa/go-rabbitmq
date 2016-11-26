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
	c := api.Client{mgmtUrl, username, password}
	fmt.Println(c)
}
