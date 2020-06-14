package managers

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
	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
	"golang.org/x/crypto/bcrypt"
)

//Six910Manager Six910Manager
type Six910Manager struct {
	Db    sdbi.Six910DB
	Proxy px.Proxy
	Log   *lg.Logger
}

//GetNew GetNew
func (m *Six910Manager) GetNew() Manager {
	return m
}

func (m *Six910Manager) hashPassword(pw string) (bool, string) {
	var suc bool
	var rtn string
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err == nil {
		rtn = string(hashedPw)
		suc = true
	}
	return suc, rtn
}

func (m *Six910Manager) validatePassword(pw string, hpw string) bool {
	var suc bool
	err := bcrypt.CompareHashAndPassword([]byte(hpw), []byte(pw))
	if err == nil {
		suc = true
	}
	return suc
}
