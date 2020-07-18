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

//OrderCommentReq OrderCommentReq
type OrderCommentReq struct {
	StoreID      int64             `json:"storeId"`
	OrderComment sdbi.OrderComment `json:"orderComment"`
}

// AddOrderComments godoc
// @Summary Add a new orderComment
// @Description Adds a new orderComment to a store
// @Tags OrderComment
// @Accept  json
// @Produce  json
// @Param orderComment body OrderCommentReq true "orderComment"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/orderComment/add [post]
func (h *Six910Handler) AddOrderComments(w http.ResponseWriter, r *http.Request) {
	var addorcURL = "/six910/rs/orderComment/add"
	var aorcc jv.Claim
	aorcc.Role = customerRole
	aorcc.URL = addorcURL
	aorcc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aorcc)
	h.Log.Debug("order comment add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aorc OrderCommentReq
			aorcsuc, aorcerr := h.ProcessBody(r, &aorc)
			h.Log.Debug("aorcsuc: ", aorcsuc)
			h.Log.Debug("aorc: ", aorc)
			h.Log.Debug("aorcerr: ", aorcerr)
			if !aorcsuc && aorcerr != nil {
				http.Error(w, aorcerr.Error(), http.StatusBadRequest)
			} else {
				aorcres := h.Manager.AddOrderComments(&aorc.OrderComment, aorc.StoreID)
				h.Log.Debug("aorcres: ", *aorcres)
				if aorcres.Success && aorcres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aorcres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aorcf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aorcf)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetOrderCommentList godoc
// @Summary Get list of orderComment
// @Description Get list of orderComment for a store
// @Tags OrderComment
// @Accept  json
// @Produce  json
// @Param orderId path string true "order Id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.OrderComment
// @Router /rs/orderComment/get/list/{orderId}/{storeId} [get]
func (h *Six910Handler) GetOrderCommentList(w http.ResponseWriter, r *http.Request) {
	var gorclURL = "/six910/rs/orderComment/list"
	var gorccl jv.Claim
	gorccl.Role = customerRole
	gorccl.URL = gorclURL
	gorccl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gorccl)
	h.Log.Debug("order comment get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var orclordIDStr = vars["orderId"]
			var orclstoreIDStr = vars["storeId"]
			orderID, sorcloiderr := strconv.ParseInt(orclordIDStr, 10, 64)
			storeID, sorclserr := strconv.ParseInt(orclstoreIDStr, 10, 64)
			var gorclres *[]sdbi.OrderComment
			if sorcloiderr == nil && sorclserr == nil {
				gorclres = h.Manager.GetOrderCommentList(orderID, storeID)
				h.Log.Debug("get order comment list: ", gorclres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.OrderComment{}
				gorclres = &nc
			}
			resJSON, _ := json.Marshal(gorclres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
