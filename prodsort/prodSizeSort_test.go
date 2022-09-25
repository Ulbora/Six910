package prodsort

import (
	"fmt"
	"testing"

	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestProductSort_SortProds(t *testing.T) {
	var prods []sdbi.Product

	var p1 sdbi.Product
	p1.ID = 1
	p1.Size = "5XL"

	var p2 sdbi.Product
	p2.ID = 2
	p2.Size = "3XL"

	var p3 sdbi.Product
	p3.ID = 3
	p3.Size = "XL"

	var p4 sdbi.Product
	p4.ID = 4
	p4.Size = "7XL"

	var p5 sdbi.Product
	p5.ID = 5
	p5.Size = "21"

	var p6 sdbi.Product
	p6.ID = 6
	p6.Size = "26"

	prods = append(prods, p1)
	prods = append(prods, p2)
	prods = append(prods, p3)
	prods = append(prods, p4)
	prods = append(prods, p5)
	prods = append(prods, p6)

	var ps ProductSort
	sps := ps.SortProds(prods)

	for i := range *sps {
		fmt.Print("ID: ", (*sps)[i].ID)
		fmt.Println(" size: ", (*sps)[i].Size)
	}
	if sps == nil {
		t.Fail()
	}

}
