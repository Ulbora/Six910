package prodsort

import (
	"fmt"
	"testing"
)

func TestAlphaNum_Sort(t *testing.T) {
	var alphaNum AlphaNum

	anum := alphaNum.Get()
	comp := anum.Sort("4XL", "XL")
	fmt.Println("comp:", comp)

	if comp {
		t.Fail()
	}
}

func TestAlphaNum_Sort2(t *testing.T) {
	var alphaNum AlphaNum

	anum := alphaNum.Get()
	comp := anum.Sort("XL", "5XL")
	fmt.Println("comp:", comp)

	if !comp {
		t.Fail()
	}
}
