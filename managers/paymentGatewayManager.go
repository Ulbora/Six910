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

//AddPaymentGateway AddPaymentGateway
func (m *Six910Manager) AddPaymentGateway(pgw *sdbi.PaymentGateway, sid int64) *ResponseID {
	var rtn ResponseID
	g := m.Db.GetStorePlugin(pgw.StorePluginsID)
	if g.StoreID == sid {
		suc, id := m.Db.AddPaymentGateway(pgw)
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

//UpdatePaymentGateway UpdatePaymentGateway
func (m *Six910Manager) UpdatePaymentGateway(pgw *sdbi.PaymentGateway, sid int64) *Response {
	var rtn Response
	p := m.Db.GetStorePlugin(pgw.StorePluginsID)
	if p.StoreID == sid {
		suc := m.Db.UpdatePaymentGateway(pgw)
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

//GetPaymentGateway GetPaymentGateway
func (m *Six910Manager) GetPaymentGateway(id int64, sid int64) *sdbi.PaymentGateway {
	var rtn *sdbi.PaymentGateway
	gw := m.Db.GetPaymentGateway(id)
	p := m.Db.GetStorePlugin(gw.StorePluginsID)
	if p.StoreID == sid {
		rtn = gw
	} else {
		var ng sdbi.PaymentGateway
		rtn = &ng
	}
	return rtn
}

//GetPaymentGatewayByName GetPaymentGatewayByName
func (m *Six910Manager) GetPaymentGatewayByName(name string, sid int64) *sdbi.PaymentGateway {
	return m.Db.GetPaymentGatewayByName(name, sid)
}

//GetPaymentGateways GetPaymentGateways
func (m *Six910Manager) GetPaymentGateways(storeID int64) *[]sdbi.PaymentGateway {
	return m.Db.GetPaymentGateways(storeID)
}

//DeletePaymentGateway DeletePaymentGateway
func (m *Six910Manager) DeletePaymentGateway(id int64, sid int64) *Response {
	var rtn Response
	gw := m.Db.GetPaymentGateway(id)
	p := m.Db.GetStorePlugin(gw.StorePluginsID)
	if p.StoreID == sid {
		suc := m.Db.DeletePaymentGateway(id)
		if suc {
			rtn.Success = true
			rtn.Code = http.StatusOK
		} else {
			rtn.Code = http.StatusBadRequest
		}
	} else {
		rtn.Code = http.StatusBadRequest
	}
	return &rtn
}
