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

And since it's a monad, you can use higher order functions like `Fmap` and `Map` to conditionally sequence code. Here is the `CopyFile` function from the [Error Handling - Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md) rewritten using `Try`:

```go
func ExampleCopyFile(src, dst string) error {
  clo := func (f *os.File) error { return f.Close() }
  rdr := New(os.Open)(src)
  defer Map(rdr, clo)
  wtr := New(os.Create)(dst)
  cpy := New2(func (w, r *os.File) (int64, error) { return w.ReadFrom(r) })
  err := Fmap2(wtr, rdr, cpy).Err
  if err != nil {
    Map(wtr, clo)
    Map(wtr, func (_ *os.File) error { return os.Remove(dst) })
    return fmt.Errorf("copy %s %s: %v", src, dst, err)
  }
  Map(wtr, clo)
  return nil
}
```

[errval](https://github.com/dnmfarrell/go-errval) was a previous attempt at a generic Either type.
