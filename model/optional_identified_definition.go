package model

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type OptionalIdentifiedDefinition[Type any] struct {
	context cameltogo.Context
}

func (optIdDef *OptionalIdentifiedDefinition[Type]) SetContext(context cameltogo.Context) {
	optIdDef.context = context
}

func (optIdDef *OptionalIdentifiedDefinition[Type]) GetContext() cameltogo.Context {
	return optIdDef.context
}
