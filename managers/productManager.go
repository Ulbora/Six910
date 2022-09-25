package managers

import (
	"net/http"
	"sync"

	psort "github.com/Ulbora/Six910/prodsort"

	sdbi "github.com/Ulbora/six910-database-interface"
)

/*
 Six910 is a shopping cart and E-commerce system.

 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//AddProduct AddProduct
func (m *Six910Manager) AddProduct(p *sdbi.Product) *ResponseID {
	var rtn ResponseID
	suc, id := m.Db.AddProduct(p)
	if suc && id != 0 {
		rtn.Success = suc
		rtn.ID = id
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}

	return &rtn
}

//UpdateProduct UpdateProduct
func (m *Six910Manager) UpdateProduct(p *sdbi.Product) *Response {
	var rtn Response
	pd := m.Db.GetProductByID(p.ID)
	if pd.StoreID == p.StoreID {
		suc := m.Db.UpdateProduct(p)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//UpdateProductQuantity UpdateProductQuantity
func (m *Six910Manager) UpdateProductQuantity(p *sdbi.Product) *Response {
	var rtn Response
	pd := m.Db.GetProductByID(p.ID)
	if pd.StoreID == p.StoreID {
		suc := m.Db.UpdateProductQuantity(p)
		if suc {
			rtn.Success = suc
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetProductByID GetProductByID
func (m *Six910Manager) GetProductByID(id int64, sid int64) *sdbi.Product {
	var rtn *sdbi.Product
	var p *sdbi.Product
	var ss *[]sdbi.Product

	var wg sync.WaitGroup
	wg.Add(1)
	go func(pid int64) {
		defer wg.Done()
		p = m.Db.GetProductByID(pid)
	}(id)

	wg.Add(1)
	go func(ssid int64, pid int64) {
		defer wg.Done()
		ss = m.Db.GetProductSubSkuList(ssid, pid)
		var ps psort.ProductSort
		ss = ps.SortProds(*ss)
	}(sid, id)

	wg.Wait()

	if p.StoreID == sid {
		rtn = p
	} else {
		var np sdbi.Product
		rtn = &np
	}
	if len(*ss) > 0 {
		rtn.SubSkuList = ss
	} else {
		var ssls = []sdbi.Product{}
		rtn.SubSkuList = &ssls
	}

	return rtn
}

//GetProductByBySku GetProductByBySku
func (m *Six910Manager) GetProductByBySku(sku string, distributorID int64, sid int64) *sdbi.Product {
	var rtn *sdbi.Product
	p := m.Db.GetProductBySku(sku, distributorID, sid)
	if p.StoreID == sid {
		rtn = p
	} else {
		var np sdbi.Product
		rtn = &np
	}
	var ssls *[]sdbi.Product
	if rtn.ID != 0 {
		ssls = m.Db.GetProductSubSkuList(sid, rtn.ID)
	}
	if ssls != nil && len(*ssls) > 0 {
		rtn.SubSkuList = ssls
	} else {
		ssls = &[]sdbi.Product{}
		rtn.SubSkuList = ssls
	}

	return rtn
}

//GetProductsByPromoted GetProductsByPromoted
func (m *Six910Manager) GetProductsByPromoted(sid int64, start int64,
	end int64) *[]sdbi.Product {
	return m.Db.GetProductsByPromoted(sid, start, end)
}

//GetProductsByName GetProductsByName
func (m *Six910Manager) GetProductsByName(name string, sid int64, start int64,
	end int64) *[]sdbi.Product {
	return m.Db.GetProductsByName(name, sid, start, end)
}

//GetProductsByCaterory GetProductsByCaterory
func (m *Six910Manager) GetProductsByCaterory(catID int64, sid int64, start int64,
	end int64) *[]sdbi.Product {
	var rtn *[]sdbi.Product
	c := m.Db.GetCategory(catID)
	if c.StoreID == sid {
		rtn = m.Db.GetProductsByCaterory(catID, start, end)
	} else {
		var np = []sdbi.Product{}
		rtn = &np
	}
	return rtn
}

//GetProductList GetProductList
func (m *Six910Manager) GetProductList(storeID int64, start int64, end int64) *[]sdbi.Product {
	return m.Db.GetProductList(storeID, start, end)
}

//GetProductSubSkuList GetProductSubSkuList
func (m *Six910Manager) GetProductSubSkuList(storeID int64, parentProdID int64) *[]sdbi.Product {
	return m.Db.GetProductSubSkuList(storeID, parentProdID)
}

//GetProductIDList GetProductIDList
func (m *Six910Manager) GetProductIDList(sid int64) *[]int64 {
	return m.Db.GetProductIDList(sid)
}

//GetProductIDListByCategories GetProductIDListByCategories
func (m *Six910Manager) GetProductIDListByCategories(sid int64, catList *[]int64) *[]int64 {
	return m.Db.GetProductIDListByCategories(sid, catList)
}

//DeleteProduct DeleteProduct
func (m *Six910Manager) DeleteProduct(id int64, sid int64) *Response {
	var rtn Response
	p := m.Db.GetProductByID(id)
	if p.StoreID == sid {
		suc := m.Db.DeleteProduct(id)
		rtn.Success = suc
		rtn.Code = http.StatusOK
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
