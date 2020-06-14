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
	//dbmod "github.com/Ulbora/six910-mysql"
	sdbi "github.com/Ulbora/six910-database-interface"
)

//LocalStoreAdminUser LocalStoreAdminUser
type LocalStoreAdminUser struct {
	Username string
	Password string
}

//User User
type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	Role        string `json:"role"`
	CustomerID  int64  `json:"customerId"`
	StoreID     int64  `json:"storeId"`
}

//ResponseID ResponseID
type ResponseID struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
}

//Response Response
type Response struct {
	Success bool `json:"success"`
}

//UserResonse UserResonse
type UserResonse struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	CustomerID int64  `json:"customerId"`
	StoreID    int64  `json:"storeId"`
}

//Manager Manager
type Manager interface {
	CreateLocalStore(auth *LocalStoreAdminUser)

	// Store --------------------------
	//Can only add a store when in external oauth mode
	AddStore(s *sdbi.Store) *ResponseID

	UpdateStore(s *sdbi.Store) *Response
	GetStore(sname string) *sdbi.Store
	GetStoreID(id int64) *sdbi.Store
	GetStoreByLocal(localDomain string) *sdbi.Store

	//can only delete a store when in the external oauth mode
	DeleteStore(id int64) bool

	//customer
	AddCustomer(c *sdbi.Customer) (bool, int64)
	UpdateCustomer(c *sdbi.Customer) bool
	GetCustomer(email string, storeID int64) *sdbi.Customer
	GetCustomerID(id int64) *sdbi.Customer
	GetCustomerList(storeID int64) *[]sdbi.Customer
	DeleteCustomer(id int64) bool

	AddUser(u *User) *ResponseID
	UpdateUser(u *User) *Response
	GetUser(id int64) *UserResonse
	GetUsers(storeID int64) *[]UserResonse
}
