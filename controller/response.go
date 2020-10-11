package controller

// IPDataResult info returned by controller
type IPDataResult struct {
	ProxyType   string
	CountryCode string
	CountryName string
	RegionName  string
	CityName    string
	ISP         string
	Domain      string
	UsageType   string
	ASN         string
	AS          string
}
