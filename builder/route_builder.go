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

func (builder *RouteBuilder) AddRoutesToCamelContext(context *cameltogo.Context) error {
	// must configure routes before rests
	err := builder.ConfigureRoutes(context)
	if err != nil {
		return err
	}

	err = builder.ConfigureRests(context)
	if err != nil {
		return err
	}

	// but populate rests before routes, as we want to turn rests into routes
	err = builder.populateRests()
	if err != nil {
		return err
	}

	// ensure routes are prepared before being populated
	for _, route := range builder.routeCollection.GetRoutes() {
		builder.routeCollection.PrepareRoute(route)
	}
	err = builder.populateRoutes()
	if err != nil {
		return err
	}

	return err
}

func (builder *RouteBuilder) ConfigureRoutes(context *cameltogo.Context) (err error) {
	// TODO
	return err
}

func (builder *RouteBuilder) ConfigureRests(context *cameltogo.Context) (err error) {
	// TODO
	return err
}

func (builder *RouteBuilder) populateRoutes() (err error) {
	// TODO
	return err
}

func (builder *RouteBuilder) populateRests() (err error) {
	// TODO
	return err
}

// Private Methods

func (builder *RouteBuilder) configureRoute(route *model.RouteDefinition) {
	cameltogo.TrySetContext(route, builder.GetContext())
}
