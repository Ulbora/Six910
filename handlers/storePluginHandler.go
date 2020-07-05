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

//AddStorePlugin AddStorePlugin
func (h *Six910Handler) AddStorePlugin(w http.ResponseWriter, r *http.Request) {
	var addspiURL = "/six910/rs/storePlugin/add"
	var aspic jv.Claim
	aspic.Role = storeAdmin
	aspic.URL = addspiURL
	aspic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &aspic)
	h.Log.Debug("store plugin add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var aspi sdbi.StorePlugins
			aspisuc, aspierr := h.ProcessBody(r, &aspi)
			h.Log.Debug("aspisuc: ", aspisuc)
			h.Log.Debug("aspi: ", aspi)
			h.Log.Debug("aspierr: ", aspierr)
			if !aspisuc && aspierr != nil {
				http.Error(w, aspierr.Error(), http.StatusBadRequest)
			} else {
				aspires := h.Manager.AddStorePlugin(&aspi)
				h.Log.Debug("aspires: ", *aspires)
				if aspires.Success && aspires.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(aspires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var aspifl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(aspifl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateStorePlugin UpdateStorePlugin
func (h *Six910Handler) UpdateStorePlugin(w http.ResponseWriter, r *http.Request) {
	var upspiURL = "/six910/rs/storePlugin/update"
	var uspic jv.Claim
	uspic.Role = storeAdmin
	uspic.URL = upspiURL
	uspic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &uspic)
	h.Log.Debug("store plugin update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var uspi sdbi.StorePlugins
			uspisuc, uspierr := h.ProcessBody(r, &uspi)
			h.Log.Debug("uspisuc: ", uspisuc)
			h.Log.Debug("uspi: ", uspi)
			h.Log.Debug("uspierr: ", uspierr)
			if !uspisuc && uspierr != nil {
				http.Error(w, uspierr.Error(), http.StatusBadRequest)
			} else {
				uspires := h.Manager.UpdateStorePlugin(&uspi)
				h.Log.Debug("uspires: ", *uspires)
				if uspires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(uspires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var uspifl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(uspifl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetStorePlugin GetStorePlugin
func (h *Six910Handler) GetStorePlugin(w http.ResponseWriter, r *http.Request) {
	var gspiURL = "/six910/rs/storePlugin/get"
	var gspic jv.Claim
	gspic.Role = customerRole
	gspic.URL = gspiURL
	gspic.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gspic)
	h.Log.Debug("store plugin get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gspiidStr = vars["id"]
			var gspistoreIDStr = vars["storeId"]
			id, gspiiderr := strconv.ParseInt(gspiidStr, 10, 64)
			storeID, gspisiderr := strconv.ParseInt(gspistoreIDStr, 10, 64)
			var gspires *sdbi.StorePlugins
			if gspiiderr == nil && gspisiderr == nil {
				gspires = h.Manager.GetStorePlugin(id, storeID)
				h.Log.Debug("gspires: ", gspires)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.StorePlugins
				gspires = &nc
			}
			resJSON, _ := json.Marshal(gspires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetStorePluginList GetStorePluginList
func (h *Six910Handler) GetStorePluginList(w http.ResponseWriter, r *http.Request) {
	var gspilURL = "/six910/rs/storePlugin/list"
	var gspicl jv.Claim
	gspicl.Role = customerRole
	gspicl.URL = gspilURL
	gspicl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gspicl)
	h.Log.Debug("store plugin get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var spilstoreIDStr = vars["storeId"]
			storeID, sspilerr := strconv.ParseInt(spilstoreIDStr, 10, 64)
			var gspilres *[]sdbi.StorePlugins
			if sspilerr == nil {
				gspilres = h.Manager.GetStorePluginList(storeID)
				h.Log.Debug("get store plugin list: ", gspilres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.StorePlugins{}
				gspilres = &nc
			}
			resJSON, _ := json.Marshal(gspilres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteStorePlugin DeleteStorePlugin
func (h *Six910Handler) DeleteStorePlugin(w http.ResponseWriter, r *http.Request) {
	var dspiURL = "/six910/rs/storePlugin/delete"
	var dspis jv.Claim
	dspis.Role = storeAdmin
	dspis.URL = dspiURL
	dspis.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dspis)
	h.Log.Debug("store plugin delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dspiidStr = vars["id"]
			var dspistoreIDStr = vars["storeId"]
			id, dspiiderr := strconv.ParseInt(dspiidStr, 10, 64)
			storeID, dspiidserr := strconv.ParseInt(dspistoreIDStr, 10, 64)
			var dspires *m.Response
			if dspiiderr == nil && dspiidserr == nil {
				dspires = h.Manager.DeleteStorePlugin(id, storeID)
				h.Log.Debug("delete store plugin: ", dspires)
				if dspires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dspires = &nc
			}
			resJSON, _ := json.Marshal(dspires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
