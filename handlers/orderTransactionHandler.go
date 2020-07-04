package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
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

//OrderTransactionReq OrderTransactionReq
type OrderTransactionReq struct {
	StoreID          int64                 `json:"storeId"`
	OrderTransaction sdbi.OrderTransaction `json:"orderTransaction"`
}

//AddOrderTransaction AddOrderTransaction
func (h *Six910Handler) AddOrderTransaction(w http.ResponseWriter, r *http.Request) {
	var addortURL = "/six910/rs/orderTransaction/add"
	var aortc jv.Claim
	aortc.Role = customerRole
	aortc.URL = addortURL
	aortc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aortc)
	h.Log.Debug("order transaction add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aort OrderTransactionReq
			aortsuc, aorterr := h.ProcessBody(r, &aort)
			h.Log.Debug("aortsuc: ", aortsuc)
			h.Log.Debug("aort: ", aort)
			h.Log.Debug("aorerr: ", aorterr)
			if !aortsuc && aorterr != nil {
				http.Error(w, aorterr.Error(), http.StatusBadRequest)
			} else {
				aortres := h.Manager.AddOrderTransaction(&aort.OrderTransaction, aort.StoreID)
				h.Log.Debug("aortres: ", *aortres)
				if aortres.Success && aortres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aortres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aortf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aortf)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetOrderTransactionList GetOrderTransactionList
func (h *Six910Handler) GetOrderTransactionList(w http.ResponseWriter, r *http.Request) {
	var gortlURL = "/six910/rs/orderTransaction/list"
	var gortcl jv.Claim
	gortcl.Role = customerRole
	gortcl.URL = gortlURL
	gortcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gortcl)
	h.Log.Debug("order transaction get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var ortlordIDStr = vars["orderId"]
			var ortlstoreIDStr = vars["storeId"]
			orderID, sortloiderr := strconv.ParseInt(ortlordIDStr, 10, 64)
			storeID, sortlserr := strconv.ParseInt(ortlstoreIDStr, 10, 64)
			var gortlres *[]sdbi.OrderTransaction
			if sortloiderr == nil && sortlserr == nil {
				gortlres = h.Manager.GetOrderTransactionList(orderID, storeID)
				h.Log.Debug("get order transaction list: ", gortlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.OrderTransaction{}
				gortlres = &nc
			}
			resJSON, _ := json.Marshal(gortlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
