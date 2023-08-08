// Package yc implements a Y Combinator
package yc

type (
	// Function is a basic function that takes one input and returns the input
	Function[T any] func(T) T
	// Transformer is a function that takes a function in T and returns a function in T
	Transformer[T any] Function[Function[T]]
)

// recursiveSimplifier is a function that takes a RecursiveSimplifier in T and returns a function in T
type recursiveSimplifier[T any] func(recursiveSimplifier[T]) Function[T]

// Y = λf.(λx.f(x x))(λx.f(x x))
func Y[T any](f Transformer[T]) Function[T] {
	g := func(h recursiveSimplifier[T]) Function[T] {
		return func(x T) T {
			return (f(h(h))(x))
		}
	}
	return g(g)
}
