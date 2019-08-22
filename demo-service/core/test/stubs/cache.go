package stubs

import (
	"net/http"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/KWRI/demo-service/core/http/health"
)

//
// Cache a stub that implements cache.ClientProvider interface.
//
type Cache struct {
	mock.Mock
}

//
// SetKey sets a value for key for duration.
//
func (s *Cache) SetKey(key string, val interface{}, expiration time.Duration) error {

	return nil
}

//
// GetKey returns key value.
//
func (s *Cache) GetKey(key string) ([]byte, error) {

	args := s.Mock.Called()
	if err := args.Get(1); err != nil {
		return args.Get(0).([]byte), err.(error)
	}

	return args.Get(0).([]byte), nil
}

//
// GetHealth returns the Cache Health info.
//
func (s *Cache) GetHealth() (*health.Data, error) {

	return &health.Data{
		Name:   "cache_stub",
		Status: http.StatusOK,
	}, nil
}
