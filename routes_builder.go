package cameltogo

type IRouteBuilder interface {
	// TODO
	AddRoutesToContext(context *Context) error
	Configure()
}
