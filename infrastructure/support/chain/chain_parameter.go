package chain

import "fmt"

type ChainParameterFunc func(input any) (any, error)

type ChainParameter func(next ChainParameterFunc) ChainParameterFunc

func BuildChainParameter(input any, mw ...ChainParameter) (any, error) {
	if len(mw) == 0 {
		return nil, fmt.Errorf("middleware chain is empty")
	}

	var (
		current ChainParameterFunc = func(input any) (any, error) {
			return input, nil
		}
		err error
	)

	// Build the middleware chain in reverse order.
	for i := len(mw) - 1; i >= 0; i-- {
		current = mw[i](current)
		if current == nil {
			err = fmt.Errorf("middleware %d returned a nil handler", i)
			break
		}
	}

	if err != nil {
		return nil, err
	}

	// Execute the middleware chain.
	result, err := current(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}
