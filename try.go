package try

type Try[A any] struct {
	Err error
	Val A
}

func Bind[A, B any](t Try[A], f func(v A) Try[B]) Try[B] {
	if t.Err == nil {
		return f(t.Val)
	}
	return Fail[B](t.Err)
}

func Compose[A, B, C any](f func(a A) Try[B], g func(b B) Try[C]) func(A) Try[C] {
	return func(a A) Try[C] { return Bind(f(a), g) }
}

func Fail[A any](err error) Try[A] { return Try[A]{Err: err} }
func Succeed[A any](val A) Try[A]  { return Try[A]{Val: val} }
