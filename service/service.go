package service

import (
	"fmt"
	"log"
	"net"

	"github.com/nullc0rp/go-ip2proxy-api/database"
)

const (
	ERROR               = "Error: %s"
	PROXYTYPE           = "proxy_type"
	COUNTRYCODE         = "country_code"
	COUNTRYNAME         = "country_name"
	REGIONNAME          = "region_name"
	CITYNAME            = "city_name"
	ISP                 = "isp"
	DOMAIN              = "domain"
	USAGETYPE           = "usage_type"
	ASN                 = "asn"
	AS                  = "'as'"
	IPFROM              = "ip_from"
	IPTO                = "ip_to"
	IPDATAQUERY         = "SELECT " + PROXYTYPE + "," + COUNTRYCODE + "," + COUNTRYNAME + "," + REGIONNAME + "," + CITYNAME + "," + ISP + "," + DOMAIN + "," + USAGETYPE + "," + ASN + "," + AS + " FROM ip2proxy_database where ip_from <= %d AND %d <= ip_to;"
	ISPCOUNTRYQUERY     = "SELECT " + ISP + " FROM ip2proxy_database where " + COUNTRYCODE + " = '%s'"
	IPCOUNTRYQUERY      = "SELECT " + IPFROM + "," + IPTO + "," + COUNTRYNAME + "," + CITYNAME + " FROM ip2proxy_database where " + COUNTRYCODE + " = '%s' LIMIT %d;"
	IPCOUNTRYTOTALQUERY = "SELECT SUM(cast(ip_to+1 as signed)-cast(ip_from as signed)) as total_ip FROM ip2proxy_database where " + COUNTRYCODE + " = '%s' LIMIT 1;"
	MOSTPROXYTYPES      = "SELECT " + PROXYTYPE + ",count(" + PROXYTYPE + ") as total FROM ip2proxy_database GROUP BY " + PROXYTYPE + " ORDER BY total DESC LIMIT 3;"
	CHECKDATA           = "Please check your data, no results for query"
	UNKNOWN             = "Unknown error"
)

//ServiceImp handles requests and interacts with the DB
type ServiceImp struct {
	DB database.Database
}

//Service interface
type Service interface {
	GetIPInfo(ip net.IP) (*IPData, error)
	GetIPCountry(country string, limit int) (*IPCountryData, error)
	GetISPCountry(country string) (*ISPCountryData, error)
	GetCountryTotal(country string) (*IPCountryTotal, error)
	MostProxyTypes() (*MostProxyTypeResult, error)
}

//GetIPInfo TODO:
func (s ServiceImp) GetIPInfo(ip net.IP) (*IPData, error) {

	//Get decimal ip value
	decimalIP := IP2int(ip)

	//Build query
	query := fmt.Sprintf(IPDATAQUERY, decimalIP, decimalIP)
	log.Print(query)

	//Fetch results
	results, err := s.DB.Query(query)
	if err != nil {
		log.Printf(ERROR, err)
		return nil, err
	}

	//Result carrier
	var ipdata IPData

	// For first row, scan the result data
	if results.Next() {
		err = results.Scan(&ipdata.ProxyType, &ipdata.CountryCode, &ipdata.CountryName, &ipdata.RegionName, &ipdata.CityName, &ipdata.ISP, &ipdata.Domain, &ipdata.UsageType, &ipdata.ASN, &ipdata.AS)
		if err != nil {
			log.Printf(ERROR, err)
			return nil, err
		}
	} else {
		log.Printf(CHECKDATA)
		return nil, NoResultError(CHECKDATA)
	}

	return &ipdata, nil
}

