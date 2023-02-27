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

func TestLift(t *testing.T) {
	if !Lift("true", strconv.ParseBool).Val {
		t.Error("Lifted function didn't return value")
	}
	if Lift("x", strconv.ParseBool).Err == nil {
		t.Error("Lifted function didn't return error")
	}
}

func TestLift2(t *testing.T) {
	f := func(x, y int) (int, error) {
		if x+y > 9 {
			return 0, errors.New("too big")
		}
		return x + y, nil
	}
	if Lift2(1, 5, f).Val != 6 {
		t.Error("Lifted function didn't return value")
	}
	if Lift2(7, 5, f).Err == nil {
		t.Error("Lifted function didn't return error")
	}
}

func TestLift3(t *testing.T) {
	if Lift3("1", 10, 32, strconv.ParseInt).Val != 1 {
		t.Error("Lifted function didn't return value")
	}
	if Lift3("x", 10, 32, strconv.ParseInt).Err == nil {
		t.Error("Lifted function didn't return error")
	}
}

func TestMap(t *testing.T) {
	i := func(s string) string { return s }
	s := Succeed[string]("x")
	if Map(s, i).Val != "x" {
		t.Error("Map didn't apply monad")
	}
	e := Fail[string](errors.New(""))
	if Map(e, i).Err == nil {
		t.Error("Map didn't return an error")
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
