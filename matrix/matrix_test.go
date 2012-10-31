package matrix

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	M, err := New(3)
	if err != nil {
		t.Fatal(err)
	}

	M, err = New(3,5)
	if err != nil {
		t.Fatal(err)
	}

	M, err = New(`1   3
				  4.0 5.1 6.2
				  1`)
	if err != nil {
		t.Fatal(err)
	}

	M, err = New(`[1;4.0 5.1 6.2]`)
	if err != nil {
		t.Fatal(err)
	}

	M.String()
}
