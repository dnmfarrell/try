Go package that provides a generic `Try` monad inspired by [Scala](https://www.scala-lang.org/api/current/scala/util/Try.html).

Try is an alternative for functions that return a value and an error. Instead of returning two values but only initializing one of them, the function returns a `Try`.

It includes standard functions like bind, map, lift and return.

    type Try[A any] struct{ ... }
        func Bind[A, B any](a Try[A], f func(A) Try[B]) Try[B]
        func Fail[A any](err error) Try[A]
        func Lift[A, B any](a A, f func(A) (B, error)) Try[B]
        func Lift2[A, B, C any](a A, b B, f func(A, B) (C, error)) Try[C]
        func Lift3[A, B, C, D any](a A, b B, c C, f func(A, B, C) (D, error)) Try[D]
        func Map[A, B any](a Try[A], f func(A) B) Try[B]
        func Succeed[A any](val A) Try[A]

This [blog post](https://blog.dnmfarrell.com/post/monads-simplify-go-error-handling/) explains the motivation behind using Try.

License
-------
Copyright 2023 David Farrell

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
