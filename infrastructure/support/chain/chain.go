package chain

import "fmt"

type ChainFunc func() error

type Chain func(next ChainFunc) ChainFunc

func BuildChain(mw ...Chain) error {
	if len(mw) == 0 {
		return fmt.Errorf("middleware chain is empty")
	}

	var (
		current ChainFunc = func() error {
			return nil
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
		return err
	}

	// Execute the middleware chain.
	return current()
}
