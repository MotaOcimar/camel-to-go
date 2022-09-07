package cameltogo

import (
	log "github.com/sirupsen/logrus"
)

type Context struct {
	BaseService

	firstStartDone bool
	// TODO
}

func (context *Context) Start() {
	// TODO
}

func (context *Context) AddRoutes(builder IRouteBuilder) error {

	Build(context)
	log.Debugf("Adding routes from builder: %+v", builder)
	err := builder.AddRoutesToContext(context)

	return err
}

// Implements IBaseService

func (context *Context) DoBuild() (err error) {
	// TODO
	return err
}

func (context *Context) DoFail(err error) {
	context.firstStartDone = false
	panic(err)
}
