package response

import (
	"net/http"

	"github.com/KWRI/demo-service/core/errors"
	"github.com/KWRI/demo-service/core/http/response/serializer"
)

//
// Provider provides a HTTP response.
//
type Provider interface {

	//
	// IsError makes check that current response is a Error response.
	//
	IsError() bool

	//
	// GetError returns current response error.
	//
	GetError() error

	//
	// GetData returns a response data.
	//
	GetData() ([]byte, error)

	//
	// SetStatus sets the response HTTP Status Code.
	//
	SetStatus(int) Provider

	//
	// GetStatus returns a response HTTP Status Code.
	//
	GetStatus() int

	//
	// SetHeader sets response header.
	//
	SetHeader(header string, value string) Provider

	//
	// GetHeader returns a response Header by header name.
	//
	GetHeader(header string) string

	//
	// SetHeaders sets the response Headers.
	//
	SetHeaders(headers http.Header) Provider

	//
	// GetHeaders returns a response Headers list.
	//
	GetHeaders() http.Header

	//
	// SetPostProcessor sets the Response post processor.
	//
	SetPostProcessor(p PostProcesser) Provider

	//
	// GetPostProcessor returns the Response post processor.
	//
	GetPostProcessor() PostProcesser

	//
	// SetSerializer sets the custom response serializer.
	//
	SetSerializer(serializer serializer.Serializer) Provider

	//
	// GetSerializer returns current Response serializer.
	//
	GetSerializer() serializer.Serializer
}

//
// Response represents a HTTP service response.
//
type Response struct {
	data          interface{}
	status        int
	error         error
	headers       http.Header
	serializer    serializer.Serializer
	postProcessor PostProcesser
}

//
// New returns a new HTTP service response instance.
//
func New(data interface{}) Provider {

	responseSerializer := serializer.NewJSON()
	response := Response{
		error:   nil,
		status:  http.StatusOK,
		headers: http.Header{},
	}

	response.SetSerializer(responseSerializer)
	response.SetPostProcessor(&DefaultPostProcessor{})

	if err, ok := data.(error); ok {
		if errorCauser := errors.Cause(err, (*errors.HTTPError)(nil)); errorCauser != nil {
			if httpError, ok := errorCauser.(errors.HTTPError); ok {
				response.data = httpError
				response.error = err
				response.status = httpError.GetStatus()
			}
		}
	} else {
		response.data = data
	}

	return &response
}

//
// NewJSON returns a new HTTP service response instance with JSON data serializer.
//
func NewJSON(data interface{}) Provider {

	return New(data).SetSerializer(serializer.NewJSON())
}

//
// NewProtobuf returns a new HTTP service response instance with Protobuf data serializer.
//
func NewProtobuf(data interface{}) Provider {

	return New(data).SetSerializer(serializer.NewProtobuf())
}

//
// IsError makes check that current response is a Error response.
//
func (r *Response) IsError() bool {

	return r.error != nil
}

//
// GetError returns current response error.
//
func (r *Response) GetError() error {

	return r.error
}

//
// GetData returns a response data.
//
func (r *Response) GetData() ([]byte, error) {

	return r.serializer.SerializeData(r.data)
}

//
// SetStatus sets the response HTTP Status Code.
//
func (r *Response) SetStatus(status int) Provider {

	r.status = status

	return r
}

//
// GetStatus returns a response HTTP Status Code.
//
func (r *Response) GetStatus() int {

	return r.status
}

//
// SetHeader sets response header.
//
func (r *Response) SetHeader(key string, value string) Provider {

	r.headers.Set(key, value)

	return r
}

//
// GetHeader returns a response Header by header name.
//
func (r *Response) GetHeader(name string) string {

	return r.headers.Get(name)
}

//
// SetHeaders sets the response Headers.
//
func (r *Response) SetHeaders(headers http.Header) Provider {

	r.headers = headers

	return r
}

//
// GetHeaders returns a response Headers list.
//
func (r *Response) GetHeaders() http.Header {

	return r.headers
}

//
// SetPostProcessor sets the Response post processor.
//
func (r *Response) SetPostProcessor(pp PostProcesser) Provider {

	pp.SetData(r.data)
	pp.SetResponse(r)
	r.postProcessor = pp

	return r
}

//
// GetPostProcessor returns the Response post processor.
//
func (r *Response) GetPostProcessor() PostProcesser {

	return r.postProcessor
}

//
// SetSerializer sets the custom Response serializer.
//
func (r *Response) SetSerializer(serializer serializer.Serializer) Provider {

	r.serializer = serializer
	r.SetHeader("Content-Type", serializer.GetContentType())

	return r
}

//
// GetSerializer returns current Response serializer.
//
func (r *Response) GetSerializer() serializer.Serializer {

	return r.serializer
}
