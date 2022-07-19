package model

type RouteDefinition struct {
	ProcessorDefinition[RouteDefinition]
	errorHandlerFactory ErrorHandlerFactory
}

// Constructor

func NewRouteDefinition() (answer *RouteDefinition) {
	// TODO
	answer = new(RouteDefinition)
	answer.self = answer
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

// Public Methods

func (route *RouteDefinition) From(uri string) *RouteDefinition {
	route.SetInput(NewFromDefinition(uri))
	return route
}

func (route *RouteDefinition) SetInput(input *FromDefinition) {
	// TODO
}
