package main

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
	"fmt"
	"net/http"
	"os"
	"strconv"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"

	hand "github.com/Ulbora/Six910/handlers"
	man "github.com/Ulbora/Six910/managers"
	db "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	sixmdb "github.com/Ulbora/six910-mysql"
	"github.com/gorilla/mux"

	jv "github.com/Ulbora/GoAuth2JwtValidator"
)

func main() {

	var mydb mdb.MyDB
	var proxy px.GoProxy
	var dbHost string
	var dbUser string
	var dbPassword string
	var dbName string

	var apiKey string

	var userHost string

	var l lg.Logger
	l.LogLevel = lg.AllLevel

	if os.Getenv("SIX910_DB_HOST") != "" {
		dbHost = os.Getenv("SIX910_DB_HOST")
	} else {
		dbHost = "localhost:3306"
	}

	if os.Getenv("SIX910_DB_USER") != "" {
		dbUser = os.Getenv("SIX910_DB_USER")
	} else {
		dbUser = "admin"
	}

	if os.Getenv("SIX910_DB_PASSWORD") != "" {
		dbPassword = os.Getenv("SIX910_DB_PASSWORD")
	} else {
		dbPassword = "admin"
	}

	if os.Getenv("SIX910_DB_DATABASE") != "" {
		dbName = os.Getenv("SIX910_DB_DATABASE")
	} else {
		dbName = "six910"
	}

	if os.Getenv("SIX910_USER_HOST") != "" {
		userHost = os.Getenv("SIX910_USER_HOST")
	} else {
		userHost = "http://localhost:3001"
	}

	if os.Getenv("SIX910_API_KEY") != "" {
		apiKey = os.Getenv("SIX910_API_KEY")
	} else {
		apiKey = "GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5"
	}

	mydb.Host = dbHost         // "localhost:3306"
	mydb.User = dbUser         // "admin"
	mydb.Password = dbPassword // "admin"
	mydb.Database = dbName     // "six910"
	var dbi db.Database = &mydb

	var sdb sixmdb.Six910Mysql
	//var l lg.Logger
	l.LogLevel = lg.AllLevel
	sdb.Log = &l
	sdb.DB = dbi
	dbi.Connect()

	var sm man.Six910Manager
	sm.Db = sdb.GetNew()
	sm.Log = &l
	sm.Proxy = proxy.GetNewProxy()
	sm.UserHost = userHost

	var sh hand.Six910Handler
	sh.Manager = sm.GetNew()
	sh.APIKey = apiKey
	sh.Log = &l

	var mc jv.MockOauthClient
	//mc.MockValidate = true
	sh.ValidatorClient = mc.GetNewClient()

	router := mux.NewRouter()
	port := "3002"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	h := sh.GetNew()

	//sdb.MockAddStoreSuccess = true
	//sdb.MockStoreID = 5

	//h := sh.GetNew()

	var locacc man.LocalStoreAdminUser
	locacc.Username = "admin"
	locacc.Password = "admin"

	lstoreRes := sm.CreateLocalStore(&locacc)
	sm.Log.Debug("Creating local store", *lstoreRes)

	//store
	router.HandleFunc("/rs/store/get/{storeName}/{localDomain}", h.GetStore).Methods("GET")

	fmt.Println("Six910 (six nine ten) server is running on port " + port + "!")

	http.ListenAndServe(":"+port, router)
}

// go mod init github.com/Ulbora/Six910
