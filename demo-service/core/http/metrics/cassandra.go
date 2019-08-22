package metrics

import (
	"context"
	"regexp"
	"strings"

	"github.com/gocql/gocql"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//
	// whiteSpaceFiler filter to normalize SQL statements.
	//
	whiteSpaceFiler = regexp.MustCompile(`\s+`)

	//
	// bracesValueFilter filter to remove values inside the `{}` statement.
	//
	bracesValueFilter = regexp.MustCompile(`{(.*?)}`)

	//
	// bracketsValueFilter filter to remove values inside the `()` statement.
	//
	bracketsValueFilter = regexp.MustCompile(`\((.*?)\)`)
)

//
// CassandraObserve Cassandra metric observer struct.
//
type CassandraObserve struct {
	queryObserverMetric      HistogramVec
	connectionObserverMetric Histogram
}

//
// NewCassandraObserver creates new instance of the Cassandra Metric observer.
//
func NewCassandraObserver(metricPrefix string) *CassandraObserve {

	return &CassandraObserve{
		queryObserverMetric: func() HistogramVec {
			sv := prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:      "cassandra_query_latency",
					Namespace: metricPrefix,
					Help:      "Latency of the Cassandra query operation.",
					Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
				},
				[]string{"keyspace", "query"},
			)
			prometheus.MustRegister(sv)

			return sv
		}(),
		connectionObserverMetric: func() Histogram {
			sv := prometheus.NewHistogram(
				prometheus.HistogramOpts{
					Name:      "cassandra_connection_latency",
					Namespace: metricPrefix,
					Help:      "Latency of Cassandra connection operation.",
					Buckets:   []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
				},
			)
			prometheus.MustRegister(sv)

			return sv
		}(),
	}
}

//
// ObserveConnect realizes gocql.ConnectObserver interface.
//
func (o CassandraObserve) ObserveConnect(connect gocql.ObservedConnect) {

	o.connectionObserverMetric.Observe(connect.End.Sub(connect.Start).Seconds())
}

//
// ObserveQuery realizes gocql.QueryObserver interface.
//
func (o CassandraObserve) ObserveQuery(ctx context.Context, query gocql.ObservedQuery) {

	go func() {

		filteredQuery := whiteSpaceFiler.ReplaceAllString(query.Statement, " ")
		filteredQuery = bracesValueFilter.ReplaceAllString(filteredQuery, "{?}")
		filteredQuery = bracketsValueFilter.ReplaceAllString(filteredQuery, "(?)")
		filteredQuery = strings.Trim(filteredQuery, " ")

		o.queryObserverMetric.WithLabelValues(
			query.Keyspace,
			filteredQuery,
		).Observe(query.End.Sub(query.Start).Seconds())

	}()
}
