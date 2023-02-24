package try

import (
	"errors"
	"strconv"
	"testing"
)

func TestMonadLaws(t *testing.T) {
	// Left identity: "return bind is neutral".
	// return x >>= f eq f x
	x := 5
	f := func(i int) Try[int] { return Succeed(i * 2) }
	if Bind(Succeed(x), f).Val != f(x).Val {
		t.Errorf("Does not obey left identity law")
	}
	// Right identity: "bind return is neutral".
	// m >>= return eq m
	m := Succeed(x)
	if Bind(m, Succeed[int]).Val != m.Val {
		t.Errorf("Does not obey right identity law")
	}
	// Associativity: "bind order does not matter".
	// (m >>= f) >>= g eq m >>= (\x -> f x >>= g)
	g := func(i int) Try[int] { return Succeed(i / 2) }
	if Bind(Bind(m, f), g).Val != Bind(f(x), g).Val {
		t.Error("Does not obey associativity law")
	}
}

func TestNew(t *testing.T) {
	tryParseBool := New(strconv.ParseBool)
	if !tryParseBool("true").Val {
		t.Error("Monadic parseBool didn't return parse value")
	}
	if tryParseBool("foo").Err == nil {
		t.Error("Monadic parseBool didn't return parse error")
	}
}

func TestFmap(t *testing.T) {
	i := func(s string) string { return s }
	s := Succeed[string]("x")
	if Fmap(s, i).Val != "x" {
		t.Error("Fmap didn't apply monad")
	}
	e := Fail[string](errors.New(""))
	if Fmap(e, i).Err == nil {
		t.Error("Fmap didn't return an error")
	}
}

func TestBind(t *testing.T) {
	i := func(s string) Try[string] { return Succeed(s) }
	s := Succeed[string]("x")
	if Bind(s, i).Val != "x" {
		t.Error("Bind didn't apply monad")
	}
	e := Fail[string](errors.New(""))
	if Bind(e, i).Err == nil {
		t.Error("Bind didn't return an error")
	}
}
