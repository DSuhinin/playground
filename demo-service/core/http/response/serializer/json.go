package serializer

import (
	"encoding/json"
)

const (
	//
	// JSONSerializerType the type of current serializer.
	//
	JSONSerializerType = "json"
)

//
// JSON represents the JSON data serializer.
//
type JSON struct {
	serializerType string
}

//
// NewJSONSerializer creates
//
func NewJSON() *JSON {

	return &JSON{
		serializerType: JSONSerializerType,
	}
}

//
// GetType returns the serializer type.
//
func (s JSON) GetType() string {

	return s.serializerType
}

//
// SerializeData makes data serialization.
//
func (s JSON) SerializeData(data interface{}) ([]byte, error) {
	if data == (interface{})(nil) {
		return nil, nil
	}

	return json.Marshal(data)
}

//
// GetContentType returns the Content Type for current serializer.
//
func (s JSON) GetContentType() string {

	return "application/json"
}
