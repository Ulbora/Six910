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

//AddressReq AddressReq
type AddressReq struct {
	StoreID int64        `json:"storeId"`
	Address sdbi.Address `json:"address"`
}

//AddAddress AddAddress
func (h *Six910Handler) AddAddress(w http.ResponseWriter, r *http.Request) {
	var addadURL = "/six910/rs/address/add"
	var aadc jv.Claim
	aadc.Role = customerRole
	aadc.URL = addadURL
	aadc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aadc)
	h.Log.Debug("address add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var addReq AddressReq
			aadsuc, aaderr := h.ProcessBody(r, &addReq)
			h.Log.Debug("aadsuc: ", aadsuc)
			h.Log.Debug("addReq: ", addReq)
			h.Log.Debug("aaderr: ", aaderr)
			if !aadsuc && aaderr != nil {
				http.Error(w, aaderr.Error(), http.StatusBadRequest)
			} else {
				aadres := h.Manager.AddAddress(&addReq.Address, addReq.StoreID)
				h.Log.Debug("aadres: ", *aadres)
				if aadres.Success && aadres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aadres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aadfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aadfl)
		fmt.Fprint(w, string(resJSON))
	}
}
