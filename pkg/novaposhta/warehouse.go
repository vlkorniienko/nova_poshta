package novaposhta

type WareHouseResponse struct {
	Success bool `json:"success"`
	Data    []struct {
		SiteKey                        string `json:"SiteKey"`
		Description                    string `json:"Description"`
		ShortAddress                   string `json:"ShortAddress"`
		Phone                          string `json:"Phone"`
		TypeOfWarehouse                string `json:"TypeOfWarehouse"`
		Ref                            string `json:"Ref"`
		Number                         string `json:"Number"`
		CityRef                        string `json:"CityRef"`
		CityDescription                string `json:"CityDescription"`
		SettlementRef                  string `json:"SettlementRef"`
		SettlementDescription          string `json:"SettlementDescription"`
		SettlementAreaDescription      string `json:"SettlementAreaDescription"`
		SettlementRegionsDescription   string `json:"SettlementRegionsDescription"`
		SettlementTypeDescription      string `json:"SettlementTypeDescription"`
		PostFinance                    string `json:"PostFinance"`
		PaymentAccess                  string `json:"PaymentAccess"`
		POSTerminal                    string `json:"POSTerminal"`
		InternationalShipping          string `json:"InternationalShipping"`
		SelfServiceWorkplacesCount     string `json:"SelfServiceWorkplacesCount"`
		TotalMaxWeightAllowed          string `json:"TotalMaxWeightAllowed"`
		PlaceMaxWeightAllowed          string `json:"PlaceMaxWeightAllowed"`
		SendingLimitationsOnDimensions struct {
			Width  int `json:"Width"`
			Height int `json:"Height"`
			Length int `json:"Length"`
		} `json:"SendingLimitationsOnDimensions"`
		ReceivingLimitationsOnDimensions struct {
			Width  int `json:"Width"`
			Height int `json:"Height"`
			Length int `json:"Length"`
		} `json:"ReceivingLimitationsOnDimensions"`
		Reception struct {
			Monday    string `json:"Monday"`
			Tuesday   string `json:"Tuesday"`
			Wednesday string `json:"Wednesday"`
			Thursday  string `json:"Thursday"`
			Friday    string `json:"Friday"`
			Saturday  string `json:"Saturday"`
			Sunday    string `json:"Sunday"`
		} `json:"Reception"`
		Delivery struct {
			Monday    string `json:"Monday"`
			Tuesday   string `json:"Tuesday"`
			Wednesday string `json:"Wednesday"`
			Thursday  string `json:"Thursday"`
			Friday    string `json:"Friday"`
			Saturday  string `json:"Saturday"`
			Sunday    string `json:"Sunday"`
		} `json:"Delivery"`
		Schedule struct {
			Monday    string `json:"Monday"`
			Tuesday   string `json:"Tuesday"`
			Wednesday string `json:"Wednesday"`
			Thursday  string `json:"Thursday"`
			Friday    string `json:"Friday"`
			Saturday  string `json:"Saturday"`
			Sunday    string `json:"Sunday"`
		} `json:"Schedule"`
		DistrictCode        string `json:"DistrictCode"`
		WarehouseStatus     string `json:"WarehouseStatus"`
		WarehouseStatusDate string `json:"WarehouseStatusDate"`
		CategoryOfWarehouse string `json:"CategoryOfWarehouse"`
		Direct              string `json:"Direct"`
		RegionCity          string `json:"RegionCity"`
		WarehouseForAgent   string `json:"WarehouseForAgent"`
		PostomatFor         string `json:"PostomatFor,omitempty"`
	} `json:"data"`
	Errors   []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
	Info     struct {
		TotalCount int `json:"totalCount"`
	} `json:"info"`
	MessageCodes []interface{} `json:"messageCodes"`
	ErrorCodes   []interface{} `json:"errorCodes"`
	WarningCodes []interface{} `json:"warningCodes"`
	InfoCodes    []interface{} `json:"infoCodes"`
}