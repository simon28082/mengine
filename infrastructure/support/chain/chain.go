package chain

type HandlerFunc func(next HandlerFunc) (HandlerFunc, error)

func BuildChain(middlewares ...HandlerFunc) error {
	var (
		handler HandlerFunc = func(handlerFunc HandlerFunc) (HandlerFunc, error) {
			return nil, nil
		}
		err error
	)
	for i := 0; i < len(middlewares); i++ {
		handler, err = middlewares[i](handler)
		if err != nil {
			return err
		}
	}
	return nil
}
