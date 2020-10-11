package service

import (
	"net"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/nullc0rp/go-ip2proxy-api/service"

	"github.com/nullc0rp/go-ip2proxy-api/database"
	"github.com/stretchr/testify/assert"
)

func TestGetIPInfoHappy(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"proxy_type", "country_code", "country_name", "region_name", "city_name", "isp", "domain", "usage_type", "asn", "as"}).
		AddRow("PUB", "PL", "Poland", "Mazowieckie", "Warsaw", "Opera Software ASA", "opera.com", "DCH", "1299", "as")

	mock.ExpectQuery("SELECT proxy_type,country_code,country_name,region_name,city_name,isp,domain,usage_type,asn,'as' FROM ip2proxy_database where ip_from <= 168430081 AND 168430081 <= ip_to;").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Test case
	ip := net.ParseIP("10.10.10.1")

	//Execution
	result, err := service.GetIPInfo(ip)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, "Warsaw", result.CityName, "")
	assert.Equal(t, "Mazowieckie", result.RegionName, "")
	assert.Equal(t, "Opera Software ASA", result.ISP, "")
}

func TestGetIPInfoNoResults(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"proxy_type", "country_code", "country_name", "region_name", "city_name", "isp", "domain", "usage_type", "asn", "as"})

	mock.ExpectQuery("SELECT proxy_type,country_code,country_name,region_name,city_name,isp,domain,usage_type,asn,'as' FROM ip2proxy_database where ip_from <= 168430081 AND 168430081 <= ip_to;").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Test case
	ip := net.ParseIP("10.10.10.1")

	//Execution
	result, err := service.GetIPInfo(ip)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Nil(t, result, "No result")
	assert.Equal(t, "No results for query, Please check your data, no results for query", err.Error(), "")
}

func TestGetIPCountryHappy(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"ip_from", "ip_to", "country_name", "region_name"}).
		AddRow(16778241, 16778241, "Australia", "Victoria").
		AddRow(16778242, 16778249, "Australia", "Victoria").
		AddRow(16778252, 16778259, "Australia", "Victoria")
	mock.ExpectQuery("SELECT ip_from,ip_to,country_name,city_name FROM ip2proxy_database where country_code = 'AR' LIMIT 10;").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.GetIPCountry("AR", 10)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 10, result.Total, "")
	assert.Equal(t, "Victoria", result.IPList[0].CityName, "")
}

func TestGetIPCountryNoResult(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"ip_from", "ip_to", "country_name", "region_name"})
	mock.ExpectQuery("SELECT ip_from,ip_to,country_name,city_name FROM ip2proxy_database where country_code = 'AR' LIMIT 10;").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.GetIPCountry("AR", 10)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 0, result.Total, "")
	assert.Empty(t, result.IPList, "")
}

func TestGetISPCountryHappy(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"isp"}).
		AddRow("ISP1").AddRow("ISP2").AddRow("ISP3").AddRow("ISP4")
	mock.ExpectQuery("SELECT isp FROM ip2proxy_database where country_code = 'AR'").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.GetISPCountry("AR")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 4, result.Total, "")
}

func TestGetISPCountryNoResults(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"isp"})
	mock.ExpectQuery("SELECT isp FROM ip2proxy_database where country_code = 'AR'").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.GetISPCountry("AR")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 0, result.Total, "")
}

func TestGetCountryTotalHappy(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"total_ip"}).
		AddRow(1337)
	mock.ExpectQuery(".*").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.GetCountryTotal("AR")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 1337, result.Total, "")
}

func TestGetCountryTotalNoResult(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"total_ip"}).
		AddRow(0)
	mock.ExpectQuery(".*").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.GetCountryTotal("AR")

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 0, result.Total, "")
}

func TestMostProxyTypesHappy(t *testing.T) {

	// Create mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expected data result
	rows := sqlmock.NewRows([]string{"proxy_type", "total"}).
		AddRow("ISP", 100).AddRow("SOMETHING", 200).AddRow("ELSE", 300)
	mock.ExpectQuery(".*").WillReturnRows(rows)

	//Instance services
	database := &database.DatabaseImpl{
		Connection: db,
	}

	service := &service.ServiceImp{
		DB: database,
	}

	//Execution
	result, err := service.MostProxyTypes()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, 3, len(result.ProxyTypeList), "")
}
