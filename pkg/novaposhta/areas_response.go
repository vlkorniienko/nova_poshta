package novaposhta

type AreasResponse struct {
	Success    bool          `json:"success"`
	Data       []Data        `json:"data"`
	Errors     []interface{} `json:"errors"`
	ErrorCodes []interface{} `json:"errorCodes"`
}

type Data struct {
	Ref         string `json:"Ref"`
	AreasCenter string `json:"AreasCenter"`
	Description string `json:"Description"`
}
