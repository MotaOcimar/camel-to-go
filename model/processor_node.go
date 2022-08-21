package model

import cameltogo "github.com/MotaOcimar/camel-to-go"

type ProcessorNode interface {
	cameltogo.ContextAware

	GetLabel() string
	SetParent(ProcessorNode)
	GetParent() ProcessorNode
	IsOutputNode() bool
	ConfigureChild(ProcessorNode)
	AddToOutputList(ProcessorNode)
}

func AddOutput(proc ProcessorNode, output ProcessorNode) {
	context := proc.GetContext()
	if context == nil {
		route := GetRoute(proc)
		if route != nil {
			context = route.GetContext()
		}
	}

	// inject context
	cameltogo.TrySetContext(output, context)

	if !proc.IsOutputNode() {
		AddOutput(proc.GetParent(), output)
		return
	}

	output.SetParent(proc)
	proc.ConfigureChild(output)
	proc.AddToOutputList(output)
}

func GetRoute(node ProcessorNode) *RouteDefinition {
	if node == nil {
		return nil
	}

	def := node
	// drill to the top
	for def != nil && def.GetParent() != nil {
		def = def.GetParent()
	}

	answer, ok := def.(*RouteDefinition)
	if ok {
		return answer
	} else {
		// not found
		return nil
	}
}
