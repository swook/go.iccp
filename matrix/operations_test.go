package matrix

import (
	"testing"
)

func TestOperations(t *testing.T) {
	A, err := New("1,2,3;4,5,6")
	B, err := New("4,5,6;7,8,9")
	err = A.Add(B)
	if err != nil {
		t.Fatal(err)
	}

	err = A.Add(1.2)
	if err != nil {
		t.Fatal(err)
	}

	C, err := New("1;2;3")
	err = A.Multiply(C)
	if err != nil {
		t.Fatal(err)
	}
}
