package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

var (
	mgmtUrl  = "localhost:15672/"
	username = "guest"
	password = "guest"
)

type ObjectTotals struct {
	Consumers   int `json:"consumers"`
	Queues      int `json:"queues"`
	Exchanges   int `json:"exchanges"`
	Connections int `json:"connections"`
	Channels    int `json:"channels"`
}

type MessageStats struct {
	Publish      int `"json:publish"`
	PublishIn    int `"json:publish_in"`
	PublishOut   int `"json:publish_out"`
	Confirm      int `"json:confirm"`
	Deliver      int `"json:deliver"`
	DeliverNoAck int `"json:deliver_noack"`
	Get          int `"json:get"`
	GetNoAck     int `"json:get_noack"`
	DeliverGet   int `"json:deliver_get"`
	Redeliver    int `"json:redeliver"`
	Return       int `"json:return"`
}

type Overview struct {
	ClusterName       string       `json:"cluster_name"`
	Contexts          []string     `json:"contexts"`
	ErlangFullVersion []string     `json:"erlang_full_version"`
	ErlangVersion     string       `json:"erlang_version"`
	ExchangeTypes     string       `json:"exchange_types"`
	Listeners         []string     `json:"listeners"`
	ManagementVersion string       `json:"management_version"`
	MessageStats      MessageStats `json:"message_stats"`
	Node              string       `json:"node"`
	ObjectTotals
}

func makeRequest(method string, endpoint string) *http.Response {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s/api/", path.Join(mgmtUrl, endpoint))
	req, _ := http.NewRequest(method, url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return resp
}

func main() {

	resp := makeRequest("api", "GET")
	log.Println(resp)
}
