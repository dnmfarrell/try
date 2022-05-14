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

func TestLift(t *testing.T) {
	f := func(i int) int { return i + 2 }
	g := Lift(f)
	if g(Succeed[int](5)).Val != 7 {
		t.Error("Map doesn't call lifted function on success")
	}
	if g(Fail[int](errors.New("ruhroh"))).Val == 2 {
		t.Error("Map calls lifted function on failure")
	}
}

func TestLift2(t *testing.T) {
	f := func(x int, y int) int { return x + y }
	g := Lift2(f)
	if g(Succeed[int](5), Succeed[int](4)).Val != 9 {
		t.Error("Map2 doesn't call lifted function on success")
	}
	if g(Fail[int](errors.New("ruhroh")), Succeed[int](4)).Val == 4 {
		t.Error("Map2 calls lifted function on Fail[A]")
	}
	if g(Succeed[int](4), Fail[int](errors.New("ruhroh"))).Val == 4 {
		t.Error("Map2 calls lifted function on Fail[B]")
	}
}
