package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	m "github.com/Ulbora/Six910/managers"
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

//AddCustomer AddCustomer
func (h *Six910Handler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	var addCusURL = "/six910/rs/customer/add"
	var acc jv.Claim
	acc.Role = storeAdmin
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
	ucc.Role = storeAdmin
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