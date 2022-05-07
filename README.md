try
---
Go package that provides a generic `Try` monad inspired by [Scala](https://www.scala-lang.org/api/current/scala/util/Try.html).

`Try` is an alternative for functions that return a value and an error. Instead of returning two values but only initializing one of them, the function returns a `Try`.

There are two constructors, `Succeed` and `Fail`:

```go
func MightFail() Try[string] {
	err := foo()
	if err != nil {
		return Fail[string](Err: err)
	}
	return Succeed[string](Val: "foo")
}
// instead of:
func MightFail() (string, error) {
	err := foo()
	if err != nil {
		return "", err
	}
	return "foo", nil
}
```

And since it's a monad, you can use the `Bind` function to conditionally sequence code.


[errval](https://github.com/dnmfarrell/go-errval) was a previous attempt at a generic Either type.
