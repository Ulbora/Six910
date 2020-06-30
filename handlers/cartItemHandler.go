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

//CartItemReq CartItemReq
type CartItemReq struct {
	CustomerID int64         `json:"customerId"`
	StoreID    int64         `json:"storeId"`
	CartItem   sdbi.CartItem `json:"cartItem"`
}

//AddCartItem AddCartItem
func (h *Six910Handler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	var addciURL = "/six910/rs/cartItem/add"
	var acic jv.Claim
	acic.Role = customerRole
	acic.URL = addciURL
	acic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &acic)
	h.Log.Debug("cart item add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var cireq CartItemReq
			acirsuc, acierr := h.ProcessBody(r, &cireq)
			h.Log.Debug("acirsuc: ", acirsuc)
			h.Log.Debug("cireq: ", cireq)
			h.Log.Debug("acierr: ", acierr)
			if !acirsuc && acierr != nil {
				http.Error(w, acierr.Error(), http.StatusBadRequest)
			} else {
				acires := h.Manager.AddCartItem(&cireq.CartItem, cireq.CustomerID, cireq.StoreID)
				h.Log.Debug("acires: ", *acires)
				if acires.Success && acires.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(acires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var acifl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(acifl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateCartItem UpdateCartItem
func (h *Six910Handler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	var upciURL = "/six910/rs/cartItem/update"
	var ucic jv.Claim
	ucic.Role = customerRole
	ucic.URL = upciURL
	ucic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ucic)
	h.Log.Debug("cart item update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ucireq CartItemReq
			ucisuc, ucierr := h.ProcessBody(r, &ucireq)
			h.Log.Debug("ucisuc: ", ucisuc)
			h.Log.Debug("ucireq: ", ucireq)
			h.Log.Debug("ucierr: ", ucierr)
			if !ucisuc && ucierr != nil {
				http.Error(w, ucierr.Error(), http.StatusBadRequest)
			} else {
				ucires := h.Manager.UpdateCartItem(&ucireq.CartItem, ucireq.CustomerID, ucireq.StoreID)
				h.Log.Debug("ucires: ", *ucires)
				if ucires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ucires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ucifl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ucifl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetCarItem GetCarItem
func (h *Six910Handler) GetCarItem(w http.ResponseWriter, r *http.Request) {
	var gCiURL = "/six910/rs/cartItem/get"
	var gci2 jv.Claim
	gci2.Role = customerRole
	gci2.URL = gCiURL
	gci2.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gci2)
	h.Log.Debug("dist get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var cartIDStr = vars["cartId"]
			var prodIDStr = vars["prodId"]
			var storeIDStr = vars["storeId"]
			cartID, cartIDerr := strconv.ParseInt(cartIDStr, 10, 64)
			prodID, prodIDerr := strconv.ParseInt(prodIDStr, 10, 64)
			storeID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gcires *sdbi.CartItem
			if cartIDerr == nil && prodIDerr == nil && serr == nil {
				gcires = h.Manager.GetCarItem(cartID, prodID, storeID)
				h.Log.Debug("gcires: ", *gcires)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.CartItem
				gcires = &nc
			}
			resJSON, _ := json.Marshal(gcires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetCartItemList GetCartItemList
func (h *Six910Handler) GetCartItemList(w http.ResponseWriter, r *http.Request) {
	var gCilURL = "/six910/rs/cartItem/list"
	var gcil2 jv.Claim
	gcil2.Role = customerRole
	gcil2.URL = gCilURL
	gcil2.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gcil2)
	h.Log.Debug("dist get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var cartIDlStr = vars["cartId"]
			var cidlStr = vars["cid"]
			var storeIDlStr = vars["storeId"]
			cartIDl, cartIDerrl := strconv.ParseInt(cartIDlStr, 10, 64)
			cIDl, cIDerrl := strconv.ParseInt(cidlStr, 10, 64)
			storeIDl, serrl := strconv.ParseInt(storeIDlStr, 10, 64)
			var gcilres *[]sdbi.CartItem
			if cartIDerrl == nil && cIDerrl == nil && serrl == nil {
				gcilres = h.Manager.GetCartItemList(cartIDl, cIDl, storeIDl)
				h.Log.Debug("gcilres: ", *gcilres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var ncl = []sdbi.CartItem{}
				gcilres = &ncl
			}
			resJSON, _ := json.Marshal(gcilres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteCartItem DeleteCartItem
func (h *Six910Handler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	var gCiddURL = "/six910/rs/cartItem/delete"
	var gci2d jv.Claim
	gci2d.Role = customerRole
	gci2d.URL = gCiddURL
	gci2d.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gci2d)
	h.Log.Debug("dist delete id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var IDStrd = vars["id"]
			var prodIDStrd = vars["prodId"]
			var cartIDStrd = vars["cartId"]
			IDd, serrd := strconv.ParseInt(IDStrd, 10, 64)
			prodIDd, prodIDerrd := strconv.ParseInt(prodIDStrd, 10, 64)
			cartIDd, cartIDerrd := strconv.ParseInt(cartIDStrd, 10, 64)
			var gciresd *m.Response
			if cartIDerrd == nil && prodIDerrd == nil && serrd == nil {
				gciresd = h.Manager.DeleteCartItem(IDd, prodIDd, cartIDd)
				if gciresd.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			resJSON, _ := json.Marshal(gciresd)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
