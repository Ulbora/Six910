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

// AddCategory godoc
// @Summary Add a new category
// @Description Adds a new category to a store
// @Tags Category
// @Accept  json
// @Produce  json
// @Param category body six910-database-interface.Category true "category"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.ResponseID
// @Router /rs/category/add [post]
func (h *Six910Handler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var addCatURL = "/six910/rs/category/add"
	var acatc jv.Claim
	acatc.Role = storeAdmin
	acatc.URL = addCatURL
	acatc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &acatc)
	h.Log.Debug("cat add authorized: ", auth)
	if auth {
		h.SetContentType(w)
		acOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", acOk)
		if !acOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var acat sdbi.Category
			acatsuc, acaterr := h.ProcessBody(r, &acat)
			h.Log.Debug("acatsuc: ", acatsuc)
			h.Log.Debug("acat: ", acat)
			h.Log.Debug("acaterr: ", acaterr)
			if !acatsuc && acaterr != nil {
				http.Error(w, acaterr.Error(), http.StatusBadRequest)
			} else {
				acatres := h.Manager.AddCategory(&acat)
				h.Log.Debug("acatres: ", *acatres)
				if acatres.Success && acatres.ID != 0 {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(acatres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var actfl m.ResponseID
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(actfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update category data
// @Tags Category
// @Accept  json
// @Produce  json
// @Param category body six910-database-interface.Category true "category"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/category/update [put]
func (h *Six910Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var upcatURL = "/six910/rs/category/update"
	var ucatc jv.Claim
	ucatc.Role = storeAdmin
	ucatc.URL = upcatURL
	ucatc.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &ucatc)
	h.Log.Debug("cat update authorized: ", auth)
	if auth {
		h.SetContentType(w)
		ucOk := h.CheckContent(r)
		h.Log.Debug("conOk: ", ucOk)
		if !ucOk {
			http.Error(w, "json required", http.StatusUnsupportedMediaType)
		} else {
			var ucat sdbi.Category
			ucatsuc, ucaterr := h.ProcessBody(r, &ucat)
			h.Log.Debug("ucatsuc: ", ucatsuc)
			h.Log.Debug("ucat: ", ucat)
			h.Log.Debug("ucaterr: ", ucaterr)
			if !ucatsuc && ucaterr != nil {
				http.Error(w, ucaterr.Error(), http.StatusBadRequest)
			} else {
				ucatres := h.Manager.UpdateCategory(&ucat)
				h.Log.Debug("ucatres: ", *ucatres)
				if ucatres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				resJSON, _ := json.Marshal(ucatres)
				fmt.Fprint(w, string(resJSON))
			}
		}
	} else {
		var ucatfl m.Response
		w.WriteHeader(http.StatusUnauthorized)
		resJSON, _ := json.Marshal(ucatfl)
		fmt.Fprint(w, string(resJSON))
	}
}

// GetCategory godoc
// @Summary Get details of a category by id
// @Description Get details of a category
// @Tags Category
// @Accept  json
// @Produce  json
// @Param id path string true "category id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} six910-database-interface.Category
// @Router /rs/category/get/id/{id}/{storeId} [get]
func (h *Six910Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	var gcatURL = "/six910/rs/catetory/get"
	var gcatc jv.Claim
	gcatc.Role = customerRole
	gcatc.URL = gcatURL
	gcatc.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &gcatc)
	h.Log.Debug("cat get id authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var gcatidStr = vars["id"]
			var gcatstoreIDStr = vars["storeId"]
			id, gcatiderr := strconv.ParseInt(gcatidStr, 10, 64)
			storeID, gcatsiderr := strconv.ParseInt(gcatstoreIDStr, 10, 64)
			var gctres *sdbi.Category
			if gcatiderr == nil && gcatsiderr == nil {
				gctres = h.Manager.GetCategory(id, storeID)
				h.Log.Debug("gctres: ", gctres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc sdbi.Category
				gctres = &nc
			}
			resJSON, _ := json.Marshal(gctres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetCategoryList godoc
// @Summary Get list of categories
// @Description Get list of categories for a store
// @Tags Category
// @Accept  json
// @Produce  json
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Category
// @Router /rs/category/get/list/{storeId} [get]
func (h *Six910Handler) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	var gcatlURL = "/six910/rs/category/list"
	var gcatcl jv.Claim
	gcatcl.Role = customerRole
	gcatcl.URL = gcatlURL
	gcatcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("cat get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var ctlstoreIDStr = vars["storeId"]
			storeID, scatlerr := strconv.ParseInt(ctlstoreIDStr, 10, 64)
			var gcatlres *[]sdbi.Category
			if scatlerr == nil {
				gcatlres = h.Manager.GetCategoryList(storeID)
				h.Log.Debug("get cat list: ", gcatlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc = []sdbi.Category{}
				gcatlres = &nc
			}
			resJSON, _ := json.Marshal(gcatlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// GetSubCategoryList godoc
// @Summary Get list of sub categories
// @Description Get list of sub categories for a store
// @Tags Category
// @Accept  json
// @Produce  json
// @Param catId path string true "category id"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {array} six910-database-interface.Category
// @Router /rs/category/get/sub/list/{catId} [get]
func (h *Six910Handler) GetSubCategoryList(w http.ResponseWriter, r *http.Request) {
	var gscatlURL = "/six910/rs/subCategory/list"
	var gscatcl jv.Claim
	gscatcl.Role = customerRole
	gscatcl.URL = gscatlURL
	gscatcl.Scope = "read"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processAPIKeySecurity(r)
	h.Log.Debug("sub cat get list authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 1 {
			h.Log.Debug("vars: ", vars)
			var sctlcatIDStr = vars["catId"]
			catID, sscatlerr := strconv.ParseInt(sctlcatIDStr, 10, 64)
			var gscatlres *[]sdbi.Category
			if sscatlerr == nil {
				gscatlres = h.Manager.GetSubCategoryList(catID)
				h.Log.Debug("get sub cat list: ", gscatlres)
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nsc = []sdbi.Category{}
				gscatlres = &nsc
			}
			resJSON, _ := json.Marshal(gscatlres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category from the store
// @Tags Category
// @Accept  json
// @Produce  json
// @Param id path string true "category id"
// @Param storeId path string true "store storeId"
// @Param apiKey header string false "apiKey required for non OAuth2 stores only"
// @Param storeName header string true "store name"
// @Param localDomain header string true "store localDomain"
// @Param Authorization header string true "token"
// @Param clientId header string false "OAuth2 client ID only for OAuth2 stores"
// @Param userId header string false "User ID only for OAuth2 stores"
// @Success 200 {object} managers.Response
// @Router /rs/category/delete/{id}/{storeId} [delete]
func (h *Six910Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var dcatURL = "/six910/rs/category/delete"
	var dcts jv.Claim
	dcts.Role = storeAdmin
	dcts.URL = dcatURL
	dcts.Scope = "write"
	h.Log.Debug("client: ", h.ValidatorClient)
	auth := h.processSecurity(r, &dcts)
	h.Log.Debug("cat delete authorized: ", auth)
	h.SetContentType(w)
	if auth {
		vars := mux.Vars(r)
		h.Log.Debug("vars: ", len(vars))
		if vars != nil && len(vars) == 2 {
			h.Log.Debug("vars: ", vars)
			var dctidStr = vars["id"]
			var dctstoreIDStr = vars["storeId"]
			id, dctiderr := strconv.ParseInt(dctidStr, 10, 64)
			storeID, dctidserr := strconv.ParseInt(dctstoreIDStr, 10, 64)
			var dctres *m.Response
			if dctiderr == nil && dctidserr == nil {
				dctres = h.Manager.DeleteCategory(id, storeID)
				h.Log.Debug("deleteCat: ", dctres)
				if dctres.Success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				var nc m.Response
				dctres = &nc
			}
			resJSON, _ := json.Marshal(dctres)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
