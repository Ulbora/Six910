package handlers

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

import (
	"encoding/json"
	"errors"
	"net/http"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	m "github.com/Ulbora/Six910/managers"
)

//Six910Handler Six910Handler
type Six910Handler struct {
	Manager         m.Manager
	Proxy           px.Proxy
	Log             *lg.Logger
	APIKey          string
	ValidatorClient jv.Client
	ValidationURL   string
}

//GetNew GetNew
func (h *Six910Handler) GetNew() Handler {
	return h
}

//CheckContent CheckContent
func (h *Six910Handler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		// http.Error(w, "json required", http.StatusUnsupportedMediaType)
		rtn = true
	}
	return rtn
}

//SetContentType SetContentType
func (h *Six910Handler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//ProcessBody ProcessBody
func (h *Six910Handler) ProcessBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	h.Log.Debug("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}
