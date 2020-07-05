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

//ShipmentItemReq ShipmentItemReq
type ShipmentItemReq struct {
	StoreID      int64             `json:"storeId"`
	ShipmentItem sdbi.ShipmentItem `json:"shipmentItem"`
}

//AddShipmentItem AddShipmentItem
func (h *Six910Handler) AddShipmentItem(w http.ResponseWriter, r *http.Request) {
	var addshiURL = "/six910/rs/shipmentItem/add"
	var ashic jv.Claim
	ashic.Role = storeAdmin
	ashic.URL = addshiURL
	ashic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ashic)
	h.Log.Debug("shipment item add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var shir ShipmentItemReq
			ashisuc, ashierr := h.ProcessBody(r, &shir)
			h.Log.Debug("ashisuc: ", ashisuc)
			h.Log.Debug("shir: ", shir)
			h.Log.Debug("ashierr: ", ashierr)
			if !ashisuc && ashierr != nil {
				http.Error(w, ashierr.Error(), http.StatusBadRequest)
			} else {
				ashires := h.Manager.AddShipmentItem(&shir.ShipmentItem, shir.StoreID)
				h.Log.Debug("ashires: ", *ashires)
				if ashires.Success && ashires.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ashires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ashif m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ashif)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateShipmentItem UpdateShipmentItem
func (h *Six910Handler) UpdateShipmentItem(w http.ResponseWriter, r *http.Request) {
	var ushiURL = "/six910/rs/shipmentItem/update"
	var ushic jv.Claim
	ushic.Role = storeAdmin
	ushic.URL = ushiURL
	ushic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ushic)
	h.Log.Debug("shipment item update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var shir ShipmentItemReq
			ushisuc, ushierr := h.ProcessBody(r, &shir)
			h.Log.Debug("ushisuc: ", ushisuc)
			h.Log.Debug("shir: ", shir)
			h.Log.Debug("ushierr: ", ushierr)
			if !ushisuc && ushierr != nil {
				http.Error(w, ushierr.Error(), http.StatusBadRequest)
			} else {
				ushires := h.Manager.UpdateShipmentItem(&shir.ShipmentItem, shir.StoreID)
				h.Log.Debug("ushires: ", *ushires)
				if ushires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ushires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ushif m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ushif)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetShipmentItem GetShipmentItem
func (h *Six910Handler) GetShipmentItem(w http.ResponseWriter, r *http.Request) {
	var gshiURL = "/six910/rs/shipmentItem/get"
	var gshic jv.Claim
	gshic.Role = customerRole
	gshic.URL = gshiURL
	gshic.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshic)
	h.Log.Debug("shipment item get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gshiidStr = vars["id"]
			var gshistoreIDStr = vars["storeId"]
			id, gshiiderr := strconv.ParseInt(gshiidStr, 10, 64)
			storeID, gshisiderr := strconv.ParseInt(gshistoreIDStr, 10, 64)
			var gshires *sdbi.ShipmentItem
			if gshiiderr == nil && gshisiderr == nil {
				gshires = h.Manager.GetShipmentItem(id, storeID)
				h.Log.Debug("gshires: ", gshires)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.ShipmentItem
				gshires = &nc
			}
			resJSON, _ := json.Marshal(gshires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetShipmentItemList GetShipmentItemList
func (h *Six910Handler) GetShipmentItemList(w http.ResponseWriter, r *http.Request) {
	var gshilURL = "/six910/rs/shipmentItem/list"
	var gshicl jv.Claim
	gshicl.Role = customerRole
	gshicl.URL = gshilURL
	gshicl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshicl)
	h.Log.Debug("shipment item get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var shiloidStr = vars["shipmentId"]
			var shilstoreIDStr = vars["storeId"]
			shipmentID, sshiloiderr := strconv.ParseInt(shiloidStr, 10, 64)
			storeID, sshilserr := strconv.ParseInt(shilstoreIDStr, 10, 64)
			var gshilres *[]sdbi.ShipmentItem
			if sshiloiderr == nil && sshilserr == nil {
				gshilres = h.Manager.GetShipmentItemList(shipmentID, storeID)
				h.Log.Debug("get shipment item list: ", gshilres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.ShipmentItem{}
				gshilres = &nc
			}
			resJSON, _ := json.Marshal(gshilres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (h *Six910Handler) GetShipmentItemListByBox(w http.ResponseWriter, r *http.Request) {
	var gshilbURL = "/six910/rs/shipmentItem/list/box"
	var gshibcl jv.Claim
	gshibcl.Role = customerRole
	gshibcl.URL = gshilbURL
	gshibcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gshibcl)
	h.Log.Debug("shipment item get list by box authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var boxNumberStr = vars["boxNumber"]
			var shilbstoreIDStr = vars["storeId"]
			boxNumber, boxNumbererr := strconv.ParseInt(boxNumberStr, 10, 64)
			storeID, sshilbserr := strconv.ParseInt(shilbstoreIDStr, 10, 64)
			var gshilbres *[]sdbi.ShipmentItem
			if boxNumbererr == nil && sshilbserr == nil {
				gshilbres = h.Manager.GetShipmentItemListByBox(boxNumber, storeID)
				h.Log.Debug("get shipment item list by box: ", gshilbres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.ShipmentItem{}
				gshilbres = &nc
			}
			resJSON, _ := json.Marshal(gshilbres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteShipmentItem DeleteShipmentItem
func (h *Six910Handler) DeleteShipmentItem(w http.ResponseWriter, r *http.Request) {
	var dshiURL = "/six910/rs/shipmentItem/delete"
	var dshic jv.Claim
	dshic.Role = storeAdmin
	dshic.URL = dshiURL
	dshic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dshic)
	h.Log.Debug("shipment item delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dshiidStr = vars["id"]
			var dshistoreIDStr = vars["storeId"]
			id, dshiiderr := strconv.ParseInt(dshiidStr, 10, 64)
			storeID, dshiidserr := strconv.ParseInt(dshistoreIDStr, 10, 64)
			var dshires *m.Response
			if dshiiderr == nil && dshiidserr == nil {
				dshires = h.Manager.DeleteShipmentItem(id, storeID)
				h.Log.Debug("dshires: ", *dshires)
				if dshires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dshires = &nc
			}
			resJSON, _ := json.Marshal(dshires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
