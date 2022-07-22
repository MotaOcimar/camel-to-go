package model

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type FromDefinition struct {
	OptionalIdentifiedDefinition[FromDefinition]

	uri      *string
	endpoint cameltogo.Endpoint
}

// Constructor
func NewFromDefinition(uri string) *FromDefinition {
	answer := new(FromDefinition)
	answer.SetUri(uri)
	return answer
}

// Getters and Setters

func (fromDef *FromDefinition) SetUri(uri string) {
	fromDef.clear()
	fromDef.uri = new(string)
	*fromDef.uri = uri
}

// Public Methods

func (fromDef *FromDefinition) GetEndpointUri() *string {
	if fromDef.uri != nil {
		return fromDef.uri
	} else if fromDef.endpoint != nil {
		return fromDef.endpoint.GetEndpointUri()
	} else {
		return nil
	}
}

func (fromDef *FromDefinition) GetLabel() string {
	uri := fromDef.GetEndpointUri()

	if uri == nil {
		return "no uri supplied"
	}
	return *uri
}

func (fromDef *FromDefinition) String() string {
	return "From[" + fromDef.GetLabel() + "]"
}

// Private Methods

func (fromDef *FromDefinition) clear() {
	fromDef.endpoint = nil
	fromDef.uri = nil
}
