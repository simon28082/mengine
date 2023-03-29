package chain

//type Input any
//
//type Next any

type Chain func(input any, next Chain) Chain

func Arrange(input any, chinas ...Chain) Chain {
	for i := range chinas {
		v := chinas(input)
	}
}
