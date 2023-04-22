package chain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildChain(t *testing.T) {
	var steps = make([]int, 0)
	mw1 := func(next ChainParameterFunc) ChainParameterFunc {
		return func(input any) (interface{}, error) {
			input = fmt.Sprintf(`%s-mw1before`, input.(string))
			result, err := next(input)
			assert.Nil(t, err)
			fmt.Println("middlewareFunc1: after")
			steps = append(steps, 3)
			return result, err
		}
	}

	mw2 := func(next ChainParameterFunc) ChainParameterFunc {
		return func(input interface{}) (interface{}, error) {
			input = fmt.Sprintf(`%s-mw2before`, input.(string))
			result, err := next(input)
			fmt.Println("middlewareFunc2: after")
			steps = append(steps, 2)
			return result, err
		}
	}

	myMiddleware := func(next ChainParameterFunc) ChainParameterFunc {
		return func(input interface{}) (interface{}, error) {
			input = fmt.Sprintf(`%s-MyMiddlewareBefore`, input.(string))
			result, err := next(input)
			fmt.Println("MyMiddleware: after")
			steps = append(steps, 1)
			return result, err
		}
	}

	result, err := BuildChainParameter("hello", mw1, mw2, myMiddleware)
	assert.Nil(t, err)
	assert.Equal(t, `hello-mw1before-mw2before-MyMiddlewareBefore`, result)
	assert.Equal(t, []int{1, 2, 3}, steps)

	fmt.Println(result, err)
}
