package model

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type ProcessorDefinition[Type any] struct {
	OptionalIdentifiedDefinition[Type]

	self         *Type
	parent       ProcessorNode
	isOutputNode bool
}

// Public Methods

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
	AddOutput(proc, &NewToDefinition(uri).ProcessorDefinition)
	return proc.self
}

func (proc *ProcessorDefinition[Type]) Log(text string) *Type {
	// TODO
	return proc.self
}

// Implements ProcessorNode

func (proc *ProcessorDefinition[Type]) SetParent(parent ProcessorNode) {
	proc.parent = parent
}

func (proc *ProcessorDefinition[Type]) GetParent() ProcessorNode {
	return proc.parent
}

func (proc *ProcessorDefinition[Type]) GetLabel() string {
	return ""
}

func (proc *ProcessorDefinition[Type]) IsOutputNode() bool {
	return proc.isOutputNode
}

func (proc *ProcessorDefinition[Type]) ConfigureChild(ProcessorNode) {
	// To be overridden
}

func (proc *ProcessorDefinition[Type]) AddToOutputList(ProcessorNode) {
	// To be overridden
}
