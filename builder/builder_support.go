package builder

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type BuilderSupport struct {
	context cameltogo.Context
}

// Setters and Getters
func (builder *BuilderSupport) GetContext() cameltogo.Context {
	return builder.context
}

func (builder *BuilderSupport) SetContext(context cameltogo.Context) {
	builder.context = context
}
