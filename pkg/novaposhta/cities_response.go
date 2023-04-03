package novaposhta

type CitiesResponse struct {
	Success bool `json:"success"`
	Data    []struct {
		TotalCount int         `json:"TotalCount"`
		Addresses  []Addresses `json:"Addresses"`
	} `json:"data"`
	Errors       []interface{} `json:"errors"`
	MessageCodes []interface{} `json:"messageCodes"`
	ErrorCodes   []interface{} `json:"errorCodes"`
	WarningCodes []interface{} `json:"warningCodes"`
	InfoCodes    []interface{} `json:"infoCodes"`
}

type Addresses struct {
	Present                string `json:"Present"`
	Warehouses             int    `json:"Warehouses"`
	MainDescription        string `json:"MainDescription"`
	Area                   string `json:"Area"`
	Region                 string `json:"Region"`
	SettlementTypeCode     string `json:"SettlementTypeCode"`
	Ref                    string `json:"Ref"`
	DeliveryCity           string `json:"DeliveryCity"`
	AddressDeliveryAllowed bool   `json:"AddressDeliveryAllowed"`
	StreetsAvailability    bool   `json:"StreetsAvailability"`
	ParentRegionTypes      string `json:"ParentRegionTypes"`
	ParentRegionCode       string `json:"ParentRegionCode"`
	RegionTypes            string `json:"RegionTypes"`
	RegionTypesCode        string `json:"RegionTypesCode"`
}
