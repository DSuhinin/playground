package di

import (
	"testing"

	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
)

//
// TestNewContainer :: ensure that container is created and has a di.builder.
//
func TestNewContainer(t *testing.T) {
	container, err := NewContainer()

	assert.NoError(t, err)
	assert.NotEmpty(t, container)
	assert.NotEmpty(t, container.builder)
}

//
// TestNewContainer :: ensure that Build method creates context.
//
func TestContainer_Build(t *testing.T) {
	container, err := NewContainer()

	assert.NoError(t, err)
	assert.Nil(t, container.ctx)

	container.Build()

	assert.NotEmpty(t, container.ctx)
}

//
// TestContainer_RegisterDependency :: check that RegisterDependency doesn't cause error.
//
func TestContainer_RegisterDependency(t *testing.T) {
	container, err := NewContainer()

	assert.NoError(t, err)
	assert.NoError(t, container.RegisterDependency(
		"TestDependency",
		func(ctx di.Container) (interface{}, error) {
			return nil, nil
		},
		nil,
	))
}
