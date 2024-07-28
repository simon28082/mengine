package option

type Option[T any] func(T)

func Apply[T any](t T, options ...Option[T]) {
	for i := range options {
		options[i](t)
	}
}

//
//func Initial[T any](options ...Option[T]) T {
//	var t = new(T)
//	for i := range options {
//		options[i](t)
//	}
//	return t
//}
