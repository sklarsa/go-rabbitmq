package api

type ErlangApplication struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

type Rate struct {
	Rate float64 `json:"rate"`
}

type QueueTotals struct {
	Messages               int `"json:messages"`
	MessagesReady          int `"json:messages_ready"`
	MessagesUnacknowledged int `"json:messages_unacknowledged"`
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
