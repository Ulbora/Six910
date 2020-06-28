package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

//AddStore AddStore
func (h *Six910Handler) AddStore(w http.ResponseWriter, r *http.Request) {
	var addStoreURL = "/six910/rs/store/add"
	var c jv.Claim
	c.Role = superAdmin
	c.URL = addStoreURL
	c.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("store add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		asOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", asOk)
		if !asOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var str sdbi.Store
			sasuc, saerr := h.ProcessBody(r, &str)
			h.Log.Debug("sasuc: ", sasuc)
			h.Log.Debug("str: ", str)
			h.Log.Debug("saerr: ", saerr)
			if !sasuc && saerr != nil {
				http.Error(w, saerr.Error(), http.StatusBadRequest)
			} else {
				asres := h.Manager.AddStore(&str)
				h.Log.Debug("asres: ", *asres)
				if asres.Success && asres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(asres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var asfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(asfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateStore UpdateStore
func (h *Six910Handler) UpdateStore(w http.ResponseWriter, r *http.Request) {
	var upStoreURL = "/six910/rs/store/update"
	var c jv.Claim
	c.Role = storeAdmin
	c.URL = upStoreURL
	c.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &c)
	//auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("store update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		usOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", usOk)
		if !usOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ustr sdbi.Store
			uasuc, uaerr := h.ProcessBody(r, &ustr)
			h.Log.Debug("uasuc: ", uasuc)
			h.Log.Debug("ustr: ", ustr)
			h.Log.Debug("uaerr: ", uaerr)
			if !uasuc && uaerr != nil {
				http.Error(w, uaerr.Error(), http.StatusBadRequest)
			} else {
				usres := h.Manager.UpdateStore(&ustr)
				h.Log.Debug("usres: ", *usres)
				if usres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(usres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var usfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(usfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetStore GetStore
func (h *Six910Handler) GetStore(w http.ResponseWriter, r *http.Request) {
	var gStoreURL = "/six910/rs/store/get"
	var gsc jv.Claim
	gsc.Role = storeAdmin
	gsc.URL = gStoreURL
	gsc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gsc)
	//auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("store get authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var storeName = vars["storeName"]
			var localDomain = vars["localDomain"]
			gsres := h.Manager.GetStore(storeName, localDomain)
			h.Log.Debug("getStore: ", gsres)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(gsres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteStore DeleteStore
func (h *Six910Handler) DeleteStore(w http.ResponseWriter, r *http.Request) {
	var dStoreURL = "/six910/rs/store/delete"
	var dsc jv.Claim
	dsc.Role = storeAdmin
	dsc.URL = dStoreURL
	dsc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dsc)
	//auth := h.ValidatorClient.Authorize(r, &c, h.ValidationURL)
	h.Log.Debug("store delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var storeName = vars["storeName"]
			var localDomain = vars["localDomain"]
			dsres := h.Manager.DeleteStore(storeName, localDomain)
			h.Log.Debug("del Store: ", dsres)
			if dsres.Success {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			resJSON, _ := json.Marshal(dsres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
