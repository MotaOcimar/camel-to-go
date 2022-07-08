package model

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type ProcessorDefinition struct {
}

func (proc ProcessorDefinition) SetBody(expression cameltogo.Expression) (answer any) {
	return answer
}

func (proc ProcessorDefinition) SetHeader(name string, expression cameltogo.Expression) (answer any) {
	return answer
}

func (proc ProcessorDefinition) To(uri string) (answer any) {
	return answer
}
