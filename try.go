// Package try provides an either monad for bundling return vales and
// exceptions. It includes standard functions like bind, fmap and return.
package try

// Try is a monadic either type which contains a generic value or an error.
type Try[A any] struct {
	Err error
	Val A
}

// New converts a function that returns a value and an error into one that
// returns a Try.
func New[A, B any](f func(a A) (B, error)) func(A) Try[B] {
	return func(a A) Try[B] {
		val, err := f(a)
		if err == nil {
			return Succeed[B](val)
		}
		return Fail[B](err)
	}
}

// Fail constructs a Try that wraps an error (aka "return").
func Fail[A any](err error) Try[A] { return Try[A]{Err: err} }

// Succeed constructs a Try that wraps a value (aka "return").
func Succeed[A any](val A) Try[A] { return Try[A]{Val: val} }

// Fmap applies a Try to a non-monadic function, wrapping the return value as a
// Try. The function is only called when the Try is not an error.
func Fmap[A, B any](a Try[A], f func(A) B) Try[B] {
	if a.Err == nil {
		return Succeed[B](f(a.Val))
	}
	return Fail[B](a.Err)
}

// Bind applies a Try to a function which accepts a value and returns a Try.
// The function is only called if the Try is not an error.
func Bind[A, B any](a Try[A], f func(A) Try[B]) Try[B] {
	if a.Err == nil {
		return f(a.Val)
	}
	return Fail[B](a.Err)
}

// Compose takes two functions and combines them into a new function.
func Compose[A, B, C any](f func(A) Try[B], g func(B) Try[C]) func(A) Try[C] {
	return func(a A) Try[C] { return Bind(f(a), g) }
}
