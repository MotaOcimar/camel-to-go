package cameltogo

type ContextAware interface {
	GetContext() *Context
	SetContext(context *Context)
}

func TrySetContext[T ContextAware](object T, context *Context) T {
	if context != nil {
		object.SetContext(context)
	}
	return object
}
