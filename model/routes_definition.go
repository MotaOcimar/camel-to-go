package model

type RoutesDefinition struct {
	OptionalIdentifiedDefinition[RouteDefinition]

	routes              []*RouteDefinition
	errorHandlerFactory ErrorHandlerFactory
}

// Setters and Getters

func (routes *RoutesDefinition) GetErrorHandlerFactory() (handler ErrorHandlerFactory) {
	return routes.errorHandlerFactory
}

func (routes *RoutesDefinition) GetRoutes() []*RouteDefinition {
	return routes.routes
}

func (rd *RoutesDefinition) SetRoutes(routes []*RouteDefinition) {
	rd.routes = routes
}

func (rd *RoutesDefinition) AddRoute(route *RouteDefinition) {
	rd.SetRoutes(append(rd.GetRoutes(), route))
}

// Public Methods

func (routes *RoutesDefinition) From(uri string) *RouteDefinition {
	route := routes.createRoute()
	route.From(uri)
	return routes.Route(route)
}

func (routes *RoutesDefinition) Route(route *RouteDefinition) *RouteDefinition {
	// must set the error handler if not already set on the route
	handler := routes.GetErrorHandlerFactory()
	if handler != nil {
		route.SetErrorHandlerFactoryIfNil(handler)
	}
	routes.AddRoute(route)
	return route
}

func (routes *RoutesDefinition) PrepareRoute(route *RouteDefinition) {
	// TODO
}

// Private Methods

func (routes *RoutesDefinition) createRoute() (route *RouteDefinition) {
	route = NewRouteDefinition()
	route.SetContext(routes.GetContext())
	handler := routes.GetErrorHandlerFactory()
	if handler != nil {
		route.SetErrorHandlerFactoryIfNil(handler)
	}
	return route
}
