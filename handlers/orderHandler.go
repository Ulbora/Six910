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

//AddOrder AddOrder
func (h *Six910Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var addorURL = "/six910/rs/order/add"
	var aorc jv.Claim
	aorc.Role = customerRole
	aorc.URL = addorURL
	aorc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aorc)
	h.Log.Debug("order add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aor sdbi.Order
			aorsuc, aorerr := h.ProcessBody(r, &aor)
			h.Log.Debug("aorsuc: ", aorsuc)
			h.Log.Debug("aor: ", aor)
			h.Log.Debug("aorerr: ", aorerr)
			if !aorsuc && aorerr != nil {
				http.Error(w, aorerr.Error(), http.StatusBadRequest)
			} else {
				aorres := h.Manager.AddOrder(&aor)
				h.Log.Debug("aorres: ", *aorres)
				if aorres.Success && aorres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aorres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aorf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aorf)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateOrder UpdateOrder
func (h *Six910Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var uorURL = "/six910/rs/order/update"
	var uorc jv.Claim
	uorc.Role = customerRole
	uorc.URL = uorURL
	uorc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uorc)
	h.Log.Debug("order update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uor sdbi.Order
			uorsuc, uorerr := h.ProcessBody(r, &uor)
			h.Log.Debug("uorsuc: ", uorsuc)
			h.Log.Debug("uor: ", uor)
			h.Log.Debug("uorerr: ", uorerr)
			if !uorsuc && uorerr != nil {
				http.Error(w, uorerr.Error(), http.StatusBadRequest)
			} else {
				uorres := h.Manager.UpdateOrder(&uor)
				h.Log.Debug("uorres: ", *uorres)
				if uorres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uorres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uorf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uorf)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetOrder GetOrder
func (h *Six910Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var gorURL = "/six910/rs/order/get"
	var gorc jv.Claim
	gorc.Role = customerRole
	gorc.URL = gorURL
	gorc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gorc)
	h.Log.Debug("order get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var goridStr = vars["id"]
			var gorstoreIDStr = vars["storeId"]
			id, goriderr := strconv.ParseInt(goridStr, 10, 64)
			storeID, gorsiderr := strconv.ParseInt(gorstoreIDStr, 10, 64)
			var gorres *sdbi.Order
			if goriderr == nil && gorsiderr == nil {
				gorres = h.Manager.GetOrder(id, storeID)
				h.Log.Debug("gorres: ", gorres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Order
				gorres = &nc
			}
			resJSON, _ := json.Marshal(gorres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetOrderList GetOrderList
func (h *Six910Handler) GetOrderList(w http.ResponseWriter, r *http.Request) {
	var gorlURL = "/six910/rs/order/list"
	var gorcl jv.Claim
	gorcl.Role = customerRole
	gorcl.URL = gorlURL
	gorcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gorcl)
	h.Log.Debug("order get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var orlcidStr = vars["cid"]
			var orlstoreIDStr = vars["storeId"]
			cID, sorlciderr := strconv.ParseInt(orlcidStr, 10, 64)
			storeID, sorlserr := strconv.ParseInt(orlstoreIDStr, 10, 64)
			var gorlres *[]sdbi.Order
			if sorlciderr == nil && sorlserr == nil {
				gorlres = h.Manager.GetOrderList(cID, storeID)
				h.Log.Debug("get order list: ", gorlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Order{}
				gorlres = &nc
			}
			resJSON, _ := json.Marshal(gorlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteOrder DeleteOrder
func (h *Six910Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var dorURL = "/six910/rs/order/delete"
	var dorc jv.Claim
	dorc.Role = customerRole
	dorc.URL = dorURL
	dorc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dorc)
	h.Log.Debug("order delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var doridStr = vars["id"]
			var dorstoreIDStr = vars["storeId"]
			id, doriderr := strconv.ParseInt(doridStr, 10, 64)
			storeID, doridserr := strconv.ParseInt(dorstoreIDStr, 10, 64)
			var dorres *m.Response
			if doriderr == nil && doridserr == nil {
				dorres = h.Manager.DeleteOrder(id, storeID)
				h.Log.Debug("dorres: ", *dorres)
				if dorres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dorres = &nc
			}
			resJSON, _ := json.Marshal(dorres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
