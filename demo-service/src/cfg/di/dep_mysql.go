package di

import (
	"github.com/sarulabs/di"

	"github.com/KWRI/demo-service/src/dao/adapter"
)

//
// Definition name.
//
const (
	DefMySQL = "MySQL"
)

//
// initMySQL dependency registrar.
//
func (c *Container) initMySQL() error {

	return c.RegisterDependency(
		DefMySQL,
		func(ctx di.Container) (interface{}, error) {
			return adapter.NewConnection(
				c.GetLogger(),
				c.GetConfig().GetMySQLDSN(),
			)
		},
		nil,
	)
}

//
// GetMySQLDatabase dependency getter.
//
func (c *Container) GetMySQLDatabase() *adapter.Connection {
	return c.Container.Get(DefMySQL).(*adapter.Connection)
}
