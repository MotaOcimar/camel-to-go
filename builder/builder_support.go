package builder

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
)

type BuilderSupport struct {
	context *cameltogo.Context
}

// Setters and Getters
func (builder *BuilderSupport) GetContext() *cameltogo.Context {
	return builder.context
}

func (builder *RouteBuilder) SetContext(context *cameltogo.Context) {
	if context != nil {
		builder.context = context
	}
}
