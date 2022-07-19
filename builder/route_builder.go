package builder

import (
	cameltogo "github.com/MotaOcimar/camel-to-go"
	"github.com/MotaOcimar/camel-to-go/model"
)

type RouteBuilder struct {
	BuilderSupport
	routeCollection *model.RoutesDefinition
}

// Setters and Getters

func (builder *RouteBuilder) GetRouteCollection() *model.RoutesDefinition {
	return builder.routeCollection
}

func (builder *RouteBuilder) SetRouteCollection(routeCollection model.RoutesDefinition) {
	*builder.routeCollection = routeCollection
}

// Public Methods

func Configure() {}

func (builder *RouteBuilder) From(uri string) (answer *model.RouteDefinition) {
	builder.GetRouteCollection().SetContext(builder.GetContext())
	answer = builder.GetRouteCollection().From(uri)
	builder.configureRoute(answer)
	return answer
}

// Private Methods

func (builder *RouteBuilder) configureRoute(route *model.RouteDefinition) {
	cameltogo.TrySetContext(route, builder.GetContext())
}
