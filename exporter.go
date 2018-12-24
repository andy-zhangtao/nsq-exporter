package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var old_backend_depth, old_in_flight_count, old_deferred_count, old_message_count, old_requeue_count, old_timeout_count int

func recordMetrics() {
	url, host := generateRequestURL()
	logrus.WithFields(logrus.Fields{"NSQD-URL": url}).Debug(ModuleName)

	stats, err := requestStats(url)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Request-NSQD-Stats-Error": err}).Error(ModuleName)
		return
	}

	logrus.WithFields(logrus.Fields{"Stats": stats}).Info(ModuleName)

	prometheus.MustRegister(opsQueuedVec)

	go func() {
		for {

			stats, err = requestStats(url)
			if err != nil {
				logrus.WithFields(logrus.Fields{"Request-NSQD-Stats-Error": err}).Error(ModuleName)
				return
			}

			stats.Host = host
			opsQueuedVec.Reset()

			oldStats := stats

			setupMetrics(stats, oldStats)

			time.Sleep(1 * time.Minute)
		}
	}()
}

func setupMetrics(stats, oldStat nsqdStats) {

	var channels int
	var paused, backend_depth, in_flight_count, deferred_count, message_count, requeue_count, timeout_count int

	for _, s := range stats.Topics {
		channels += len(s.Channels)
		for _, c := range s.Channels {
			backend_depth += c.BackendDepth
			in_flight_count += c.InFlightCount
			deferred_count += c.DeferredCount
			message_count += c.MessageCount
			requeue_count += c.RequeueCount
			timeout_count += c.TimeoutCount
		}

		if s.Paused {
			paused ++
		}
	}

	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Topics"}).Set(float64(len(stats.Topics)))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Channels"}).Set(float64(channels))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Paused"}).Set(float64(paused))

	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_MemoryUsed"}).Set(float64(stats.Memory.HeapInUseBytes))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_MemoryIdle"}).Set(float64(stats.Memory.HeapIdleBytes))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_MemoryGC"}).Set(float64(stats.Memory.GcTotalRuns))

	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "MemoryGC"}).Set(float64(stats.Memory.GcTotalRuns - oldStat.Memory.GcTotalRuns))

	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Backend_Depth"}).Set(float64(backend_depth))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_In_Flight_Count"}).Set(float64(in_flight_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Deferred_Count"}).Set(float64(deferred_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Message_Count"}).Set(float64(message_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Requeue_Count"}).Set(float64(requeue_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Total_Timeout_Count"}).Set(float64(timeout_count))

	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Backend_Depth"}).Set(float64(backend_depth - old_backend_depth))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "In_Flight_Count"}).Set(float64(in_flight_count - old_in_flight_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Deferred_Count"}).Set(float64(deferred_count - old_deferred_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Message_Count"}).Set(float64(message_count - old_message_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Requeue_Count"}).Set(float64(requeue_count - old_requeue_count))
	opsQueuedVec.With(prometheus.Labels{"host": stats.Host, "metrics": "Timeout_Count"}).Set(float64(timeout_count - old_timeout_count))

	old_requeue_count = requeue_count
	old_timeout_count = timeout_count
	old_message_count = message_count
	old_deferred_count = deferred_count
	old_backend_depth = backend_depth
	old_in_flight_count = in_flight_count
}

var opsQueuedVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "nsq",
	Subsystem: "stats",
	Name:      "gauge",
	Help:      "NSQD Stats Info",
}, []string{"host", "metrics"})

func exporter() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":80", nil)
}
