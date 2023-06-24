package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "Chatapp"

	GatewaySystem  = "Gateway"
	AuthSystem     = "Authhentication"
	AccountSystem  = "Account"
	ChatroomSystem = "Chatroom"
	IconSystem     = "Icon"
	MessageSystem  = "Message"
	UserSystem     = "User"
)

var TotalConnections = prometheus.NewGauge(prometheus.GaugeOpts{
	Name:      "connections_total",
	Help:      "Total number of connections",
	Namespace: namespace,
})

var PingCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "ping_total",
	Help:      "Total number of ping messages",
	Subsystem: GatewaySystem,
	Namespace: namespace,
})

var PongCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "pong_total",
	Help:      "Total number of pong messages",
	Subsystem: GatewaySystem,
	Namespace: namespace,
})

var PingPongLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "ping_pong_latency",
	Help:      "Latency of ping pong messages",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: GatewaySystem,
	Namespace: namespace,
})

var RequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "requests_total",
	Help:      "Total number of requests",
	Namespace: namespace,
})

var RequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "request_latency",
	Help:      "Latency of requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Namespace: namespace,
})

var AuthRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "auth_requests_total",
	Help:      "Total number of auth requests",
	Subsystem: AuthSystem,
	Namespace: namespace,
})

var AuthRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "auth_request_latency",
	Help:      "Latency of auth requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: AuthSystem,
	Namespace: namespace,
})

var AccountRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "account_requests_total",
	Help:      "Total number of account requests",
	Subsystem: AccountSystem,
	Namespace: namespace,
})

var AccountRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "account_request_latency",
	Help:      "Latency of account requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: AccountSystem,
	Namespace: namespace,
})

var ChatroomRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "chatroom_requests_total",
	Help:      "Total number of chatroom requests",
	Subsystem: ChatroomSystem,
	Namespace: namespace,
})

var ChatroomRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "chatroom_request_latency",
	Help:      "Latency of chatroom requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: ChatroomSystem,
	Namespace: namespace,
})

var IconRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "icon_requests_total",
	Help:      "Total number of icon requests",
	Subsystem: IconSystem,
	Namespace: namespace,
})

var IconRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "icon_request_latency",
	Help:      "Latency of icon requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: IconSystem,
	Namespace: namespace,
})

var MessageRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "message_requests_total",
	Help:      "Total number of message requests",
	Subsystem: MessageSystem,
	Namespace: namespace,
})

var MessageRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "message_request_latency",
	Help:      "Latency of message requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: MessageSystem,
	Namespace: namespace,
})

var UserRequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name:      "user_requests_total",
	Help:      "Total number of user requests",
	Subsystem: UserSystem,
	Namespace: namespace,
})

var UserRequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:      "user_request_latency",
	Help:      "Latency of user requests",
	Buckets:   []float64{0.1, 0.5, 1, 5, 10, 50, 100, 500, 1000},
	Subsystem: UserSystem,
	Namespace: namespace,
})

func init() {
	prometheus.MustRegister(
		// Total stats
		TotalConnections,
		RequestCounter,
		RequestLatency,

		//websocket stats
		PingCounter,
		PongCounter,
		PingPongLatency,

		//services
		AuthRequestCounter,
		AuthRequestLatency,
		AccountRequestCounter,
		AccountRequestLatency,
		ChatroomRequestCounter,
		ChatroomRequestLatency,
		IconRequestCounter,
		IconRequestLatency,
		MessageRequestCounter,
		MessageRequestLatency,
		UserRequestCounter,
		UserRequestLatency,
	)

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":9100", nil)
}
