package service

// IPData info returned by controller
type IPData struct {
	ProxyType   string `json:"proxy_type"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	RegionName  string `json:"region_name"`
	CityName    string `json:"city_name"`
	ISP         string `json:"isp"`
	Domain      string `json:"domain"`
	UsageType   string `json:"usage_type"`
	ASN         string `json:"asn"`
	AS          string `json:"as"`
}

//IPDataSimple raw data from DB
type IPDataSimple struct {
	IPFrom      uint32 `json:"ip_from"`
	IPTo        uint32 `json:"ip_to"`
	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

//IPDataResult data formated for response
type IPDataResult struct {
	IP          string `json:"ip_address"`
	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
}

//ISPDataResult data formated for response
type ISPDataResult struct {
	Name string `json:"isp"`
}

//IPCountryData data formated for response
type IPCountryData struct {
	Total  int `json:"total"`
	IPList []*IPDataResult
}

//ISPCountryData data formated for response
type ISPCountryData struct {
	Total   int `json:"total"`
	ISPList []*ISPDataResult
}

//IPCountryTotal raw data from DB
type IPCountryTotal struct {
	Total int `json:"total_ip"`
}

//MostProxyType data formated for response
type MostProxyType struct {
	ProxyType string `json:"proxy_type"`
	Total     int    `json:"total"`
}

//MostProxyTypeResult data formated for response
type MostProxyTypeResult struct {
	ProxyTypeList []*MostProxyType
}
