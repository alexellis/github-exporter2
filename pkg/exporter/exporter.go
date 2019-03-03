package exporter

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var once = sync.Once{}

// RegisterExporter registers with Prometheus for tracking
func RegisterExporter(exporter *Exporter) {
	once.Do(func() {
		prometheus.MustRegister(exporter)
	})
}

type Exporter struct {
	repoStarsCounterVec *prometheus.CounterVec
}

func NewExporter() *Exporter {

	repoStarsCounterVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "repo_stars",
			Help: "GitHub repo stars",
		},
		[]string{"repo", "owner"},
	)

	e := Exporter{
		repoStarsCounterVec: repoStarsCounterVec,
	}

	return &e
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	e.repoStarsCounterVec.Describe(ch)

}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	e.repoStarsCounterVec.WithLabelValues("inlets", "alexellis").Add(1)
	e.repoStarsCounterVec.Collect(ch)

}
