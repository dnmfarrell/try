// Package try provides an either monad for bundling return values and
// errors. It includes standard functions like bind, map, lift and return.
package try

// Try is a monadic either type which contains a generic value or an error.
type Try[A any] struct {
	Val A
	Err error
}

// Fail constructs a Try that wraps an error (aka "return").
func Fail[A any](err error) Try[A] { return Try[A]{Err: err} }

// Succeed constructs a Try that wraps a value (aka "return").
func Succeed[A any](val A) Try[A] { return Try[A]{Val: val} }

// Map applies a Try to a non-monadic function, wrapping the return value as a
// Try. The function is only called when the Try is not an error.
func Map[A, B any](a Try[A], f func(A) B) Try[B] {
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

// Lift wraps the return values of func(A) (B,error) into Try[B].
func Lift[A, B any](a A, f func(A) (B, error)) Try[B] {
	b, e := f(a)
	return Try[B]{b, e}
}

// Lift2 wraps the return values of func(A,B) (C,error) into Try[C].
func Lift2[A, B, C any](a A, b B, f func(A, B) (C, error)) Try[C] {
	c, e := f(a, b)
	return Try[C]{c, e}
}

// Lift3 wraps the return values of func(A,B,C) (D,error) into Try[D].
func Lift3[A, B, C, D any](a A, b B, c C, f func(A, B, C) (D, error)) Try[D] {
	d, e := f(a, b, c)
	return Try[D]{d, e}
}
