package cache

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-redis/redis"

	"github.com/KWRI/demo-service/core/errors"
	"github.com/KWRI/demo-service/core/http/health"
	"github.com/KWRI/demo-service/core/http/metrics"
	"github.com/KWRI/demo-service/core/log"
)

//
// Nil Redis nil reply, .e.g. when key does not exist.
//
var Nil = redis.Nil

const (
	//
	// ServiceName name returning in health info.
	//
	ServiceName = "redis_cache"
)

//
// Redis client instance.
//
type Redis struct {
	// TODO:: I want to get rid of this hardcoded dependency here!!!
	client             *redis.Client
	logger             log.Logger
	writeLatencyMetric metrics.Summary
	readLatencyMetric  metrics.Summary
	writeErrorMetric   metrics.Counter
	readErrorMetric    metrics.Counter
}

//
// NewRedisClient returns a new Redis client instance.
//
func NewRedisClient(logger log.Logger, client *redis.Client, metricPrefix string) *Redis {

	return &Redis{
		client:             client,
		logger:             logger,
		writeLatencyMetric: metrics.NewRedisWriteLatency(metricPrefix),
		writeErrorMetric:   metrics.NewRedisWriteError(metricPrefix),
		readLatencyMetric:  metrics.NewRedisReadLatency(metricPrefix),
		readErrorMetric:    metrics.NewRedisReadError(metricPrefix),
	}
}

//
// SetKey sets a value for key for duration.
//
func (r *Redis) SetKey(key string, val interface{}, expiration time.Duration) error {

	var value string
	switch v := val.(type) {
	case string:
		value = v
	case []byte:
		value = string(v)
	default:
		vv, err := json.Marshal(v)
		if nil != err {
			return errors.WithMessage(
				err, "kit.cache@Redis.SetKey [value (%s) marshalling error for key (%s)]", val, key,
			)
		}
		value = string(vv)
	}

	begin := time.Now()
	defer func() {
		r.writeLatencyMetric.Observe(time.Since(begin).Seconds())
	}()

	if err := r.client.Set(key, value, expiration).Err(); nil != err {
		r.writeErrorMetric.Inc()
		return errors.WithMessage(
			err, "kit.cache@Redis.SetKey [key (%s) setting error for val (%s)]", key, val,
		)
	}

	return nil
}

//
// GetKey returns key value.
//
func (r *Redis) GetKey(key string) ([]byte, error) {

	begin := time.Now()
	defer func() {
		r.readLatencyMetric.Observe(time.Since(begin).Seconds())
	}()

	result, err := r.client.Get(key).Result()
	if err != nil {
		if err != Nil {
			r.readErrorMetric.Inc()
			return nil, errors.WithMessage(err, "kit.cache@Redis.SetKey [error reading key (%s) value]", key)
		}
	}

	return []byte(result), nil
}

//
// GetHealth returns the Redis Cache Health info.
//
func (r *Redis) GetHealth() (*health.Data, error) {

	h := health.Data{
		Name:   ServiceName,
		Status: http.StatusOK,
	}

	startTime := time.Now()
	defer func() {
		h.Latency = time.Now().Sub(startTime).Seconds()
	}()

	if _, err := r.client.Ping().Result(); err != nil {
		h.Status = http.StatusBadRequest
		return &h, errors.WithMessage(err, "redis cache health check error")

	}

	return &h, nil
}

//
// Shutdown closes connection with Redis service. Method needs for graceful shutdown.
//
func (r *Redis) Shutdown() {

	if err := r.client.Close(); err != nil {
		r.logger.Error("redis connection has been closed with an error: %+v", err)
		return
	}

	r.logger.Info("redis connection has been closed")
}
