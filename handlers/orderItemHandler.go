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

//OrderItemReq OrderItemReq
type OrderItemReq struct {
	StoreID   int64          `json:"storeId"`
	OrderItem sdbi.OrderItem `json:"orderItem"`
}

// AddOrderItem godoc
// @Summary Add a new orderItem
// @Description Adds a new orderItem to a store
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Param orderItem body OrderItemReq true "orderItem"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/orderItem/add [post]
func (h *Six910Handler) AddOrderItem(w http.ResponseWriter, r *http.Request) {
	var addoriURL = "/six910/rs/orderItem/add"
	var aoric jv.Claim
	aoric.Role = customerRole
	aoric.URL = addoriURL
	aoric.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aoric)
	h.Log.Debug("order item add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aori OrderItemReq
			aorisuc, aorierr := h.ProcessBody(r, &aori)
			h.Log.Debug("aorisuc: ", aorisuc)
			h.Log.Debug("aori: ", aori)
			h.Log.Debug("aorerr: ", aorierr)
			if !aorisuc && aorierr != nil {
				http.Error(w, aorierr.Error(), http.StatusBadRequest)
			} else {
				aorires := h.Manager.AddOrderItem(&aori.OrderItem, aori.StoreID)
				h.Log.Debug("aorires: ", *aorires)
				if aorires.Success && aorires.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aorires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aorif m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aorif)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateOrderItem godoc
// @Summary Update a orderItem
// @Description Update orderItem data
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Param orderItem body OrderItemReq true "orderItem"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/orderItem/update [put]
func (h *Six910Handler) UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	var uoriURL = "/six910/rs/orderItem/update"
	var uoric jv.Claim
	uoric.Role = customerRole
	uoric.URL = uoriURL
	uoric.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uoric)
	h.Log.Debug("order item update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uori OrderItemReq
			uorisuc, uorierr := h.ProcessBody(r, &uori)
			h.Log.Debug("uorisuc: ", uorisuc)
			h.Log.Debug("uori: ", uori)
			h.Log.Debug("uorierr: ", uorierr)
			if !uorisuc && uorierr != nil {
				http.Error(w, uorierr.Error(), http.StatusBadRequest)
			} else {
				uorires := h.Manager.UpdateOrderItem(&uori.OrderItem, uori.StoreID)
				h.Log.Debug("uorires: ", *uorires)
				if uorires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uorires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uorif m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uorif)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetOrderItem godoc
// @Summary Get details of a orderItem by id
// @Description Get details of a orderItem
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Param id path string true "orderItem id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.OrderItem
// @Router /rs/orderItem/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetOrderItem(w http.ResponseWriter, r *http.Request) {
	var goriURL = "/six910/rs/orderItem/get"
	var goric jv.Claim
	goric.Role = customerRole
	goric.URL = goriURL
	goric.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &goric)
	h.Log.Debug("order item get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var goriidStr = vars["id"]
			var goristoreIDStr = vars["storeId"]
			id, goriiderr := strconv.ParseInt(goriidStr, 10, 64)
			storeID, gorisiderr := strconv.ParseInt(goristoreIDStr, 10, 64)
			var gorires *sdbi.OrderItem
			if goriiderr == nil && gorisiderr == nil {
				gorires = h.Manager.GetOrderItem(id, storeID)
				h.Log.Debug("gorires: ", gorires)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.OrderItem
				gorires = &nc
			}
			resJSON, _ := json.Marshal(gorires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetOrderItemList godoc
// @Summary Get list of orderItem
// @Description Get list of orderItem for a store
// @Tags OrderItem
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
// @Success 200 {array} six910-database-interface.OrderItem
// @Router /rs/orderItem/get/list/{orderId}/{storeId} [get]
func (h *Six910Handler) GetOrderItemList(w http.ResponseWriter, r *http.Request) {
	var gorilURL = "/six910/rs/orderItem/list"
	var goricl jv.Claim
	goricl.Role = customerRole
	goricl.URL = gorilURL
	goricl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &goricl)
	h.Log.Debug("order item get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var orlordIDStr = vars["orderId"]
			var orlstoreIDStr = vars["storeId"]
			orderID, soriloiderr := strconv.ParseInt(orlordIDStr, 10, 64)
			storeID, sorilserr := strconv.ParseInt(orlstoreIDStr, 10, 64)
			var gorilres *[]sdbi.OrderItem
			if soriloiderr == nil && sorilserr == nil {
				gorilres = h.Manager.GetOrderItemList(orderID, storeID)
				h.Log.Debug("get order item list: ", gorilres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.OrderItem{}
				gorilres = &nc
			}
			resJSON, _ := json.Marshal(gorilres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteOrderItem godoc
// @Summary Delete a orderItem
// @Description Delete a orderItem from the store
// @Tags OrderItem
// @Accept  json
// @Produce  json
// @Param id path string true "orderItem id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/orderItem/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	var doriURL = "/six910/rs/orderItem/delete"
	var doric jv.Claim
	doric.Role = customerRole
	doric.URL = doriURL
	doric.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &doric)
	h.Log.Debug("order item delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var doriidStr = vars["id"]
			var doristoreIDStr = vars["storeId"]
			id, doriiderr := strconv.ParseInt(doriidStr, 10, 64)
			storeID, doriidserr := strconv.ParseInt(doristoreIDStr, 10, 64)
			var dorires *m.Response
			if doriiderr == nil && doriidserr == nil {
				dorires = h.Manager.DeleteOrderItem(id, storeID)
				h.Log.Debug("dorres: ", *dorires)
				if dorires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dorires = &nc
			}
			resJSON, _ := json.Marshal(dorires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
