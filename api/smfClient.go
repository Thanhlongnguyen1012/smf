package api

import (
	"fmt"
	"net/http"
	"smf/models"
	"time"
)

type smfClientStruct struct {
	BaseURL    string
	HTTPClient *http.Client
}

var SmfClient = smfClientStruct{
	BaseURL: "http://localhost:9090", // địa chỉ của UDM
	HTTPClient: &http.Client{
		Timeout: 5 * time.Second,
	},
}

// luồng pduSession establishment
func (s *smfClientStruct) PduSessionCreation(data models.SMContextCreateData) (*http.Response, error) {
	//chuyển đổi struct thành file dạng byte để gửi đi trong phần body của POST
	//jsonData, err := json.Marshal(data)
	//gửi request đến nrf để thực hiện yêu cầu query UDM
	// truy xuất dữ liệu đăng kí phiên
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/nudm-sdm/v2/%s/sm-data", s.BaseURL, data.Supi), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	udmResponse, err := s.HTTPClient.Do(req)
	// if udmResponse.StatusCode == 201 {
	// 	fmt.Printf("gửi bản tin smContextCreated")
	// }
	return udmResponse, err
}

// func (s *smfClientStruct) PostSmContextsResponse(data models.SMContextCreatedData) (*http.Response, error) {
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req, err := http.NewRequest("POST", s.BaseURL+"/nf/", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	return s.HTTPClient.Do(req)
// }
