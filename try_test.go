package try

import (
	"errors"
	"testing"
)

func TestMonadLaws(t *testing.T) {
	// Compose(f, unit) == f
	a := 5
	f := func(i int) Try[int] { return Succeed[int](i * 2) }
	if Compose[int, int, int](f, Succeed[int])(a).Val != f(a).Val {
		t.Errorf("Success does not obey left identity law")
	}
	// Compose(unit, f) == f
	if Compose[int, int, int](Succeed[int], f)(a).Val != f(a).Val {
		t.Errorf("Success does not obey right identity law")
	}
	// bind(bind(m,f), g) == bind(m, fg)
	m := Succeed[int](a)
	g := func(i int) Try[int] { return Succeed[int](i / 2) }
	if Bind(Bind(m, f), g).Val != Bind(m, Compose[int, int, int](f, g)).Val {
		t.Error("Success does not obey associativity law")
	}
	M := Fail[int](errors.New("womp womp"))
	if Bind(Bind(M, f), g).Err != Bind(M, Compose[int, int, int](f, g)).Err {
		t.Error("Failure does not obey associativity law")
	}
}
