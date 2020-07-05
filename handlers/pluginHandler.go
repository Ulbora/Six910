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

//AddPlugin AddPlugin
func (h *Six910Handler) AddPlugin(w http.ResponseWriter, r *http.Request) {
	var addpiURL = "/six910/rs/plugin/add"
	var apic jv.Claim
	apic.Role = superAdmin
	apic.URL = addpiURL
	apic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &apic)
	h.Log.Debug("plugin add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var api sdbi.Plugins
			apisuc, apierr := h.ProcessBody(r, &api)
			h.Log.Debug("apisuc: ", apisuc)
			h.Log.Debug("apit: ", api)
			h.Log.Debug("apierr: ", apierr)
			if !apisuc && apierr != nil {
				http.Error(w, apierr.Error(), http.StatusBadRequest)
			} else {
				apires := h.Manager.AddPlugin(&api)
				h.Log.Debug("apires: ", *apires)
				if apires.Success && apires.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(apires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var apifl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(apifl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdatePlugin UpdatePlugin
func (h *Six910Handler) UpdatePlugin(w http.ResponseWriter, r *http.Request) {
	var uppiURL = "/six910/rs/plugin/update"
	var upic jv.Claim
	upic.Role = superAdmin
	upic.URL = uppiURL
	upic.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &upic)
	h.Log.Debug("plugins update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var upi sdbi.Plugins
			upisuc, upierr := h.ProcessBody(r, &upi)
			h.Log.Debug("upisuc: ", upisuc)
			h.Log.Debug("upi: ", upi)
			h.Log.Debug("upierr: ", upierr)
			if !upisuc && upierr != nil {
				http.Error(w, upierr.Error(), http.StatusBadRequest)
			} else {
				upires := h.Manager.UpdatePlugin(&upi)
				h.Log.Debug("upires: ", *upires)
				if upires.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(upires)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var upifl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(upifl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetPlugin GetPlugin
func (h *Six910Handler) GetPlugin(w http.ResponseWriter, r *http.Request) {
	var gpiURL = "/six910/rs/plugin/get"
	var gpic jv.Claim
	gpic.Role = customerRole
	gpic.URL = gpiURL
	gpic.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gpic)
	h.Log.Debug("plugin get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var gpiidStr = vars["id"]
			id, gcatiderr := strconv.ParseInt(gpiidStr, 10, 64)
			var gpires *sdbi.Plugins
			if gcatiderr == nil {
				gpires = h.Manager.GetPlugin(id)
				h.Log.Debug("gpires: ", *gpires)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Plugins
				gpires = &nc
			}
			resJSON, _ := json.Marshal(gpires)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetPluginList GetPluginList
func (h *Six910Handler) GetPluginList(w http.ResponseWriter, r *http.Request) {
	var gpilURL = "/six910/rs/plugin/list"
	var gpicl jv.Claim
	gpicl.Role = customerRole
	gpicl.URL = gpilURL
	gpicl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gpicl)
	h.Log.Debug("plugin get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var pistartStr = vars["start"]
			var piendStr = vars["end"]
			plstart, pistlerr := strconv.ParseInt(pistartStr, 10, 64)
			plend, piedlerr := strconv.ParseInt(piendStr, 10, 64)
			var gpllres *[]sdbi.Plugins
			if pistlerr == nil && piedlerr == nil {
				gpllres = h.Manager.GetPluginList(plstart, plend)
				h.Log.Debug("get plugin list: ", *gpllres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Plugins{}
				gpllres = &nc
			}
			resJSON, _ := json.Marshal(gpllres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeletePlugin DeletePlugin
func (h *Six910Handler) DeletePlugin(w http.ResponseWriter, r *http.Request) {
	var dplURL = "/six910/rs/plugin/delete"
	var dpls jv.Claim
	dpls.Role = storeAdmin
	dpls.URL = dplURL
	dpls.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dpls)
	h.Log.Debug("plugin delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var dplidStr = vars["id"]
			id, dpliderr := strconv.ParseInt(dplidStr, 10, 64)
			var dplres *m.Response
			if dpliderr == nil {
				dplres = h.Manager.DeletePlugin(id)
				h.Log.Debug("delete plugin: ", dplres)
				if dplres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dplres = &nc
			}
			resJSON, _ := json.Marshal(dplres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
