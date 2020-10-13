package controller

import (
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/nullc0rp/go-ip2proxy-api/mocks"
	"github.com/nullc0rp/go-ip2proxy-api/service"
	"github.com/stretchr/testify/assert"
)

func TestGetIPInfoHappy(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/ip/10.10.10.1", strings.NewReader("address=10.10.10.1"))
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"address": "10.10.10.1",
	})

	//Test case
	ip := net.ParseIP("10.10.10.1")

	//Mock response
	ipResponse := &service.IPData{
		CountryName: "Javalandia",
	}

	//Expects setup
	mockService.EXPECT().GetIPInfo(ip).Return(ipResponse, nil)

	controllerInstance.GetIpInfo(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"proxy_type\":\"\",\"country_code\":\"\",\"country_name\":\"Javalandia\",\"region_name\":\"\",\"city_name\":\"\",\"isp\":\"\",\"domain\":\"\",\"usage_type\":\"\",\"asn\":\"\",\"as\":\"\"}")
}

func TestGetIPInfoBadIP(t *testing.T) {
	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/ip/10.10.10.asd", strings.NewReader("address=10.10.10.1"))
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"address": "10.10.10.asd",
	})

	controllerInstance.GetIpInfo(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "Bad IP address")

}

func TestGetIPListHappy(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR", strings.NewReader("limit=3"))
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Test case
	ipDataResult := &service.IPDataResult{
		CountryName: "Argentina",
	}

	//Mock response
	ipResponse := &service.IPCountryData{
		Total:  100,
		IPList: []*service.IPDataResult{ipDataResult},
	}

	//Expects setup
	mockService.EXPECT().GetIPCountry("AR", 50).Return(ipResponse, nil)

	controllerInstance.GetIpList(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total\":100,\"IPList\":[{\"ip_address\":\"\",\"country_name\":\"Argentina\",\"city_name\":\"\"}]}")
}

func TestGetIPListBadCountry(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR", strings.NewReader("limit=3"))
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "ARSAR'<someScript>;'Select",
	})

	//Test case
	ipDataResult := &service.IPDataResult{
		CountryName: "Argentina",
	}

	//Mock response
	ipResponse := &service.IPCountryData{
		Total:  100,
		IPList: []*service.IPDataResult{ipDataResult},
	}

	//Expects setup
	mockService.EXPECT().GetIPCountry("AR", 50).Return(ipResponse, nil)

	controllerInstance.GetIpList(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total\":100,\"IPList\":[{\"ip_address\":\"\",\"country_name\":\"Argentina\",\"city_name\":\"\"}]}")
}

func TestGetIPListNoResult(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR", strings.NewReader("limit=3"))
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "ARSAR'<someScript>;'Select",
	})

	//Mock response
	ipResponse := &service.IPCountryData{
		Total:  0,
		IPList: nil,
	}

	//Expects setup
	mockService.EXPECT().GetIPCountry("AR", 50).Return(ipResponse, nil)

	controllerInstance.GetIpList(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total\":0,\"IPList\":null}")
}

func TestGetIPListErrorService(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR", strings.NewReader("limit=3"))
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "ARSAR'<someScript>;'Select",
	})

	//Expects setup
	mockService.EXPECT().GetIPCountry("AR", 50).Return(nil, errors.New("some dirty info"))

	controllerInstance.GetIpList(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "Service error")
}

func TestGetIPTotalHappy(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR/total", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Mock Response
	ipDataResult := &service.IPCountryTotal{
		Total: 100,
	}

	//Expects setup
	mockService.EXPECT().GetCountryTotal("AR").Return(ipDataResult, nil)

	controllerInstance.GetIPTotalCountry(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total_ip\":100}")
}

func TestGetIPTotalNoResult(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR/total", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Mock Response
	ipDataResult := &service.IPCountryTotal{
		Total: 0,
	}

	//Expects setup
	mockService.EXPECT().GetCountryTotal("AR").Return(ipDataResult, nil)

	controllerInstance.GetIPTotalCountry(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total_ip\":0}")
}

func TestGetIPTotalError(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR/total", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Expects setup
	mockService.EXPECT().GetCountryTotal("AR").Return(nil, errors.New("some dirty info"))

	controllerInstance.GetIPTotalCountry(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "Service error")
}

func TestGetISPListHappy(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR/isp", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Test case
	ipDataResult := &service.ISPDataResult{
		Name: "ISP Name",
	}

	//Mock response
	ipResponse := &service.ISPCountryData{
		Total:   100,
		ISPList: []*service.ISPDataResult{ipDataResult},
	}

	//Expects setup
	mockService.EXPECT().GetISPCountry("AR").Return(ipResponse, nil)

	controllerInstance.GetISPCountry(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total\":100,\"ISPList\":[{\"isp\":\"ISP Name\"}]}")
}

func TestGetISPListNoResult(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR/isp", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Mock response
	ipResponse := &service.ISPCountryData{
		Total:   0,
		ISPList: nil,
	}

	//Expects setup
	mockService.EXPECT().GetISPCountry("AR").Return(ipResponse, nil)

	controllerInstance.GetISPCountry(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"total\":0,\"ISPList\":null}")
}

func TestGetISPListErrorService(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/country/AR/isp", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{
		"country": "AR",
	})

	//Expects setup
	mockService.EXPECT().GetISPCountry("AR").Return(nil, errors.New("some dirty info"))

	controllerInstance.GetISPCountry(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "Service error")
}

func TestGetMostProxyTypesHappy(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/proxytypes", nil)
	w := httptest.NewRecorder()

	//Mock response+
	proxyType := &service.MostProxyType{
		ProxyType: "Proxytype1",
		Total:     100,
	}

	proxyType2 := &service.MostProxyType{
		ProxyType: "Proxytype2",
		Total:     101,
	}

	ipResponse := &service.MostProxyTypeResult{
		ProxyTypeList: []*service.MostProxyType{proxyType, proxyType2},
	}

	//Expects setup
	mockService.EXPECT().MostProxyTypes().Return(ipResponse, nil)

	controllerInstance.GetMostProxyTypes(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"ProxyTypeList\":[{\"proxy_type\":\"Proxytype1\",\"total\":100},{\"proxy_type\":\"Proxytype2\",\"total\":101}]}")
}

func TestGetMostProxyTypesNoResult(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/proxytypes", nil)
	w := httptest.NewRecorder()

	ipResponse := &service.MostProxyTypeResult{
		ProxyTypeList: []*service.MostProxyType{},
	}

	//Expects setup
	mockService.EXPECT().MostProxyTypes().Return(ipResponse, nil)

	controllerInstance.GetMostProxyTypes(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "{\"ProxyTypeList\":[]}")
}

func TestGetMostProxyTypesErrorService(t *testing.T) {

	//Mocked service setup
	controller := gomock.NewController(t)
	defer controller.Finish()
	var controllerInstance Controller
	mockService := mocks.NewMockService(controller)

	//Instance Controller
	controllerInstance = &ControllerImpl{
		Service: mockService,
	}

	r, _ := http.NewRequest("GET", "/proxytypes", nil)
	w := httptest.NewRecorder()

	//Expects setup
	mockService.EXPECT().MostProxyTypes().Return(nil, errors.New("some dirty info"))

	controllerInstance.GetMostProxyTypes(w, r)

	assert.Equal(t, string(w.Body.Bytes()), "Service error")
}
