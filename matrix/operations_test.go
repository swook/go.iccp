package matrix

import (
	"fmt"
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
	fmt.Println(A)

	err = A.Add("1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(A)
}
