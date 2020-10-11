package controller

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nullc0rp/go-ip2proxy-api/service"
)

//Controller handles requests and filter common requests
type ControllerImpl struct {
	Service service.Service
}

//Controller interface
type Controller interface {
	GetIpInfo(w http.ResponseWriter, r *http.Request)
	GetIpList(w http.ResponseWriter, r *http.Request)
	GetISPCountry(w http.ResponseWriter, r *http.Request)
	GetIPTotalCountry(w http.ResponseWriter, r *http.Request)
	GetMostProxyTypes(w http.ResponseWriter, r *http.Request)
}

const (
	MAXROWS      = 1000
	SERVICEERROR = "Service error"
	BADIPADDRESS = "Bad IP address"
	ERRORMARSHAL = "Error Marshal"
	CONTENTYPE   = "Content-Type"
	APPJSON      = "application/json"
	BADCOUNTRY   = "Bad country code"
	ERRORLIMIT   = "Error parsing limit"
	ERROR        = "Error"
	COUNTRY      = "country"
	ADDRESS      = "address"
)

//GetIpInfo is the controller for IP Information endpoint
func (c ControllerImpl) GetIpInfo(w http.ResponseWriter, r *http.Request) {

	// Get path vars
	vars := mux.Vars(r)

	log.Printf("Received request for ip info: %s\n", vars[ADDRESS])

	// Parse and validate ip address
	ip := net.ParseIP(vars[ADDRESS])
	if ip == nil {
		log.Println(BADIPADDRESS)
		WriteError(w, BADIPADDRESS)
		return
	}

	// Get service data
	result, err := c.Service.GetIPInfo(ip)
	if err != nil {
		log.Println(ERROR, ip, err)
		WriteError(w, SERVICEERROR)
		return
	}

	//Parse result data in Json format
	jData, err := json.Marshal(result)
	if err != nil {
		log.Println(ERRORMARSHAL)
		WriteError(w, SERVICEERROR)
		return
	}

	w.Header().Set(CONTENTYPE, APPJSON)
	w.Write(jData)
}

//GetIpList is the controller to ammount of addresses determined by limit parameter or 50 by default.
func (c ControllerImpl) GetIpList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Get and filter input values
	country := OnlyCountryCode(vars[COUNTRY])
	log.Printf("Received request for ip list by country %s\n", country)
	limit := OnlyInt(r.URL.Query().Get("limit"))
	log.Printf("Limit %s\n", limit)

	//Parse limit
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(ERRORLIMIT, err)
		intLimit = 0
	}

	// Default limit value
	if intLimit == 0 || intLimit > MAXROWS {
		intLimit = 50
	}

	// Get service data
	result, err := c.Service.GetIPCountry(country, intLimit)
	if err != nil {
		log.Println(ERROR, err)
		WriteError(w, SERVICEERROR)
		return
	}

	// Result json
	jData, err := json.Marshal(result)
	if err != nil {
		log.Println(ERRORMARSHAL)
		WriteError(w, SERVICEERROR)
		return
	}

	w.Header().Set(CONTENTYPE, APPJSON)
	w.Write(jData)
}

// GetISPCountry is the controller to get all the ISP by country
func (c ControllerImpl) GetISPCountry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Get and filter input values
	country := OnlyCountryCode(vars[COUNTRY])
	log.Printf("Received request for isp list by country %s\n", country)
	if len(country) < 2 {
		log.Println(BADCOUNTRY)
		WriteError(w, BADCOUNTRY)
		return
	}

	//Get service data
	result, err := c.Service.GetISPCountry(country)
	if err != nil {
		log.Println(ERROR, err)
		WriteError(w, SERVICEERROR)
		return
	}

	// Result json
	jData, err := json.Marshal(result)
	if err != nil {
		WriteError(w, SERVICEERROR)
		return
	}
	w.Header().Set(CONTENTYPE, APPJSON)
	w.Write(jData)
}

// GetIPTotalCountry is the controller to get the IP count for a country
func (c ControllerImpl) GetIPTotalCountry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Get and filter input values
	country := OnlyCountryCode(vars[COUNTRY])
	log.Printf("Received request for ip count by country %s\n", country)
	if len(country) < 2 {
		log.Println(BADCOUNTRY)
		WriteError(w, BADCOUNTRY)
		return
	}

	//Get service data
	result, err := c.Service.GetCountryTotal(country)
	if err != nil {
		log.Println(ERROR, err)
		WriteError(w, SERVICEERROR)
		return
	}

	// Result json
	jData, err := json.Marshal(result)
	if err != nil {
		WriteError(w, SERVICEERROR)
		return
	}
	w.Header().Set(CONTENTYPE, APPJSON)
	w.Write(jData)
}

// GetMostProxyTypes is the controller to get the top 3 most proxy types
func (c ControllerImpl) GetMostProxyTypes(w http.ResponseWriter, r *http.Request) {

	log.Println("Received request for most proxy types")

	//Get service data
	result, err := c.Service.MostProxyTypes()
	if err != nil {
		log.Println(ERROR, err)
		WriteError(w, SERVICEERROR)
		return
	}

	// Result json
	jData, err := json.Marshal(result)
	if err != nil {
		WriteError(w, SERVICEERROR)
		return
	}
	w.Header().Set(CONTENTYPE, APPJSON)
	w.Write(jData)
}
