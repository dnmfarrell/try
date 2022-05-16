package try

type Try[A any] struct {
	Err error
	Val A
}

func Fmap[A, B any](a Try[A], f func(a A) Try[B]) Try[B] {
	if a.Err == nil {
		return f(a.Val)
	}
	return Fail[B](a.Err)
}

func Fmap2[A, B, C any](a Try[A], b Try[B], f func(a A, b B) Try[C]) Try[C] {
	if a.Err == nil {
		if b.Err == nil {
			return f(a.Val, b.Val)
		}
		return Fail[C](b.Err)
	}
	return Fail[C](a.Err)
}

func Compose[A, B, C any](f func(a A) Try[B], g func(b B) Try[C]) func(A) Try[C] {
	return func(a A) Try[C] { return Fmap(f(a), g) }
}

func Map[A, B any](a Try[A], f func(A) B) Try[B] {
	if a.Err == nil {
		return Succeed[B](f(a.Val))
	}
	return Fail[B](a.Err)
}

func Map2[A, B, C any](a Try[A], b Try[B], f func(A, B) C) Try[C] {
	if a.Err == nil {
		if b.Err == nil {
			return Succeed[C](f(a.Val, b.Val))
		}
		return Fail[C](b.Err)
	}
	return Fail[C](a.Err)
}

func New[A, B any](f func(a A) (B, error)) func(A) Try[B] {
	return func(a A) Try[B] {
		val, err := f(a)
		if err == nil {
			return Succeed[B](val)
		}
		return Fail[B](err)
	}
}

func New2[A, B, C any](f func(a A, b B) (C, error)) func(A, B) Try[C] {
	return func(a A, b B) Try[C] {
		val, err := f(a, b)
		if err == nil {
			return Succeed[C](val)
		}
		return Fail[C](err)
	}
}

func Lift[A, B any](f func(A) B) func(Try[A]) Try[B] {
	return func(a Try[A]) Try[B] { return Map(a, f) }
}

func Lift2[A, B, C any](f func(A, B) C) func(Try[A], Try[B]) Try[C] {
	return func(a Try[A], b Try[B]) Try[C] { return Map2(a, b, f) }
}

func Fail[A any](err error) Try[A] { return Try[A]{Err: err} }

func Succeed[A any](val A) Try[A] { return Try[A]{Val: val} }
