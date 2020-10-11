package controller

import (
	"testing"
)

func TestGetIPInfoHappy(t *testing.T) {
	//Work in progress
	/*
		//Mocked service setup
		controller := gomock.NewController(t)
		defer controller.Finish()
		mockService := mocks.NewMockService(controller)

		//Test case
		ip := net.ParseIP("10.10.10.1")

		//Mock response
		ipResponse := &service.IPData{
			CountryName: "Javalandia",
		}

		//Expects setup
		mockService.EXPECT().GetIPInfo(ip).Return(ipResponse)

		controllerTest := &ControllerImpl{
			Service: mockService,
		}

		//Execution
		result, err := controllerTest.GetIPInfo(ip)

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Equal(t, "Warsaw", result.CityName, "")
		assert.Equal(t, "Mazowieckie", result.RegionName, "")
		assert.Equal(t, "Opera Software ASA", result.ISP, "")
	*/
}
