package try

import (
	"testing"
)

func TestMonadLaws(t *testing.T) {
	// Left identity: "return bind is neutral".
	// return x >>= f eq f x
	x := 5
	f := func(i int) Try[int] { return Succeed[int](i * 2) }
	if Bind[int](Succeed[int](x), f).Val != f(x).Val {
		t.Errorf("Does not obey left identity law")
	}
	// Right identity: "bind return is neutral".
	// m >>= return eq m
	m := Succeed[int](x)
	if Bind[int](m, Succeed[int]).Val != m.Val {
		t.Errorf("Does not obey right identity law")
	}
	// Associativity: "bind order does not matter".
	// (m >>= f) >>= g eq m >>= (\x -> f x >>= g)
	g := func(i int) Try[int] { return Succeed[int](i / 2) }
	if Bind[int](Bind[int](m, f), g).Val != Bind[int](f(x), g).Val {
		t.Error("Does not obey associativity law")
	}
}
