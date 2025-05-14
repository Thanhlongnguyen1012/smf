package api

import (
	"net/http"
	"smf/models"
	"time"
)

type smfClientStruct struct {
	BaseURL    string
	HTTPClient *http.Client
}

var SmfClient = smfClientStruct{
	BaseURL: "http://localhost:9090", // địa chỉ UDM hoặc AMF
	HTTPClient: &http.Client{
		Timeout: 5 * time.Second,
	},
}

func (s *smfClientStruct) PduSessionCreation(data models.SMContextCreateData) (*http.Response, error) {
	/*jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}*/
	//req, err := http.NewRequest("POST", s.BaseURL+"/nf/pdu-session", bytes.NewBuffer(jsonData))
	req, err := http.NewRequest("POST", s.BaseURL+"/nf/pdu-session", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.HTTPClient.Do(req)
}
