package types

type Calc[T any] func(string) T
type Result[T any] func(ori, val T) T
type Transfer[T any] func(<-chan string, Calc[T]) <-chan T
