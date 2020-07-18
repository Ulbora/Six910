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

//ShipmentReq ShipmentReq
type ShipmentReq struct {
	StoreID  int64         `json:"storeId"`
	Shipment sdbi.Shipment `json:"shipment"`
}

// AddShipment godoc
// @Summary Add a new shipment
// @Description Adds a new shipment to a store
// @Tags Shipment
// @Accept  json
// @Produce  json
// @Param shipment body ShipmentReq true "shipment"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/shipment/add [post]
func (h *Six910Handler) AddShipment(w http.ResponseWriter, r *http.Request) {
	var addshURL = "/six910/rs/shipment/add"
	var ashc jv.Claim
	ashc.Role = storeAdmin
	ashc.URL = addshURL
	ashc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ashc)
	h.Log.Debug("shipment add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var shr ShipmentReq
			ashsuc, asherr := h.ProcessBody(r, &shr)
			h.Log.Debug("ashsuc: ", ashsuc)
			h.Log.Debug("shr: ", shr)
			h.Log.Debug("asherr: ", asherr)
			if !ashsuc && asherr != nil {
				http.Error(w, asherr.Error(), http.StatusBadRequest)
			} else {
				ashres := h.Manager.AddShipment(&shr.Shipment, shr.StoreID)
				h.Log.Debug("ashres: ", *ashres)
				if ashres.Success && ashres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ashres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ashf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ashf)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateShipment godoc
// @Summary Update a shipment
// @Description Update shipment data
// @Tags Shipment
// @Accept  json
// @Produce  json
// @Param shipment body ShipmentReq true "shipment"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shipment/update [put]
func (h *Six910Handler) UpdateShipment(w http.ResponseWriter, r *http.Request) {
	var ushURL = "/six910/rs/shipment/update"
	var ushc jv.Claim
	ushc.Role = storeAdmin
	ushc.URL = ushURL
	ushc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ushc)
	h.Log.Debug("shipment update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var shr ShipmentReq
			ushsuc, usherr := h.ProcessBody(r, &shr)
			h.Log.Debug("ushsuc: ", ushsuc)
			h.Log.Debug("shr: ", shr)
			h.Log.Debug("usherr: ", usherr)
			if !ushsuc && usherr != nil {
				http.Error(w, usherr.Error(), http.StatusBadRequest)
			} else {
				ushres := h.Manager.UpdateShipment(&shr.Shipment, shr.StoreID)
				h.Log.Debug("ushres: ", *ushres)
				if ushres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ushres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ushf m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ushf)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetShipment godoc
// @Summary Get details of a shipment by id
// @Description Get details of a shipment
// @Tags Shipment
// @Accept  json
// @Produce  json
// @Param id path string true "shipment id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Shipment
// @Router /rs/shipment/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetShipment(w http.ResponseWriter, r *http.Request) {
	var gshURL = "/six910/rs/shipment/get"
	var gshc jv.Claim
	gshc.Role = customerRole
	gshc.URL = gshURL
	gshc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshc)
	h.Log.Debug("shipment get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gshidStr = vars["id"]
			var gshstoreIDStr = vars["storeId"]
			id, gshiderr := strconv.ParseInt(gshidStr, 10, 64)
			storeID, gshsiderr := strconv.ParseInt(gshstoreIDStr, 10, 64)
			var gshres *sdbi.Shipment
			if gshiderr == nil && gshsiderr == nil {
				gshres = h.Manager.GetShipment(id, storeID)
				h.Log.Debug("gshres: ", gshres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Shipment
				gshres = &nc
			}
			resJSON, _ := json.Marshal(gshres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetShipmentList godoc
// @Summary Get list of shipment
// @Description Get list of shipment for a store
// @Tags Shipment
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
// @Success 200 {array} six910-database-interface.Shipment
// @Router /rs/shipment/get/list/{orderId}/{storeId} [get]
func (h *Six910Handler) GetShipmentList(w http.ResponseWriter, r *http.Request) {
	var gshlURL = "/six910/rs/shipment/list"
	var gshcl jv.Claim
	gshcl.Role = customerRole
	gshcl.URL = gshlURL
	gshcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshcl)
	h.Log.Debug("shipment get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var shloidStr = vars["orderId"]
			var shlstoreIDStr = vars["storeId"]
			orderID, sshloiderr := strconv.ParseInt(shloidStr, 10, 64)
			storeID, sshlserr := strconv.ParseInt(shlstoreIDStr, 10, 64)
			var gshlres *[]sdbi.Shipment
			if sshloiderr == nil && sshlserr == nil {
				gshlres = h.Manager.GetShipmentList(orderID, storeID)
				h.Log.Debug("get shipment list: ", gshlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Shipment{}
				gshlres = &nc
			}
			resJSON, _ := json.Marshal(gshlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteShipment godoc
// @Summary Delete a shipment
// @Description Delete a shipment from the store
// @Tags Shipment
// @Accept  json
// @Produce  json
// @Param id path string true "shipment id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shipment/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteShipment(w http.ResponseWriter, r *http.Request) {
	var dshURL = "/six910/rs/shipment/delete"
	var dshc jv.Claim
	dshc.Role = storeAdmin
	dshc.URL = dshURL
	dshc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dshc)
	h.Log.Debug("shipment delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dshidStr = vars["id"]
			var dshstoreIDStr = vars["storeId"]
			id, dshiderr := strconv.ParseInt(dshidStr, 10, 64)
			storeID, dshidserr := strconv.ParseInt(dshstoreIDStr, 10, 64)
			var dshres *m.Response
			if dshiderr == nil && dshidserr == nil {
				dshres = h.Manager.DeleteShipment(id, storeID)
				h.Log.Debug("dshres: ", *dshres)
				if dshres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dshres = &nc
			}
			resJSON, _ := json.Marshal(dshres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
