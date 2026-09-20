package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	thingsCreated prometheus.Counter
}

func NewPrometheus() *Prometheus {
	thingsCreated := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "things_created_total",
			Help: "Total number of created things",
		},
	)

	prometheus.MustRegister(thingsCreated)

	return &Prometheus{
		thingsCreated: thingsCreated,
	}
}

func (m *Prometheus) ThingCreated() {
	m.thingsCreated.Inc()
}
