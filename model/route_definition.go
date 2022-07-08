package model

type RouteDefinition struct {
	uri string
	ProcessorDefinition
}

func (route RouteDefinition) Log(text string) RouteDefinition {
	// TODO
	return route
}
