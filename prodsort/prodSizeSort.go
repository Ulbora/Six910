package prodsort

import (
	"sort"

	sdbi "github.com/Ulbora/six910-database-interface"
)

//ProcuctSizeSort ProcuctSizeSort
type ProcuctSizeSort []sdbi.Product

//Len Len
func (p ProcuctSizeSort) Len() int {
	return len(p)
}

//Swap Swap
func (p ProcuctSizeSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ProcuctSizeSort) Less(i, j int) bool {
	var rtn bool

	var v1 = p[i].Size
	var v2 = p[j].Size

	var alpha AlphaNum
	anu := alpha.Get()
	rtn = anu.Sort(v1, v2)

	return rtn
}

//ProductSort ProductSort
type ProductSort struct {
}

//SortProds Sort
func (s ProductSort) SortProds(prods []sdbi.Product) *[]sdbi.Product {
	sort.Sort(ProcuctSizeSort(prods))

	return &prods
}
