package models

type SMContextCreateData struct {
	Supi string `json:"supi,omitempty"`
	//Pei            string      `json:"pei,omitempty"`
	//Gpsi           string      `json:"gpsi,omitempty"`
	PduSessionId int32  `json:"pduSessionId,omitempty"`
	Dnn          string `json:"dnn,omitempty"`
	SNssai       string `json:"sNssai,omitempty"`
	//SNssai       *Snssai `json:"sNssai,omitempty"`
	//ServingNfId    string      `json:"servingNfId"`
	//Guami          *Guami      `json:"guami,omitempty"`
	//ServingNetwork *PlmnId     `json:"servingNetwork"`
	//RequestType    RequestType `json:"requestType,omitempty"`
	//AnType         AccessType  `json:"anType"`
	//RatType        RatType     `json:"ratType,omitempty"`
	//UeLocation     *UserLocation `json:"ueLocation,omitempty"`
}
