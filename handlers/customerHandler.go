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

//AddCustomer AddCustomer
func (h *Six910Handler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var addCusURL = "/six910/rs/customer/add"
	var acc jv.Claim
	acc.Role = customerRole
	acc.URL = addCusURL
	acc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.ValidatorClient.Authorize(r, &acc, h.ValidationURL)
	h.Log.Debug("cus add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cust sdbi.Customer
			scsuc, scerr := h.ProcessBody(r, &cust)
			h.Log.Debug("scsuc: ", scsuc)
			h.Log.Debug("cust: ", cust)
			h.Log.Debug("scerr: ", scerr)
			if !scsuc && scerr != nil {
				http.Error(w, scerr.Error(), http.StatusBadRequest)
			} else {
				acres := h.Manager.AddCustomer(&cust)
				h.Log.Debug("acres: ", *acres)
				if acres.Success && acres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(acres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var acfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(acfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateCustomer UpdateCustomer
func (h *Six910Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var upCusURL = "/six910/rs/customer/update"
	var ucc jv.Claim
	ucc.Role = customerRole
	ucc.URL = upCusURL
	ucc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ucc)
	//auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("customer update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ucus sdbi.Customer
			ucsuc, ucerr := h.ProcessBody(r, &ucus)
			h.Log.Debug("ucsuc: ", ucsuc)
			h.Log.Debug("ucus: ", ucus)
			h.Log.Debug("ucerr: ", ucerr)
			if !ucsuc && ucerr != nil {
				http.Error(w, ucerr.Error(), http.StatusBadRequest)
			} else {
				ucres := h.Manager.UpdateCustomer(&ucus)
				h.Log.Debug("ucres: ", *ucres)
				if ucres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ucres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ucfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ucfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetCustomer GetCustomer
func (h *Six910Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	var gCusURL = "/six910/rs/customer/get"
	var gcc jv.Claim
	gcc.Role = customerRole
	gcc.URL = gCusURL
	gcc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gcc)
	//auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("cus get authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var storeName = vars["email"]
			var storeIDStr = vars["storeId"]
			storeID, err := strconv.ParseInt(storeIDStr, 10, 64)
			var gcres *sdbi.Customer
			if err == nil {
				gcres = h.Manager.GetCustomer(storeName, storeID)
				h.Log.Debug("getCust: ", gcres)
				w.WriteHeader(http.StatusOK)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
				var nc sdbi.Customer
				gcres = &nc
			}
			resJSON, _ := json.Marshal(gcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetCustomerID GetCustomerID
func (h *Six910Handler) GetCustomerID(w http.ResponseWriter, r *http.Request) {
	var gCus2URL = "/six910/rs/customer/get/id"
	var gcc2 jv.Claim
	gcc2.Role = customerRole
	gcc2.URL = gCus2URL
	gcc2.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gcc2)
	//auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("cus get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var cidStr = vars["id"]
			var storeIDStr = vars["storeId"]
			cid, cerr := strconv.ParseInt(cidStr, 10, 64)
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gcres *sdbi.Customer
			if cerr == nil && serr == nil {
				gcres = h.Manager.GetCustomerID(cid, storeID)
				h.Log.Debug("getCustId: ", gcres)
				w.WriteHeader(http.StatusOK)
			} else {
				//http.Error(w, err.Error(), http.StatusBadRequest)
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Customer
				gcres = &nc
			}
			resJSON, _ := json.Marshal(gcres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
