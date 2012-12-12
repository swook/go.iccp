package matrix

import (
	"testing"
)

func TestOperations(t *testing.T) {
	A, err := New("1,2,3;4,5,6")
	B, err := New("4,5,6;7,8,9")
	C, err := A.Add(B)
	if err != nil {
		t.Fatal(err)
	}
	println(C.String())

	D, err := A.Add(1.2)
	if err != nil {
		t.Fatal(err)
	}
	println(D.String())

	E, err := New("1;2;3")
	F, err := A.Multiply(E)
	if err != nil {
		t.Fatal(err)
	}
	println(F.String())

	G, err := F.OuterProduct(E)
	if err != nil {
		t.Fatal(err)
	}
	println(G.String())
}
