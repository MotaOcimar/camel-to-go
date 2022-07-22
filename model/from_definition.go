package model

type FromDefinition struct {
	OptionalIdentifiedDefinition[FromDefinition]

	uri *string
}

// Constructor
func NewFromDefinition(uri string) *FromDefinition {
	answer := new(FromDefinition)
	answer.SetUri(uri)
	return answer
}

// Getters and Setters

func (fromDef *FromDefinition) SetUri(uri string) {
	fromDef.uri = new(string)
	*fromDef.uri = uri
}

// Public Methods

func (fromDef *FromDefinition) GetLabel() string {

	if fromDef.uri == nil {
		return "no uri supplied"
	}
	return *fromDef.uri
}

func (fromDef *FromDefinition) String() string {
	return "From[" + fromDef.GetLabel() + "]"
}
