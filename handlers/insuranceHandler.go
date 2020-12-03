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

// AddInsurance godoc
// @Summary Add shipping insurance provider
// @Description Adds shipping insurance provider to a store
// @Tags Insurance (Shipping Insurance)
// @Accept  json
// @Produce  json
// @Param insurance body six910-database-interface.Insurance true "insurance"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/insurance/add [post]
func (h *Six910Handler) AddInsurance(w http.ResponseWriter, r *http.Request) {
	var addinsURL = "/six910/rs/insurance/add"
	var ainsc jv.Claim
	ainsc.Role = storeAdmin
	ainsc.URL = addinsURL
	ainsc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ainsc)
	h.Log.Debug("insurance add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ains sdbi.Insurance
			ainssuc, ainserr := h.ProcessBody(r, &ains)
			h.Log.Debug("ainssuc: ", ainssuc)
			h.Log.Debug("ains: ", ains)
			h.Log.Debug("ainserr: ", ainserr)
			if !ainssuc && ainserr != nil {
				http.Error(w, ainserr.Error(), http.StatusBadRequest)
			} else {
				ainsres := h.Manager.AddInsurance(&ains)
				h.Log.Debug("ainsres: ", *ainsres)
				if ainsres.Success && ainsres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ainsres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ainsfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ainsfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateInsurance godoc
// @Summary Update shipping insurance provider
// @Description Update shipping insurance provider data
// @Tags Insurance (Shipping Insurance)
// @Accept  json
// @Produce  json
// @Param insurance body six910-database-interface.Insurance true "insurance"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/insurance/update [put]
func (h *Six910Handler) UpdateInsurance(w http.ResponseWriter, r *http.Request) {
	var upinsURL = "/six910/rs/insurance/update"
	var uinsc jv.Claim
	uinsc.Role = storeAdmin
	uinsc.URL = upinsURL
	uinsc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uinsc)
	h.Log.Debug("insurance update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uins sdbi.Insurance
			uinssuc, uinserr := h.ProcessBody(r, &uins)
			h.Log.Debug("uinssuc: ", uinssuc)
			h.Log.Debug("uins: ", uins)
			h.Log.Debug("uinserr: ", uinserr)
			if !uinssuc && uinserr != nil {
				http.Error(w, uinserr.Error(), http.StatusBadRequest)
			} else {
				uinsres := h.Manager.UpdateInsurance(&uins)
				h.Log.Debug("uinsres: ", *uinsres)
				if uinsres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uinsres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uinsfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uinsfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetInsurance godoc
// @Summary Get details of a shipping insurance provider by id
// @Description Get details of a shipping insurance provider
// @Tags Insurance (Shipping Insurance)
// @Accept  json
// @Produce  json
// @Param id path string true "insurance id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Insurance
// @Router /rs/insurance/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetInsurance(w http.ResponseWriter, r *http.Request) {
	var ginsURL = "/six910/rs/insurance/get"
	var ginsc jv.Claim
	ginsc.Role = customerRole
	ginsc.URL = ginsURL
	ginsc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("insurance get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var ginsidStr = vars["id"]
			var ginsstoreIDStr = vars["storeId"]
			id, ginsiderr := strconv.ParseInt(ginsidStr, 10, 64)
			storeID, ginssiderr := strconv.ParseInt(ginsstoreIDStr, 10, 64)
			var ginsres *sdbi.Insurance
			if ginsiderr == nil && ginssiderr == nil {
				ginsres = h.Manager.GetInsurance(id, storeID)
				h.Log.Debug("ginsres: ", ginsres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Insurance
				ginsres = &nc
			}
			resJSON, _ := json.Marshal(ginsres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetInsuranceList godoc
// @Summary Get list of shipping insurance providers
// @Description Get list of shipping insurance providers for a store
// @Tags Insurance (Shipping Insurance)
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Insurance
// @Router /rs/insurance/get/list/{storeId} [get]
func (h *Six910Handler) GetInsuranceList(w http.ResponseWriter, r *http.Request) {
	var ginslURL = "/six910/rs/insurance/list"
	var ginscl jv.Claim
	ginscl.Role = customerRole
	ginscl.URL = ginslURL
	ginscl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("insurance get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var inslstoreIDStr = vars["storeId"]
			storeID, sinslerr := strconv.ParseInt(inslstoreIDStr, 10, 64)
			var ginslres *[]sdbi.Insurance
			if sinslerr == nil {
				ginslres = h.Manager.GetInsuranceList(storeID)
				h.Log.Debug("get insurance list: ", ginslres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Insurance{}
				ginslres = &nc
			}
			resJSON, _ := json.Marshal(ginslres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteInsurance godoc
// @Summary Delete a shipping insurance provider
// @Description Delete a shipping insurance provider from the store
// @Tags Insurance (Shipping Insurance)
// @Accept  json
// @Produce  json
// @Param id path string true "insurance id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/insurance/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteInsurance(w http.ResponseWriter, r *http.Request) {
	var dinsURL = "/six910/rs/insurance/delete"
	var dinss jv.Claim
	dinss.Role = storeAdmin
	dinss.URL = dinsURL
	dinss.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dinss)
	h.Log.Debug("insurance delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dinsidStr = vars["id"]
			var dinsstoreIDStr = vars["storeId"]
			id, dinsiderr := strconv.ParseInt(dinsidStr, 10, 64)
			storeID, dinsidserr := strconv.ParseInt(dinsstoreIDStr, 10, 64)
			var dinsres *m.Response
			if dinsiderr == nil && dinsidserr == nil {
				dinsres = h.Manager.DeleteInsurance(id, storeID)
				h.Log.Debug("dinsres: ", dinsres)
				if dinsres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dinsres = &nc
			}
			resJSON, _ := json.Marshal(dinsres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
