package models

type Guami struct {
	PlmnId *PlmnId `json:"plmnId"`
	AmfId  string  `json:"amfId"`
}
