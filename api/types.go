package api

type QueueTotals struct {
	Messages               int `"messages"`
	MessagesReady          int `"messages_ready"`
	MessagesUnacknowledged int `"messages_unacknowledged"`
}

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

type ExchangeType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type Listener struct {
	Node      string `json:"node"`
	Protocol  string `json:"protocol"`
	IPAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}

type Context struct {
	Node        string `json:"node"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Port        string `json:"port"`
}

type Overview struct {
	ClusterName            string         `json:"cluster_name"`
	Contexts               []Context      `json:"contexts"`
	ErlangFullVersion      string         `json:"erlang_full_version"`
	ErlangVersion          string         `json:"erlang_version"`
	ExchangeTypes          []ExchangeType `json:"exchange_types"`
	Listeners              []Listener     `json:"listeners"`
	ManagementVersion      string         `json:"management_version"`
	MessageStats           MessageStats   `json:"message_stats"`
	Node                   string         `json:"node"`
	ObjectTotals           ObjectTotals   `json:"object_totals"`
	QueueTotals            QueueTotals    `json:"queue_totals"`
	RabbitMQVersion        string         `json:"rabbitmq_version"`
	RatesMode              string         `json:"rates_modes"`
	StatisticsDbEventQueue int            `json:"statistics_db_event_queue"`
	StatisticsDbNode       string         `json:"statistics_db_node"`
}
