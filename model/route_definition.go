package model

type RouteDefinition struct {
	ProcessorDefinition[RouteDefinition]

	errorHandlerFactory ErrorHandlerFactory
	input               *FromDefinition
}

// Constructor
func NewRouteDefinition() (answer *RouteDefinition) {
	answer = new(RouteDefinition)
	answer.self = answer
	answer.isOutputNode = true
	return answer
}

// Getters and Setters

func (route *RouteDefinition) SetErrorHandlerFactory(errorHandlerFactory ErrorHandlerFactory) {
	route.errorHandlerFactory = errorHandlerFactory
}

func (route *RouteDefinition) SetErrorHandlerFactoryIfNil(errorHandlerFactory ErrorHandlerFactory) {
	if route.errorHandlerFactory == nil {
		route.SetErrorHandlerFactory(errorHandlerFactory)
	}
}

func (route *RouteDefinition) SetInput(input *FromDefinition) {
	if route.input != nil && input != nil && route.input != input {
		panic("Only one input is allowed per route. Cannot accept input: " + input.String())
	}

	route.input = input

}

// Public Methods

func (route *RouteDefinition) From(uri string) *RouteDefinition {
	route.SetInput(NewFromDefinition(uri))
	return route
}
