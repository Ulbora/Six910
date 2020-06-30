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

//AddProduct AddProduct
func (h *Six910Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var addprodURL = "/six910/rs/product/add"
	var aprodc jv.Claim
	aprodc.Role = storeAdmin
	aprodc.URL = addprodURL
	aprodc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aprodc)
	h.Log.Debug("product add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aprod sdbi.Product
			aprodsuc, aproderr := h.ProcessBody(r, &aprod)
			h.Log.Debug("aprodsuc: ", aprodsuc)
			h.Log.Debug("aprod: ", aprod)
			h.Log.Debug("aproderr: ", aproderr)
			if !aprodsuc && aproderr != nil {
				http.Error(w, aproderr.Error(), http.StatusBadRequest)
			} else {
				aprodres := h.Manager.AddProduct(&aprod)
				h.Log.Debug("aprodres: ", *aprodres)
				if aprodres.Success && aprodres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aprodres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aprodfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aprodfl)
		fmt.Fprint(w, string(resJSON))
	}
}
