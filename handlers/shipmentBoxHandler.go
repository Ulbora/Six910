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

//ShipmentBoxReq ShipmentBoxReq
type ShipmentBoxReq struct {
	StoreID     int64            `json:"storeId"`
	ShipmentBox sdbi.ShipmentBox `json:"shipmentBox"`
}

// AddShipmentBox godoc
// @Summary Add a new shipmentBox
// @Description Adds a new shipmentBox to a store
// @Tags ShipmentBox
// @Accept  json
// @Produce  json
// @Param shipmentBox body ShipmentBoxReq true "shipmentBox"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/shipmentBox/add [post]
func (h *Six910Handler) AddShipmentBox(w http.ResponseWriter, r *http.Request) {
	var addshbURL = "/six910/rs/shipmentBox/add"
	var ashbc jv.Claim
	ashbc.Role = storeAdmin
	ashbc.URL = addshbURL
	ashbc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ashbc)
	h.Log.Debug("shipment box add authorized: ", auth)
	h.SetContentType(w)
	if auth {
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var shbr ShipmentBoxReq
			ashbsuc, ashberr := h.ProcessBody(r, &shbr)
			h.Log.Debug("ashbsuc: ", ashbsuc)
			h.Log.Debug("shbr: ", shbr)
			h.Log.Debug("ashberr: ", ashberr)
			if !ashbsuc && ashberr != nil {
				http.Error(w, ashberr.Error(), http.StatusBadRequest)
			} else {
				ashbres := h.Manager.AddShipmentBox(&shbr.ShipmentBox, shbr.StoreID)
				h.Log.Debug("ashbres: ", *ashbres)
				if ashbres.Success && ashbres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ashbres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ashbf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ashbf)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateShipmentBox godoc
// @Summary Update a shipmentBox
// @Description Update shipmentBox data
// @Tags ShipmentBox
// @Accept  json
// @Produce  json
// @Param shipmentBox body ShipmentBoxReq true "shipmentBox"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/shipmentBox/update [put]
func (h *Six910Handler) UpdateShipmentBox(w http.ResponseWriter, r *http.Request) {
	var ushbURL = "/six910/rs/shipmentBox/update"
	var ushbc jv.Claim
	ushbc.Role = storeAdmin
	ushbc.URL = ushbURL
	ushbc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ushbc)
	h.Log.Debug("shipment box update authorized: ", auth)
	h.SetContentType(w)
	if auth {
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var shbr ShipmentBoxReq
			ushbsuc, ushberr := h.ProcessBody(r, &shbr)
			h.Log.Debug("ushbsuc: ", ushbsuc)
			h.Log.Debug("shbr: ", shbr)
			h.Log.Debug("ushberr: ", ushberr)
			if !ushbsuc && ushberr != nil {
				http.Error(w, ushberr.Error(), http.StatusBadRequest)
			} else {
				ushbres := h.Manager.UpdateShipmentBox(&shbr.ShipmentBox, shbr.StoreID)
				h.Log.Debug("ushbres: ", *ushbres)
				if ushbres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ushbres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ushbf m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ushbf)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetShipmentBox godoc
// @Summary Get details of a shipmentBox by id
// @Description Get details of a shipmentBox
// @Tags ShipmentBox
// @Accept  json
// @Produce  json
// @Param id path string true "shipmentBox id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.ShipmentBox
// @Router /rs/shipmentBox/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetShipmentBox(w http.ResponseWriter, r *http.Request) {
	var gshbURL = "/six910/rs/shipmentBox/get"
	var gshbc jv.Claim
	gshbc.Role = customerRole
	gshbc.URL = gshbURL
	gshbc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshbc)
	h.Log.Debug("shipment box get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gshbidStr = vars["id"]
			var gshbstoreIDStr = vars["storeId"]
			id, gshbiderr := strconv.ParseInt(gshbidStr, 10, 64)
			storeID, gshbsiderr := strconv.ParseInt(gshbstoreIDStr, 10, 64)
			var gshbres *sdbi.ShipmentBox
			if gshbiderr == nil && gshbsiderr == nil {
				gshbres = h.Manager.GetShipmentBox(id, storeID)
				h.Log.Debug("gshbres: ", gshbres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.ShipmentBox
				gshbres = &nc
			}
			resJSON, _ := json.Marshal(gshbres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetShipmentBoxList godoc
// @Summary Get list of shipmentBox
// @Description Get list of shipmentBox for a store
// @Tags ShipmentBox
// @Accept  json
// @Produce  json
// @Param shipmentId path string true "shipment Id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.ShipmentBox
// @Router /rs/shipmentBox/get/list/{shipmentId}/{storeId} [get]
func (h *Six910Handler) GetShipmentBoxList(w http.ResponseWriter, r *http.Request) {
	var gshblURL = "/six910/rs/shipmentBox/list"
	var gshbcl jv.Claim
	gshbcl.Role = customerRole
	gshbcl.URL = gshblURL
	gshbcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshbcl)
	h.Log.Debug("shipment box get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var shbloidStr = vars["shipmentId"]
			var shblstoreIDStr = vars["storeId"]
			shipmentID, sshbloiderr := strconv.ParseInt(shbloidStr, 10, 64)
			storeID, sshblserr := strconv.ParseInt(shblstoreIDStr, 10, 64)
			var gshblres *[]sdbi.ShipmentBox
			if sshbloiderr == nil && sshblserr == nil {
				gshblres = h.Manager.GetShipmentBoxList(shipmentID, storeID)
				h.Log.Debug("get shipment box list: ", gshblres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.ShipmentBox{}
				gshblres = &nc
			}
			resJSON, _ := json.Marshal(gshblres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteShipmentBox godoc
// @Summary Delete a shipmentBox
// @Description Delete a shipmentBox from the store
// @Tags ShipmentBox
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
// @Router /rs/shipmentBox/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteShipmentBox(w http.ResponseWriter, r *http.Request) {
	var dshbURL = "/six910/rs/shipmentBox/delete"
	var dshbc jv.Claim
	dshbc.Role = storeAdmin
	dshbc.URL = dshbURL
	dshbc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dshbc)
	h.Log.Debug("shipment box delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dshbidStr = vars["id"]
			var dshbstoreIDStr = vars["storeId"]
			id, dshbiderr := strconv.ParseInt(dshbidStr, 10, 64)
			storeID, dshbidserr := strconv.ParseInt(dshbstoreIDStr, 10, 64)
			var dshbres *m.Response
			if dshbiderr == nil && dshbidserr == nil {
				dshbres = h.Manager.DeleteShipmentBox(id, storeID)
				h.Log.Debug("dshbres: ", *dshbres)
				if dshbres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dshbres = &nc
			}
			resJSON, _ := json.Marshal(dshbres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
