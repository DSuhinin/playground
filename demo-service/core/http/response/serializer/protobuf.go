package serializer

import (
	"github.com/golang/protobuf/proto"

	"github.com/KWRI/demo-service/core/errors"
	"github.com/KWRI/demo-service/core/http/response/serializer/protobuf"
)

const (
	//
	// ProtobufSerializerType the type of current serializer.
	//
	ProtobufSerializerType = "protobuf"
)

//
// Protobuf represents the Protobuf data serializer.
//
type Protobuf struct {
	serializerType string
}

//
// NewProtobuf creates
//
func NewProtobuf() *Protobuf {

	return &Protobuf{
		serializerType: ProtobufSerializerType,
	}
}

//
// GetType returns the serializer type.
//
func (s Protobuf) GetType() string {

	return s.serializerType
}

//
// SerializeData makes data serialization.
//
func (s Protobuf) SerializeData(data interface{}) ([]byte, error) {

	if data == (interface{})(nil) {
		return nil, nil
	}

	if error, ok := data.(errors.HTTPError); ok {
		data = &protobuf.HttpError{
			Code:    uint32(error.GetCode()),
			Message: error.GetMessage(),
		}
	}

	if protoData, ok := data.(proto.Message); ok {
		return proto.Marshal(protoData)
	}

	return nil, nil
}

//
// GetContentType returns the Content Type for current serializer.
//
func (s Protobuf) GetContentType() string {

	return "application/protobuf"
}
