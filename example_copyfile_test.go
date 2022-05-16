package try_test

import (
	"fmt"
	"github.com/dnmfarrell/try"
	"os"
	"testing"
)

// ExampleCopyFile uses Try to condense the CopyFile function from the
// Error Handling Problem Overview:
// https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md
func CopyFile(src, dst string) error {
	clo := func(f *os.File) error { return f.Close() }
	rdr := try.New(os.Open)(src)
	defer try.Map(rdr, clo)
	wtr := try.New(os.Create)(dst)
	cpy := try.New2(func(w, r *os.File) (int64, error) { return w.ReadFrom(r) })
	err := try.Fmap2(wtr, rdr, cpy).Err
	if err != nil {
		try.Map(wtr, clo)
		try.Map(wtr, func(_ *os.File) error { return os.Remove(dst) })
		return fmt.Errorf("copy %s %s: %v", src, dst, err)
	}
	try.Map(wtr, clo)
	return nil
}

func Example() {
	src := "/tmp/try-test-orig.text"
	dst := "/tmp/try-test-copy.text"
	err := CopyFile(src, dst)
	if err != nil {
		fmt.Printf("CopyFile returned an error: %e\n", err)
	}
	return
}

func TestExample(t *testing.T) {
	Example()
}
