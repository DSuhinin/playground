package di

import "github.com/KWRI/demo-service/src/cfg/config"

//
// GetConfig dependency retriever.
//
func (c *Container) GetConfig() *config.Config {

	return c.config
}
