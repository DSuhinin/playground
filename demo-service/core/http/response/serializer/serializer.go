package serializer

//
// Serializer provides an interface to work with the data serialization
//
type Serializer interface {
	//
	// GetType returns the serializer type.
	//
	GetType() string

	//
	// SerializeData makes data serialization.
	//
	SerializeData(data interface{}) ([]byte, error)

	//
	// GetContentType returns the Content Type for current serializer.
	//
	GetContentType() string
}
