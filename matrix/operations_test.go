package matrix

import (
	"testing"
	"fmt"
)

func TestOperations(t *testing.T) {
	A, _ := New("1,2,3;4,5,6")
	B, _ := New("4,5,6;7,8,9")
	C, _ := A.Add(B)
}