//GetIPCountry Gets an ammount of ip addresses for country
//To consider: the requeriment is to deliver the ammount of addresses specified by the limit value,
//This value can be used on a query as limit, but the result may return a bigger number of addresses given the data structure.
//To archive a query that returns the specified ammount of ip addresses we need to build a complex query or use Mysql variables to keep track of the given results
//Since a complex query is needed, the performance of the query may be affected, and using Mysql variables is not a good practice
//I've decided to go a simpler approach with limit. The results will be rendered in runtime and controlled before return. It exchanges performance for memory, which is acceptable in my opinion
func (s ServiceImp) GetIPCountry(country string, limit int) (*IPCountryData, error) {

	//Build query
	query := fmt.Sprintf(IPCOUNTRYQUERY, country, limit)
	log.Print(query)

	//Fetch results
	results, err := s.DB.Query(query)
	if err != nil {
		log.Printf(ERROR, err)
		return nil, err
	}

	//Result carrier
	IPList := []*IPDataResult{}
	var ipDataSimple IPDataSimple
	//Solution for the problem with limit
	counter := 0

	// For each row, scan the result into ipDataSimple
	for results.Next() {
		err = results.Scan(&ipDataSimple.IPFrom, &ipDataSimple.IPTo, &ipDataSimple.CountryName, &ipDataSimple.CityName)
		if err != nil {
			//Keep cycling
			log.Printf(ERROR, err)
		} else {
			//Creates the result object and counts the addresses
			//For each IP range, it builds the IPV4 address from IPFrom and then
			//increments IPFrom until is equal to IPTo.
			for ipDataSimple.IPFrom <= ipDataSimple.IPTo && counter < limit {
				localAdress := Int2IP(ipDataSimple.IPFrom)
				IPList = append(IPList, &IPDataResult{
					CountryName: ipDataSimple.CountryName,
					CityName:    ipDataSimple.CityName,
					IP:          localAdress.String(),
				})
				ipDataSimple.IPFrom = ipDataSimple.IPFrom + 1
				counter = counter + 1
			}
		}
	}

	//Result carrier
	IPCountryData := &IPCountryData{
		IPList: IPList,
		Total:  counter,
	}

	return IPCountryData, nil
}

//GetISPCountry Service to get all the ISP by country
func (s ServiceImp) GetISPCountry(country string) (*ISPCountryData, error) {

	//Build query
	query := fmt.Sprintf(ISPCOUNTRYQUERY, country)
	log.Print(query)

	//Fetch results
	results, err := s.DB.Query(query)
	if err != nil {
		log.Printf(ERROR, err)
		return nil, err
	}

	//Since a lot of ISP names are the same, we need to create a map for them
	//Golang alternative to "set" is to use a map with boolean for value
	set := make(map[string]bool)

	//Data holders
	ISPList := []*ISPDataResult{}
	var ispDataSimple ISPDataResult

	// For each row, scan the result into ispDataSimple
	for results.Next() {
		err = results.Scan(&ispDataSimple.Name)
		if err != nil {
			//Keep cycling
			log.Printf(ERROR, err)
		} else {
			//Save value in dictionary
			set[ispDataSimple.Name] = true
		}
	}

	//Loop over the dictionary to get the ISP names individually
	for k := range set {
		ISPList = append(ISPList, &ISPDataResult{
			Name: k,
		})
	}

	//Result carrier
	IPCountryData := &ISPCountryData{
		ISPList: ISPList,
		Total:   len(set),
	}

	return IPCountryData, nil
}

//GetCountryTotal Get the total ammount of ips for a country
func (s ServiceImp) GetCountryTotal(country string) (*IPCountryTotal, error) {

	//Build query
	query := fmt.Sprintf(IPCOUNTRYTOTALQUERY, country)
	log.Print(query)

	//Fetch results
	results, err := s.DB.Query(query)
	if err != nil {
		log.Printf(ERROR, err)
		return nil, err
	}

	//Result carrier
	ipCountryTotal := &IPCountryTotal{}

	// For first row, scan the result data
	if results.Next() {
		err = results.Scan(&ipCountryTotal.Total)
		if err != nil {
			log.Printf(ERROR, err)
			return nil, err
		}
	}

	return ipCountryTotal, nil
}

//MostProxyTypes Is the ServiceImp that gets the most proxy types in the database
//Note: The database has only one kind of proxy type in the whole table, so this always returns one result
func (s ServiceImp) MostProxyTypes() (*MostProxyTypeResult, error) {

	//Prepare query
	log.Print(MOSTPROXYTYPES)

	//Fetch results
	results, err := s.DB.Query(MOSTPROXYTYPES)
	if err != nil {
		log.Printf(ERROR, err)
		return nil, err
	}

	//Result carrier
	var mostProxyType MostProxyType
	mostProxyTypeList := []*MostProxyType{}

	// For first row, scan the result data
	for results.Next() {
		err = results.Scan(&mostProxyType.ProxyType, &mostProxyType.Total)
		if err != nil {
			//Keep cycling
			log.Printf(ERROR, err)
		} else {
			mostProxyTypeList = append(mostProxyTypeList, &MostProxyType{
				Total:     mostProxyType.Total,
				ProxyType: mostProxyType.ProxyType,
			})
		}
	}

	//Result object
	mostProxyTypeResult := &MostProxyTypeResult{
		ProxyTypeList: mostProxyTypeList,
	}

	return mostProxyTypeResult, nil
}
