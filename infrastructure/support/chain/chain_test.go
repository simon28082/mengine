package chain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildChain2(t *testing.T) {
	var steps = make([]int, 0)
	err := BuildChain(func(next HandlerFunc) (HandlerFunc, error) {
		steps = append(steps, 1)
		fmt.Println("step1")
		return next, nil
	}, func(next HandlerFunc) (HandlerFunc, error) {
		steps = append(steps, 2)
		fmt.Println("step2")
		return next, nil
	}, func(next HandlerFunc) (HandlerFunc, error) {
		steps = append(steps, 3)
		fmt.Println("step3")
		return next, nil
	})
	assert.Equal(t, []int{1, 2, 3}, steps)
	assert.Nil(t, err)
}
