package novaposhta

const (
	AddressModel        = "Address"
	AddressGeneralModel = "AddressGeneral"
)

type ModelAddressRequest struct {
	APIKey           string           `json:"apiKey"`
	ModelName        string           `json:"modelName"`
	CalledMethod     string           `json:"calledMethod"`
	MethodProperties MethodProperties `json:"methodProperties"`
}

type MethodProperties struct {
	CityName      string `json:"CityName"`
	Limit         int    `json:"Limit"`
	Language      string `json:"Language"`
	SettlementRef string `json:"SettlementRef"`
}
