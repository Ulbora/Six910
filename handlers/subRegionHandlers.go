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

//SubRegionReq SubRegionReq
type SubRegionReq struct {
	StoreID   int64          `json:"storeId"`
	SubRegion sdbi.SubRegion `json:"subRegion"`
}

//AddSubRegion AddSubRegion
func (h *Six910Handler) AddSubRegion(w http.ResponseWriter, r *http.Request) {
	var addsregURL = "/six910/rs/subRegion/add"
	var asregc jv.Claim
	asregc.Role = storeAdmin
	asregc.URL = addsregURL
	asregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &asregc)
	h.Log.Debug("sub region add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var asreg SubRegionReq
			asregsuc, asregerr := h.ProcessBody(r, &asreg)
			h.Log.Debug("asregsuc: ", asregsuc)
			h.Log.Debug("asreg: ", asreg)
			h.Log.Debug("asregerr: ", asregerr)
			if !asregsuc && asregerr != nil {
				http.Error(w, asregerr.Error(), http.StatusBadRequest)
			} else {
				asregres := h.Manager.AddSubRegion(&asreg.SubRegion, asreg.StoreID)
				h.Log.Debug("asregres: ", *asregres)
				if asregres.Success && asregres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(asregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var asregfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(asregfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//UpdateSubRegion UpdateSubRegion
func (h *Six910Handler) UpdateSubRegion(w http.ResponseWriter, r *http.Request) {
	var upsregURL = "/six910/rs/subRegion/update"
	var usregc jv.Claim
	usregc.Role = storeAdmin
	usregc.URL = upsregURL
	usregc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &usregc)
	h.Log.Debug("sub region update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var usreg SubRegionReq
			usregsuc, usregerr := h.ProcessBody(r, &usreg)
			h.Log.Debug("usregsuc: ", usregsuc)
			h.Log.Debug("usreg: ", usreg)
			h.Log.Debug("usregerr: ", usregerr)
			if !usregsuc && usregerr != nil {
				http.Error(w, usregerr.Error(), http.StatusBadRequest)
			} else {
				usregres := h.Manager.UpdateSubRegion(&usreg.SubRegion, usreg.StoreID)
				h.Log.Debug("usregres: ", *usregres)
				if usregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(usregres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var usregfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(usregfl)
		fmt.Fprint(w, string(resJSON))
	}
}

//GetSubRegion GetSubRegion
func (h *Six910Handler) GetSubRegion(w http.ResponseWriter, r *http.Request) {
	var gsregURL = "/six910/rs/subRegion/get"
	var gsregc jv.Claim
	gsregc.Role = customerRole
	gsregc.URL = gsregURL
	gsregc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gsregc)
	h.Log.Debug("sub region get id authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gsregidStr = vars["id"]
			var gsregstoreIDStr = vars["storeId"]
			id, gsregiderr := strconv.ParseInt(gsregidStr, 10, 64)
			storeID, gsregsiderr := strconv.ParseInt(gsregstoreIDStr, 10, 64)
			var gsregres *sdbi.SubRegion
			if gsregiderr == nil && gsregsiderr == nil {
				gsregres = h.Manager.GetSubRegion(id, storeID)
				h.Log.Debug("gsregres: ", gsregres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.SubRegion
				gsregres = &nc
			}
			resJSON, _ := json.Marshal(gsregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//GetSubRegionList GetSubRegionList
func (h *Six910Handler) GetSubRegionList(w http.ResponseWriter, r *http.Request) {
	var gsreglURL = "/six910/rs/subRegion/list"
	var gsregcl jv.Claim
	gsregcl.Role = customerRole
	gsregcl.URL = gsreglURL
	gsregcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gsregcl)
	h.Log.Debug("sub region get list authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var regIDlStr = vars["regionId"]
			var sreglstoreIDStr = vars["storeId"]
			regionID, regIDlerr := strconv.ParseInt(regIDlStr, 10, 64)
			storeID, sreglerr := strconv.ParseInt(sreglstoreIDStr, 10, 64)
			var gsreglres *[]sdbi.SubRegion
			if regIDlerr == nil && sreglerr == nil {
				gsreglres = h.Manager.GetSubRegionList(regionID, storeID)
				h.Log.Debug("get sub region list: ", gsreglres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.SubRegion{}
				gsreglres = &nc
			}
			resJSON, _ := json.Marshal(gsreglres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//DeleteSubRegion DeleteSubRegion
func (h *Six910Handler) DeleteSubRegion(w http.ResponseWriter, r *http.Request) {
	var dsregURL = "/six910/rs/subRegion/delete"
	var dsregs jv.Claim
	dsregs.Role = storeAdmin
	dsregs.URL = dsregURL
	dsregs.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dsregs)
	h.Log.Debug("sub region delete authorized: ", auth)
	if auth {
		h.SetContentType(w)
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dsregidStr = vars["id"]
			var dsregstoreIDStr = vars["storeId"]
			id, dsregiderr := strconv.ParseInt(dsregidStr, 10, 64)
			storeID, dsregidserr := strconv.ParseInt(dsregstoreIDStr, 10, 64)
			var dsregres *m.Response
			if dsregiderr == nil && dsregidserr == nil {
				dsregres = h.Manager.DeleteSubRegion(id, storeID)
				h.Log.Debug("dsregres: ", dsregres)
				if dsregres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dsregres = &nc
			}
			resJSON, _ := json.Marshal(dsregres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
