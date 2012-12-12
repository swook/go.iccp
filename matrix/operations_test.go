package matrix

import (
	"testing"
)

func TestOperations(t *testing.T) {
	A, err := New("1,2,3;4,5,6")
	B, err := New("4,5,6;7,8,9")
	_, err = A.Add(B)
	if err != nil {
		t.Fatal(err)
	}

	_, err = A.Add(1.2)
	if err != nil {
		t.Fatal(err)
	}

	E, err := New("1;2;3")
	F, err := A.Multiply(E)
	if err != nil {
		t.Fatal(err)
	}

	_, err = F.OuterProduct(E)
	if err != nil {
		t.Fatal(err)
	}

	H, err := New("3,5,7;2,4,8;1,6,9")
	H, _ = H.Inverse()
	H, _ = H.Inverse()

	I, err := New("3,5;2,4")
	I, _ = I.Inverse()
	I, _ = I.Inverse()

	K, err := New("30,11,12,13;14,35,16,17;18,19,20,21;22,23,24,55")
	K, _ = K.Inverse()
	K, _ = K.Inverse()
}
