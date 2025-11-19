package logs

import(
        "testing"
	"os"
	"slices"
)

// test get input
func TestGetInput(t *testing.T) {
        old := os.Stdin
	defer func() {os.Stdin = old}()

        r, w, err := os.Pipe()
	if err != nil {
	        t.Fatal(err)
	}
	os.Stdin = r

        go func() {
	        defer w.Close()
		w.Write([]byte("   run  hello    world    \n"))
	}()

        res := GetInput()
	exp := []string{"run", "hello", "world"}
	if !slices.Equal(res, exp) {
	        t.Errorf("expected: %v, got: %v", exp, res)
	}
}