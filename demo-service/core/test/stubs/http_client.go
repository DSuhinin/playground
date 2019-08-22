package stubs

import (
	"github.com/stretchr/testify/mock"

	kitHTTP "github.com/KWRI/demo-service/core/http"
)

//
// HTTPClient a stub that implements the http.ClientProvider.
//
type HTTPClient struct {
	mock.Mock
}

//
// Do executes the HTTP request and returns response body.
//
func (s *HTTPClient) Do(req *kitHTTP.ClientRequest) (*kitHTTP.ClientResponse, error) {

	args := s.Mock.Called()
	return args.Get(0).(*kitHTTP.ClientResponse), nil
}
