package managers

import (
	"net/http"

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

//AddOrderTransaction AddOrderTransaction
func (m *Six910Manager) AddOrderTransaction(t *sdbi.OrderTransaction, sid int64) *ResponseID {
	var rtn ResponseID
	fo := m.Db.GetOrder(t.OrderID)
	if fo.StoreID == sid {
		suc, id := m.Db.AddOrderTransaction(t)
		if suc && id != 0 {
			rtn.Success = suc
			rtn.ID = id
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}

//GetOrderTransactionList GetOrderTransactionList
func (m *Six910Manager) GetOrderTransactionList(orderID int64, sid int64) *[]sdbi.OrderTransaction {
	var rtn *[]sdbi.OrderTransaction
	fo := m.Db.GetOrder(orderID)
	if fo.StoreID == sid {
		rtn = m.Db.GetOrderTransactionList(orderID)
	} else {
		var nt = []sdbi.OrderTransaction{}
		rtn = &nt
	}
	return rtn
}
