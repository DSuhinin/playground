package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/KWRI/demo-service/core/errors"
	"github.com/KWRI/demo-service/core/http/response"
	"github.com/KWRI/demo-service/core/http/response/serializer/protobuf"
	"github.com/KWRI/demo-service/core/log"
	"github.com/KWRI/demo-service/core/test/helpers"
)

//
// RouterTest represents the struct for table test.
//
type RouterTest struct {
	name    string
	path    string
	handler func(RouterProvider, string, Handler)
	result  int
}

var RouterTestList = []RouterTest{
	{
		name: "TestRouterWithPostMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Post(s, h)
		},
		result: StatusOK,
	},
	{
		name: "TestRouterWithGetMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Get(s, h)
		},
		result: StatusOK,
	},
	{
		name: "TestRouterWithPutMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Put(s, h)
		},
		result: StatusOK,
	},
	{
		name: "TestRouterWithDeleteMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Delete(s, h)
		},
		result: StatusOK,
	},
	{
		name: "TestRouterWithPatchMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Patch(s, h)
		},
		result: StatusOK,
	},
	{
		name: "TestRouterWithOptionsMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Options(s, h)
		},
		result: StatusOK,
	},
	{
		name: "TestRouterWithTraceMethodAndWithoutTrailingSlash",
		path: "/path",
		handler: func(r RouterProvider, s string, h Handler) {
			r.Trace(s, h)
		},
		result: StatusOK,
	},
}

//
// TestNewRouterCallRouteListWithTrailingSlash :: test the HTTP router, setup each route without trailing
// slash but call each one with trailing slash.
//
func TestNewRouterCallRouteListWithTrailingSlash(t *testing.T) {

	helpers.ResetPrometheusMetrics()

	r := NewRouter(
		log.New(ioutil.Discard, "debug"),
		"metric_prefix",
	)

	for _, test := range RouterTestList {
		t.Run(test.name, func(t *testing.T) {
			test.handler(r, test.path, func(req *http.Request) response.Provider {
				return response.New(nil)
			})

			server := httptest.NewServer(r.GetMuxRouter())
			defer server.Close()

			resp, err := http.Post(server.URL+test.path+"/", "", nil)

			assert.Nil(t, err)
			assert.Equal(t, test.result, resp.StatusCode)
		})
	}
}

//
// TestNewRouterCallRouteListWithoutTrailingSlash :: test the HTTP router, setup each route with trailing slash
// but call each one without trailing slash.
//
func TestNewRouterCallRouteListWithoutTrailingSlash(t *testing.T) {

	helpers.ResetPrometheusMetrics()

	r := NewRouter(
		log.New(ioutil.Discard, "debug"),
		"metric_prefix",
	)

	for _, test := range RouterTestList {
		t.Run(test.name, func(t *testing.T) {
			test.handler(r, test.path+"/", func(req *http.Request) response.Provider {
				return response.New(nil)
			})

			server := httptest.NewServer(r.GetMuxRouter())
			defer server.Close()

			resp, err := http.Post(server.URL+test.path, "", nil)

			assert.Nil(t, err)
			assert.Equal(t, test.result, resp.StatusCode)
		})
	}
}

//
// TestNewRouterWithErrorResponseInProtobufFormat :: test the HTTP router,
// send the error message using Protobuf format.
//
func TestNewRouterWithErrorResponseInProtobufFormat(t *testing.T) {

	helpers.ResetPrometheusMetrics()

	r := NewRouter(
		log.New(ioutil.Discard, "debug"),
		"metric_prefix",
	)

	httpError := errors.NewHTTPError(http.StatusBadRequest, 13, "http error")

	r.Post("/", func(req *http.Request) response.Provider {
		return response.NewProtobuf(httpError)
	})

	server := httptest.NewServer(r.GetMuxRouter())
	defer server.Close()

	resp, err := http.Post(server.URL, "", nil)

	assert.Nil(t, err)
	assert.Equal(t, httpError.GetStatus(), resp.StatusCode)

	responseData, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)

	protobufError := protobuf.HttpError{}
	err = proto.Unmarshal(responseData, &protobufError)

	assert.Nil(t, err)
	assert.Equal(t, httpError.Code, int(protobufError.Code))
	assert.Equal(t, httpError.Msg, protobufError.Message)
}

//
// TestNewRouterWithErrorResponseInJSONFormat :: test the HTTP router,
// send the error message using Protobuf format.
//
func TestNewRouterWithErrorResponseInJSONFormat(t *testing.T) {

	helpers.ResetPrometheusMetrics()

	r := NewRouter(
		log.New(ioutil.Discard, "debug"),
		"metric_prefix",
	)

	httpError := errors.NewHTTPError(http.StatusBadRequest, 13, "http error")

	r.Post("/", func(req *http.Request) response.Provider {
		return response.NewJSON(httpError)
	})

	server := httptest.NewServer(r.GetMuxRouter())
	defer server.Close()

	resp, err := http.Post(server.URL, "", nil)

	assert.Nil(t, err)
	assert.Equal(t, httpError.GetStatus(), resp.StatusCode)

	responseData, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)

	jsonError := errors.HTTPError{}
	err = json.Unmarshal(responseData, &jsonError)

	assert.Nil(t, err)
	assert.Equal(t, httpError.Code, jsonError.GetCode())
	assert.Equal(t, httpError.Msg, jsonError.GetMessage())
}
