package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	DataName    = "SomeName"
	DataStatus  = 1
	DataLatency = 0.1
)

//
// TestData_GetName :: check name getter
//
func TestData_GetName(t *testing.T) {
	data := Data{Name: DataName, Status: DataStatus, Latency: DataLatency}

	assert.Equal(t, DataName, data.GetName())
}

//
// TestData_GetStatus :: check status getter
//
func TestData_GetStatus(t *testing.T) {
	data := Data{Name: DataName, Status: DataStatus, Latency: DataLatency}

	assert.Equal(t, DataStatus, data.GetStatus())
}

//
// TestData_GetLatency :: check latency getter
//
func TestData_GetLatency(t *testing.T) {
	data := Data{Name: DataName, Status: DataStatus, Latency: DataLatency}

	assert.Equal(t, DataLatency, data.GetLatency())
}
