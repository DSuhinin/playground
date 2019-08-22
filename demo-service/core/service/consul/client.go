package consul

import (
	"github.com/hashicorp/consul/api"

	"github.com/KWRI/demo-service/core/cfg/config"
	"github.com/KWRI/demo-service/core/errors"
)

//
// APIClient is a type alias for original consul type api.Client
//
type APIClient = api.Client

//
// SessionEntry is a type alias for original consul type api.SessionEntry
//
type SessionEntry = api.SessionEntry

//
// Lock is a type alias for original consul type api.Lock
//
type Lock = api.Lock

//
// LockOptions is a type alias for original consul type api.LockOptions
//
type LockOptions = api.LockOptions

//
// Client represents the struct to hold consul client and
// encapsulate consul external package inside the Virgil Core Kit.
//
type Client struct {
	client *api.Client
}

//
// NewClient creates the new instance of Consul client.
//
func NewClient(connectionInfo config.ConsulConnectionInfoProvider) (*Client, error) {

	config := api.DefaultConfig()
	config.Address = connectionInfo.GetHost()

	if connectionInfo.IsAuthorizationRequired() {
		config.HttpAuth = &api.HttpBasicAuth{
			Username: connectionInfo.GetUser(),
			Password: connectionInfo.GetPassword(),
		}
	}

	if connectionInfo.IsDCAware() {
		config.Datacenter = connectionInfo.GetDataCenter()
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, errors.WithMessage(err, "unable to instantiate consul client")
	}

	return &Client{
		client: client,
	}, nil
}

//
// GetClient returns the consul client for future use.
//
func (c Client) GetClient() *APIClient {

	return c.client
}
