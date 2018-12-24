package main

type nsqdExporter struct {
	Addr    string
	Topic   string
	Channel string
}

type nsqdStats struct {
	Host      string `json:"host"`
	Version   string `json:"version"`
	Health    string `json:"health"`
	StartTime int    `json:"start_time"`
	Topics    []struct {
		TopicName string `json:"topic_name"`
		Channels  []struct {
			ChannelName   string `json:"channel_name"`
			Depth         int    `json:"depth"`
			BackendDepth  int    `json:"backend_depth"`
			InFlightCount int    `json:"in_flight_count"`
			DeferredCount int    `json:"deferred_count"`
			MessageCount  int    `json:"message_count"`
			RequeueCount  int    `json:"requeue_count"`
			TimeoutCount  int    `json:"timeout_count"`
			Clients       []struct {
				ClientID                      string `json:"client_id"`
				Hostname                      string `json:"hostname"`
				Version                       string `json:"version"`
				RemoteAddress                 string `json:"remote_address"`
				State                         int    `json:"state"`
				ReadyCount                    int    `json:"ready_count"`
				InFlightCount                 int    `json:"in_flight_count"`
				MessageCount                  int    `json:"message_count"`
				FinishCount                   int    `json:"finish_count"`
				RequeueCount                  int    `json:"requeue_count"`
				ConnectTs                     int    `json:"connect_ts"`
				SampleRate                    int    `json:"sample_rate"`
				Deflate                       bool   `json:"deflate"`
				Snappy                        bool   `json:"snappy"`
				UserAgent                     string `json:"user_agent"`
				TLS                           bool   `json:"tls"`
				TLSCipherSuite                string `json:"tls_cipher_suite"`
				TLSVersion                    string `json:"tls_version"`
				TLSNegotiatedProtocol         string `json:"tls_negotiated_protocol"`
				TLSNegotiatedProtocolIsMutual bool   `json:"tls_negotiated_protocol_is_mutual"`
			} `json:"clients"`
			Paused               bool `json:"paused"`
			E2EProcessingLatency struct {
				Count       int         `json:"count"`
				Percentiles interface{} `json:"percentiles"`
			} `json:"e2e_processing_latency"`
		} `json:"channels"`
		Depth                int  `json:"depth"`
		BackendDepth         int  `json:"backend_depth"`
		MessageCount         int  `json:"message_count"`
		Paused               bool `json:"paused"`
		E2EProcessingLatency struct {
			Count       int         `json:"count"`
			Percentiles interface{} `json:"percentiles"`
		} `json:"e2e_processing_latency"`
	} `json:"topics"`
	Memory struct {
		HeapObjects       int `json:"heap_objects"`
		HeapIdleBytes     int `json:"heap_idle_bytes"`
		HeapInUseBytes    int `json:"heap_in_use_bytes"`
		HeapReleasedBytes int `json:"heap_released_bytes"`
		GcPauseUsec100    int `json:"gc_pause_usec_100"`
		GcPauseUsec99     int `json:"gc_pause_usec_99"`
		GcPauseUsec95     int `json:"gc_pause_usec_95"`
		NextGcBytes       int `json:"next_gc_bytes"`
		GcTotalRuns       int `json:"gc_total_runs"`
	} `json:"memory"`
}
