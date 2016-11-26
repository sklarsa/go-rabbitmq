package api

type Exchange struct {
	Name       string                 `json:"name"`
	Vhost      string                 `json:"vhost"`
	Type       string                 `json:"type"`
	Durable    bool                   `json:"durable"`
	AutoDelete bool                   `json:"auto_delete"`
	Internal   bool                   `json:"internal"`
	Arguments  map[string]interface{} `json:"arguments"`
}

type Binding struct {
	Source          string                 `json:"source"`
	Vhost           string                 `json:"vhost"`
	Destination     string                 `json:"destination"`
	DestinationType string                 `json:"destination_type"`
	RoutingKey      string                 `json:"routing_key"`
	Arguments       map[string]interface{} `json:"arguments"`
	PropertiesKey   string                 `json:"properties_key"`
}

type Vhost struct {
	SendOct                       int    `json:"send_oct"`
	SendOctDetails                Rate   `json:"send_oct_details"`
	RecvOct                       int    `json:"recv_oct"`
	RecvOctDetails                Rate   `json:"recv_oct_details"`
	Messages                      int    `json:"messages"`
	MessagesDetails               Rate   `json:"messages_details"`
	MessagesReady                 int    `json:"messages_ready"`
	MessagesReadyDetails          Rate   `json:"messages_ready_details"`
	MessagesUnacknowledged        int    `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails Rate   `json:"messages_unacknowledged_details"`
	Name                          string `json:"name"`
	Tracing                       bool   `json:"tracing"`
}

type ErlangApplication struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

type Rate struct {
	Rate float64 `json:"rate"`
}

type QueueTotals struct {
	Messages               int `json:"messages"`
	MessagesReady          int `json:"messages_ready"`
	MessagesUnacknowledged int `json:"messages_unacknowledged"`
}

type ObjectTotals struct {
	Consumers   int `json:"consumers"`
	Queues      int `json:"queues"`
	Exchanges   int `json:"exchanges"`
	Connections int `json:"connections"`
	Channels    int `json:"channels"`
}

type MessageStats struct {
	Publish      int `json:"publish"`
	PublishIn    int `json:"publish_in"`
	PublishOut   int `json:"publish_out"`
	Confirm      int `json:"confirm"`
	Deliver      int `json:"deliver"`
	DeliverNoAck int `json:"deliver_noack"`
	Get          int `json:"get"`
	GetNoAck     int `json:"get_noack"`
	DeliverGet   int `json:"deliver_get"`
	Redeliver    int `json:"redeliver"`
	Return       int `json:"return"`
}

type GenericObject struct {
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
	ClusterName            string          `json:"cluster_name"`
	Contexts               []Context       `json:"contexts"`
	ErlangFullVersion      string          `json:"erlang_full_version"`
	ErlangVersion          string          `json:"erlang_version"`
	ExchangeTypes          []GenericObject `json:"exchange_types"`
	Listeners              []Listener      `json:"listeners"`
	ManagementVersion      string          `json:"management_version"`
	MessageStats           MessageStats    `json:"message_stats"`
	Node                   string          `json:"node"`
	ObjectTotals           ObjectTotals    `json:"object_totals"`
	QueueTotals            QueueTotals     `json:"queue_totals"`
	RabbitMQVersion        string          `json:"rabbitmq_version"`
	RatesMode              string          `json:"rates_modes"`
	StatisticsDbEventQueue int             `json:"statistics_db_event_queue"`
	StatisticsDbNode       string          `json:"statistics_db_node"`
}

type Node struct {
	Applications                          []ErlangApplication `json:"applications"`
	AuthMechanisms                        []GenericObject     `json:"auth_mechanisms"`
	ClusterLinks                          []interface{}       `json:"cluster_links"`
	ConfigFiles                           []string            `json:"config_files"`
	ContextSwitches                       int                 `json:"context_switches"`
	ContextSwitchesDetails                Rate                `json:"context_switches_details"`
	Contexts                              []Context           `json:"contexts"`
	DbDir                                 string              `json:"db_dir"`
	DiskFree                              int                 `json:"disk_free"`
	DiskFreeAlarm                         bool                `json:"disk_free_alarm"`
	DiskFreeDetails                       Rate                `json:"disk_free_details"`
	DiskFreeLimit                         int                 `json:"disk_free_limit"`
	EnabledPlugins                        []string            `json:"enabled_plugins"`
	ExchangeTypes                         []GenericObject     `json:"exchange_types"`
	FdTotal                               int                 `json:"fd_total"`
	FdUsed                                int                 `json:"fd_used"`
	FdUsedDetails                         Rate                `json:"fd_used_details"`
	GcBytesReclaimed                      int                 `json:"gc_bytes_reclaimed"`
	GcBytesReclaimedDetails               Rate                `json:"gc_bytes_reclaimed_details"`
	GcNum                                 int                 `json:"gc_num"`
	GcNumDetails                          Rate                `json:"gc_num_details"`
	IoFileHandleOpenAttemptAvgTime        float64             `json:"io_file_handle_open_attempt_avg_time"`
	IoFileHandleOpenAttemptAvgTimeDetails Rate                `json:"io_file_handle_open_attempt_avg_time_details"`
	IoFileHandleOpenAttemptCount          int                 `json:"io_file_handle_open_attempt_count"`
	IoFileHandleOpenAttemptCountDetails   Rate                `json:"io_file_handle_open_attempt_count_details"`
	IoReadAvgTime                         float64             `json:"io_read_avg_time"`
	IoReadAvgTimeDetails                  Rate                `json:"io_read_avg_time_details"`
	IoReadBytes                           int                 `json:"io_read_bytes"`
	IoReadBytesDetails                    Rate                `json:"io_read_bytes_details"`
	IoReadCount                           int                 `json:"io_read_count"`
	IoReadCountDetails                    Rate                `json:"io_read_count_details"`
	IoReopenCount                         int                 `json:"io_reopen_count"`
	IoReopenCountDetails                  Rate                `json:"io_reopen_count_details"`
	IoSeekAvgTime                         float64             `json:"io_seek_avg_time"`
	IoSeekAvgTimeDetails                  Rate                `json:"io_seek_avg_time_details"`
	IoSeekCount                           int                 `json:"io_seek_count"`
	IoSeekCountDetails                    Rate                `json:"io_seek_count_details"`
	IoSyncAvgTime                         float64             `json:"io_sync_avg_time"`
	IoSyncAvgTimeDetails                  Rate                `json:"io_sync_avg_time_details"`
	IoSyncCount                           int                 `json:"io_sync_count"`
	IoSyncCountDetails                    Rate                `json:"io_sync_count_details"`
	IoWriteAvgTime                        float64             `json:"io_write_avg_time"`
	IoWriteAvgTimeDetails                 Rate                `json:"io_write_avg_time_details"`
	IoWriteBytes                          int                 `json:"io_write_bytes"`
	IoWriteBytesDetails                   Rate                `json:"io_write_bytes_details"`
	IoWriteCount                          int                 `json:"io_write_count"`
	IoWriteCountDetails                   Rate                `json:"io_write_count_details"`
	LogFile                               string              `json:"log_file"`
	MemAlarm                              bool                `json:"mem_alarm"`
	MemLimit                              int64               `json:"mem_limit"`
	MemUsed                               int                 `json:"mem_used"`
	MemUsedDetails                        Rate                `json:"mem_used_details"`
	MnesiaDiskTxCount                     int                 `json:"mnesia_disk_tx_count"`
	MnesiaDiskTxCountDetails              Rate                `json:"mnesia_disk_tx_count_details"`
	MnesiaRAMTxCount                      int                 `json:"mnesia_ram_tx_count"`
	MnesiaRAMTxCountDetails               Rate                `json:"mnesia_ram_tx_count_details"`
	MsgStoreReadCount                     int                 `json:"msg_store_read_count"`
	MsgStoreReadCountDetails              Rate                `json:"msg_store_read_count_details"`
	MsgStoreWriteCount                    int                 `json:"msg_store_write_count"`
	MsgStoreWriteCountDetails             Rate                `json:"msg_store_write_count_details"`
	Name                                  string              `json:"name"`
	NetTicktime                           int                 `json:"net_ticktime"`
	OsPid                                 string              `json:"os_pid"`
	Partitions                            []interface{}       `json:"partitions"`
	ProcTotal                             int                 `json:"proc_total"`
	ProcUsed                              int                 `json:"proc_used"`
	ProcUsedDetails                       Rate                `json:"proc_used_details"`
	Processors                            int                 `json:"processors"`
	QueueIndexJournalWriteCount           int                 `json:"queue_index_journal_write_count"`
	QueueIndexJournalWriteCountDetails    Rate                `json:"queue_index_journal_write_count_details"`
	QueueIndexReadCount                   int                 `json:"queue_index_read_count"`
	QueueIndexReadCountDetails            Rate                `json:"queue_index_read_count_details"`
	QueueIndexWriteCount                  int                 `json:"queue_index_write_count"`
	QueueIndexWriteCountDetails           Rate                `json:"queue_index_write_count_details"`
	RatesMode                             string              `json:"rates_mode"`
	RunQueue                              int                 `json:"run_queue"`
	Running                               bool                `json:"running"`
	SaslLogFile                           string              `json:"sasl_log_file"`
	SocketsTotal                          int                 `json:"sockets_total"`
	SocketsUsed                           int                 `json:"sockets_used"`
	SocketsUsedDetails                    Rate                `json:"sockets_used_details"`
	Type                                  string              `json:"type"`
	Uptime                                int                 `json:"uptime"`
}

type Permission struct {
	Vhost     string `json:"vhost"`
	User      string `json:"user"`
	Configure string `json:"configure"`
	Write     string `json:"write"`
	Read      string `json:"read"`
}

type User struct {
	Name             string `json:"name"`
	PasswordHash     string `json:"password_hash"`
	HashingAlgorithm string `json:"hashing_algorithm"`
	Tags             string `json:"tags"`
}

type Status struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type ClientProperties struct {
	Product        string       `json:"product"`
	ProductVersion string       `json:"product_version"`
	Capabilities   Capabilities `json:"capabilities"`
}

type Capabilities struct {
	ConnectionBlocked          bool `json:"connection.blocked"`
	AuthenticationFailureClose bool `json:"authentication_failure_close"`
	ConsumerCancelNotify       bool `json:"consumer_cancel_notify"`
}
type GarbageCollection struct {
	MaxHeapSize     int `json:"max_heap_size"`
	MinBinVheapSize int `json:"min_bin_vheap_size"`
	MinHeapSize     int `json:"min_heap_size"`
	FullsweepAfter  int `json:"fullsweep_after"`
	MinorGcs        int `json:"minor_gcs"`
}

type Connection struct {
	ConnectedAt       int64             `json:"connected_at"`
	ChannelMax        int               `json:"channel_max"`
	FrameMax          int               `json:"frame_max"`
	Timeout           int               `json:"timeout"`
	Vhost             string            `json:"vhost"`
	User              string            `json:"user"`
	Protocol          string            `json:"protocol"`
	SslHash           interface{}       `json:"ssl_hash"`
	SslCipher         interface{}       `json:"ssl_cipher"`
	SslKeyExchange    interface{}       `json:"ssl_key_exchange"`
	SslProtocol       interface{}       `json:"ssl_protocol"`
	AuthMechanism     string            `json:"auth_mechanism"`
	PeerCertValidity  interface{}       `json:"peer_cert_validity"`
	PeerCertIssuer    interface{}       `json:"peer_cert_issuer"`
	PeerCertSubject   interface{}       `json:"peer_cert_subject"`
	Ssl               bool              `json:"ssl"`
	PeerHost          string            `json:"peer_host"`
	Host              string            `json:"host"`
	PeerPort          int               `json:"peer_port"`
	Port              int               `json:"port"`
	Name              string            `json:"name"`
	Node              string            `json:"node"`
	Type              string            `json:"type"`
	GarbageCollection GarbageCollection `json:"garbage_collection"`
	Reductions        int               `json:"reductions"`
	Channels          int               `json:"channels"`
	State             string            `json:"state"`
	SendPend          int               `json:"send_pend"`
	SendCnt           int               `json:"send_cnt"`
	RecvCnt           int               `json:"recv_cnt"`
	RecvOctDetails    Rate              `json:"recv_oct_details"`
	RecvOct           int               `json:"recv_oct"`
	SendOctDetails    Rate              `json:"send_oct_details"`
	SendOct           int               `json:"send_oct"`
	ReductionsDetails Rate              `json:"reductions_details"`
}

type ConnectionDetails struct {
	Name     string `json:"name"`
	PeerPort int    `json:"peer_port"`
	PeerHost string `json:"peer_host"`
}

type Channel struct {
	Vhost                  string            `json:"vhost"`
	User                   string            `json:"user"`
	Number                 int               `json:"number"`
	Name                   string            `json:"name"`
	Node                   string            `json:"node"`
	GarbageCollection      GarbageCollection `json:"garbage_collection"`
	Reductions             int               `json:"reductions"`
	State                  string            `json:"state"`
	GlobalPrefetchCount    int               `json:"global_prefetch_count"`
	PrefetchCount          int               `json:"prefetch_count"`
	AcksUncommitted        int               `json:"acks_uncommitted"`
	MessagesUncommitted    int               `json:"messages_uncommitted"`
	MessagesUnconfirmed    int               `json:"messages_unconfirmed"`
	MessagesUnacknowledged int               `json:"messages_unacknowledged"`
	ConsumerCount          int               `json:"consumer_count"`
	Confirm                bool              `json:"confirm"`
	Transactional          bool              `json:"transactional"`
	IdleSince              string            `json:"idle_since"`
	ReductionsDetails      Rate              `json:"reductions_details"`
	ConnectionDetails      ConnectionDetails `json:"connection_details"`
}
