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

//AddressReq AddressReq
type AddressReq struct {
	StoreID int64        `json:"storeId"`
	Address sdbi.Address `json:"address"`
}

// AddAddress godoc
// @Summary Add a new address
// @Description Adds a new address for a customer to a store
// @Tags Address
// @Accept  json
// @Produce  json
// @Param address body AddressReq true "address"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/address/add [post]
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

// UpdateAddress godoc
// @Summary Update a address
// @Description Update address data
// @Tags Address
// @Accept  json
// @Produce  json
// @Param address body AddressReq true "address"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/address/update [put]
func (h *Six910Handler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var upadURL = "/six910/rs/address/update"
	var uadc jv.Claim
	uadc.Role = customerRole
	uadc.URL = upadURL
	uadc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uadc)
	h.Log.Debug("address update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uadd AddressReq
			uadsuc, uaderr := h.ProcessBody(r, &uadd)
			h.Log.Debug("uadsuc: ", uadsuc)
			h.Log.Debug("uadd: ", uadd)
			h.Log.Debug("uaderr: ", uaderr)
			if !uadsuc && uaderr != nil {
				http.Error(w, uaderr.Error(), http.StatusBadRequest)
			} else {
				uadres := h.Manager.UpdateAddress(&uadd.Address, uadd.StoreID)
				h.Log.Debug("uadres: ", *uadres)
				if uadres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uadres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uadfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uadfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetAddress godoc
// @Summary Get details of a address by id
// @Description Get details of a address
// @Tags Address
// @Accept  json
// @Produce  json
// @Param id path string true "address id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Address
// @Router /rs/address/get/id/{id}/{cid}/{storeId} [get]
func (h *Six910Handler) GetAddress(w http.ResponseWriter, r *http.Request) {
	var gAdURL = "/six910/rs/address/get"
	var gadc jv.Claim
	gadc.Role = customerRole
	gadc.URL = gAdURL
	gadc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gadc)
	h.Log.Debug("address get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var gadidStr = vars["id"]
			var gadcidStr = vars["cid"]
			var storeIDStr = vars["storeId"]
			gaid, aiderr := strconv.ParseInt(gadidStr, 10, 64)
			gacid, aciderr := strconv.ParseInt(gadcidStr, 10, 64)
			gastoreID, serr := strconv.ParseInt(storeIDStr, 10, 64)
			var gadres *sdbi.Address
			if aiderr == nil && aciderr == nil && serr == nil {
				gadres = h.Manager.GetAddress(gaid, gacid, gastoreID)
				h.Log.Debug("gadres: ", gadres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Address
				gadres = &nc
			}
			resJSON, _ := json.Marshal(gadres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetAddressList godoc
// @Summary Get list of address
// @Description Get list of customer addresses for a store
// @Tags Address
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Address
// @Router /rs/address/get/list/{cid}/{storeId} [get]
func (h *Six910Handler) GetAddressList(w http.ResponseWriter, r *http.Request) {
	var gadlURL = "/six910/rs/address/list"
	var gadcl jv.Claim
	gadcl.Role = customerRole
	gadcl.URL = gadlURL
	gadcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gadcl)
	h.Log.Debug("address get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var cusadlIDStr = vars["cid"]
			var storeadlIDStr = vars["storeId"]
			cadlID, cadlerr := strconv.ParseInt(cusadlIDStr, 10, 64)
			storeadlID, sadlerr := strconv.ParseInt(storeadlIDStr, 10, 64)
			var gadlres *[]sdbi.Address
			if cadlerr == nil && sadlerr == nil {
				gadlres = h.Manager.GetAddressList(cadlID, storeadlID)
				h.Log.Debug("getAdd list: ", gadlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Address{}
				gadlres = &nc
			}
			resJSON, _ := json.Marshal(gadlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteAddress godoc
// @Summary Delete a address
// @Description Delete a customer address from the store
// @Tags Address
// @Accept  json
// @Produce  json
// @Param id path string true "address id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/address/delete/{id}/{cid}/{storeId} [delete]
func (h *Six910Handler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	var daddURL = "/six910/rs/address/delete"
	var dads jv.Claim
	dads.Role = customerRole
	dads.URL = daddURL
	dads.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dads)
	h.Log.Debug("address delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 3 {
			h.Log.Debug("vars: ", vars)
			var daidStr = vars["id"]
			var dacidStr = vars["cid"]
			var dastoreIDStr = vars["storeId"]
			daid, daderr := strconv.ParseInt(daidStr, 10, 64)
			dacid, dadcerr := strconv.ParseInt(dacidStr, 10, 64)
			storedaID, dadserr := strconv.ParseInt(dastoreIDStr, 10, 64)
			var dadres *m.Response
			if daderr == nil && dadcerr == nil && dadserr == nil {
				dadres = h.Manager.DeleteAddress(daid, dacid, storedaID)
				h.Log.Debug("deleteAdd: ", dadres)
				if dadres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dadres = &nc
			}
			resJSON, _ := json.Marshal(dadres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
