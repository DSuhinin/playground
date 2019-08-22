package consul

//
// ClientProvider provides interface to work with Consul API.
//
type ClientProvider interface {
	//
	// GetClient returns the consul client.
	//
	GetClient() *APIClient
}
