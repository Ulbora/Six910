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

//ProductCategoryReq ProductCategoryReq
type ProductCategoryReq struct {
	StoreID         int64                `json:"storeId"`
	ProductCategory sdbi.ProductCategory `json:"productCategory"`
}

//AddProductCategory AddProductCategory
func (h *Six910Handler) AddProductCategory(w http.ResponseWriter, r *http.Request) {
	var addpcURL = "/six910/rs/productCategory/add"
	var apcc jv.Claim
	apcc.Role = storeAdmin
	apcc.URL = addpcURL
	apcc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &apcc)
	h.Log.Debug("product category add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var apcr ProductCategoryReq
			apcsuc, apcerr := h.ProcessBody(r, &apcr)
			h.Log.Debug("apcsuc: ", apcsuc)
			h.Log.Debug("apcr: ", apcr)
			h.Log.Debug("apcerr: ", apcerr)
			if !apcsuc && apcerr != nil {
				http.Error(w, apcerr.Error(), http.StatusBadRequest)
			} else {
				apcres := h.Manager.AddProductCategory(&apcr.ProductCategory, apcr.StoreID)
				h.Log.Debug("apcres: ", *apcres)
				if apcres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(apcres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var apcf m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(apcf)
		fmt.Fprint(w, string(resJSON))
	}
}

//DeleteProductCategory DeleteProductCategory
func (h *Six910Handler) DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	var dpcURL = "/six910/rs/productCategory/delete"
	var dpcc jv.Claim
	dpcc.Role = storeAdmin
	dpcc.URL = dpcURL
	dpcc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dpcc)
	h.Log.Debug("product category delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var dpccatidStr = vars["categoryId"]
			var dpcprodidStr = vars["productId"]
			var dprodstoreIDStr = vars["storeId"]
			cid, dpccatiderr := strconv.ParseInt(dpccatidStr, 10, 64)
			pid, dpcprodiderr := strconv.ParseInt(dpcprodidStr, 10, 64)
			storeID, dpcsiderr := strconv.ParseInt(dprodstoreIDStr, 10, 64)
			var dpcres *m.Response
			if dpccatiderr == nil && dpcprodiderr == nil && dpcsiderr == nil {
				var pc sdbi.ProductCategory
				pc.CategoryID = cid
				pc.ProductID = pid
				dpcres = h.Manager.DeleteProductCategory(&pc, storeID)
				h.Log.Debug("dpcres: ", *dpcres)
				if dpcres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dpcres = &nc
			}
			resJSON, _ := json.Marshal(dpcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
