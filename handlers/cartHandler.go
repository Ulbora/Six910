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

// AddCart godoc
// @Summary Add a new cart
// @Description Adds a new cart to a store
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param cart body six910-database-interface.Cart true "cart"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/cart/add [post]
func (h *Six910Handler) AddCart(w http.ResponseWriter, r *http.Request) {
	var addCartURL = "/six910/rs/cart/add"
	var actc jv.Claim
	actc.Role = customerRole
	actc.URL = addCartURL
	actc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cart add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var crt sdbi.Cart
			actsuc, acterr := h.ProcessBody(r, &crt)
			h.Log.Debug("actsuc: ", actsuc)
			h.Log.Debug("crt: ", crt)
			h.Log.Debug("acterr: ", acterr)
			if !actsuc && acterr != nil {
				http.Error(w, acterr.Error(), http.StatusBadRequest)
			} else {
				actres := h.Manager.AddCart(&crt)
				h.Log.Debug("actres: ", *actres)
				if actres.Success && actres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(actres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var actfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(actfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateCart godoc
// @Summary Update a cart
// @Description Update cart data
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param cart body six910-database-interface.Cart true "cart"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/cart/update [put]
func (h *Six910Handler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	var upCartURL = "/six910/rs/cart/update"
	var uctc jv.Claim
	uctc.Role = storeAdmin
	uctc.URL = upCartURL
	uctc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cart update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ucart sdbi.Cart
			uctsuc, ucterr := h.ProcessBody(r, &ucart)
			h.Log.Debug("uctsuc: ", uctsuc)
			h.Log.Debug("ucart: ", ucart)
			h.Log.Debug("ucterr: ", ucterr)
			if !uctsuc && ucterr != nil {
				http.Error(w, ucterr.Error(), http.StatusBadRequest)
			} else {
				uctres := h.Manager.UpdateCart(&ucart)
				h.Log.Debug("uctres: ", *uctres)
				if uctres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uctres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uctfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uctfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetCart godoc
// @Summary Get details of a cart by id
// @Description Get details of a cart
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param cid path string true "customer id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Cart
// @Router /rs/cart/get/{cid}/{storeId} [get]
func (h *Six910Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	var gcartURL = "/six910/rs/cart/get"
	var gctc2 jv.Claim
	gctc2.Role = customerRole
	gctc2.URL = gcartURL
	gctc2.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cart get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gccidStr = vars["cid"]
			var storeIDStr = vars["storeId"]
			cid, cerr := strconv.ParseInt(gccidStr, 10, 64)
			gcstoreID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gctres *sdbi.Cart
			if cerr == nil && serr == nil {
				gctres = h.Manager.GetCart(cid, gcstoreID)
				h.Log.Debug("gctres: ", gctres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Cart
				gctres = &nc
			}
			resJSON, _ := json.Marshal(gctres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteCart godoc
// @Summary Delete a cart
// @Description Delete a cart from the store
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param id path string true "cart id"
// @Param cid path string true "customer id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/cart/delete/{id}/{cid}/{storeId} [delete]
func (h *Six910Handler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	var dcartURL = "/six910/rs/cart/delete"
	var dcts jv.Claim
	dcts.Role = storeAdmin
	dcts.URL = dcartURL
	dcts.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cart delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var dcidStr = vars["id"]
			var dccidStr = vars["cid"]
			var dcstoreIDStr = vars["storeId"]
			id, iderr := strconv.ParseInt(dcidStr, 10, 64)
			cid, cerr := strconv.ParseInt(dccidStr, 10, 64)
			storeID, serr := strconv.ParseInt(dcstoreIDStr, 10, 64)
			var dctres *m.Response
			if iderr == nil && cerr == nil && serr == nil {
				dctres = h.Manager.DeleteCart(id, cid, storeID)
				h.Log.Debug("dctres: ", *dctres)
				if dctres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dctres = &nc
			}
			resJSON, _ := json.Marshal(dctres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
