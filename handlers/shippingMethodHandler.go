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

// AddShippingMethod godoc
// @Summary Add a new shipping method
// @Description Adds a new shipping method to a store
// @Tags ShippingMethod
// @Accept  json
// @Produce  json
// @Param shippingMethod body six910-database-interface.ShippingMethod true "shippingMethod"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/shippingMethod/add [post]
func (h *Six910Handler) AddShippingMethod(w http.ResponseWriter, r *http.Request) {
	var addsmURL = "/six910/rs/shippingMethod/add"
	var asmc jv.Claim
	asmc.Role = storeAdmin
	asmc.URL = addsmURL
	asmc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &asmc)
	h.Log.Debug("shipping method add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var asm sdbi.ShippingMethod
			asmsuc, asmerr := h.ProcessBody(r, &asm)
			h.Log.Debug("asmsuc: ", asmsuc)
			h.Log.Debug("asm: ", asm)
			h.Log.Debug("asmerr: ", asmerr)
			if !asmsuc && asmerr != nil {
				http.Error(w, asmerr.Error(), http.StatusBadRequest)
			} else {
				asmres := h.Manager.AddShippingMethod(&asm)
				h.Log.Debug("asmres: ", *asmres)
				if asmres.Success && asmres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(asmres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var asmfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(asmfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateShippingMethod godoc
// @Summary Update a shipping method
// @Description Update shipping method data
// @Tags ShippingMethod
// @Accept  json
// @Produce  json
// @Param shippingMethod body six910-database-interface.ShippingMethod true "shippingMethod"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shippingMethod/update [put]
func (h *Six910Handler) UpdateShippingMethod(w http.ResponseWriter, r *http.Request) {
	var upsmURL = "/six910/rs/shippingMethod/update"
	var usmc jv.Claim
	usmc.Role = storeAdmin
	usmc.URL = upsmURL
	usmc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &usmc)
	h.Log.Debug("shipping method update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var usm sdbi.ShippingMethod
			usmsuc, usmerr := h.ProcessBody(r, &usm)
			h.Log.Debug("usmsuc: ", usmsuc)
			h.Log.Debug("usm: ", usm)
			h.Log.Debug("usmerr: ", usmerr)
			if !usmsuc && usmerr != nil {
				http.Error(w, usmerr.Error(), http.StatusBadRequest)
			} else {
				usmres := h.Manager.UpdateShippingMethod(&usm)
				h.Log.Debug("usmres: ", *usmres)
				if usmres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(usmres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var usmfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(usmfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetShippingMethod godoc
// @Summary Get details of a shipping method by id
// @Description Get details of a shipping method
// @Tags ShippingMethod
// @Accept  json
// @Produce  json
// @Param id path string true "shippingMethod id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.ShippingMethod
// @Router /rs/shippingMethod/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetShippingMethod(w http.ResponseWriter, r *http.Request) {
	var gsmURL = "/six910/rs/shippingMethod/get"
	var gsmc jv.Claim
	gsmc.Role = customerRole
	gsmc.URL = gsmURL
	gsmc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("shipping method get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gsmidStr = vars["id"]
			var gsmstoreIDStr = vars["storeId"]
			id, gsmiderr := strconv.ParseInt(gsmidStr, 10, 64)
			storeID, gsmsiderr := strconv.ParseInt(gsmstoreIDStr, 10, 64)
			var gsmres *sdbi.ShippingMethod
			if gsmiderr == nil && gsmsiderr == nil {
				gsmres = h.Manager.GetShippingMethod(id, storeID)
				h.Log.Debug("gsmres: ", gsmres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.ShippingMethod
				gsmres = &nc
			}
			resJSON, _ := json.Marshal(gsmres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetShippingMethodList godoc
// @Summary Get list of shipping method
// @Description Get list of shipping method for a store
// @Tags ShippingMethod
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.ShippingMethod
// @Router /rs/shippingMethod/get/list/{storeId} [get]
func (h *Six910Handler) GetShippingMethodList(w http.ResponseWriter, r *http.Request) {
	var gsmlURL = "/six910/rs/shippingMethod/list"
	var gsmcl jv.Claim
	gsmcl.Role = customerRole
	gsmcl.URL = gsmlURL
	gsmcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("shipping method get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var gsmlstoreIDStr = vars["storeId"]
			storeID, gsmlerr := strconv.ParseInt(gsmlstoreIDStr, 10, 64)
			var gsmlres *[]sdbi.ShippingMethod
			if gsmlerr == nil {
				gsmlres = h.Manager.GetShippingMethodList(storeID)
				h.Log.Debug("get shipping method list: ", gsmlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.ShippingMethod{}
				gsmlres = &nc
			}
			resJSON, _ := json.Marshal(gsmlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteShippingMethod godoc
// @Summary Delete a  shipping method
// @Description Delete a  shipping method from the store
// @Tags ShippingMethod
// @Accept  json
// @Produce  json
// @Param id path string true "shippingMethod id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shippingMethod/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteShippingMethod(w http.ResponseWriter, r *http.Request) {
	var dsmURL = "/six910/rs/shippingMethod/delete"
	var dsms jv.Claim
	dsms.Role = storeAdmin
	dsms.URL = dsmURL
	dsms.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dsms)
	h.Log.Debug("shipping method delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dsmidStr = vars["id"]
			var dsmstoreIDStr = vars["storeId"]
			id, dsmiderr := strconv.ParseInt(dsmidStr, 10, 64)
			storeID, dsmidserr := strconv.ParseInt(dsmstoreIDStr, 10, 64)
			var dsmres *m.Response
			if dsmiderr == nil && dsmidserr == nil {
				dsmres = h.Manager.DeleteShippingMethod(id, storeID)
				h.Log.Debug("deleteSm: ", dsmres)
				if dsmres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dsmres = &nc
			}
			resJSON, _ := json.Marshal(dsmres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
