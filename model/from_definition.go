package model

type FromDefinition struct {
	OptionalIdentifiedDefinition[FromDefinition]
	uri string
}

// Constructor
func NewFromDefinition(uri string) *FromDefinition {
	// TODO
	return new(FromDefinition)
}

// Getters and Setters
func (fromDef *FromDefinition) SetUri(uri string) {
	fromDef.clear()
	fromDef.uri = uri
}

// Private Methods
func (fromDef *FromDefinition) clear() {
	// TODO
	fromDef.uri = ""
}
