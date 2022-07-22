package model

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type ProcessorDefinition[Type any] struct {
	OptionalIdentifiedDefinition[Type]

	self *Type
}

func (proc *ProcessorDefinition[Type]) SetBody(expression cameltogo.Expression) *Type {
	// TODO
	return proc.self
}

func (proc *ProcessorDefinition[Type]) SetHeader(name string, expression cameltogo.Expression) *Type {
	// TODO
	return proc.self
}

func (proc *ProcessorDefinition[Type]) RemoveHeader(name string) *Type {
	// TODO
	return proc.self
}

func (proc *ProcessorDefinition[Type]) To(uri string) *Type {
	// TODO
	return proc.self
}

func (proc *ProcessorDefinition[Type]) Log(text string) *Type {
	// TODO
	return proc.self
}
