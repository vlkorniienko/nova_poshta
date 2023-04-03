package novaposhta

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vlkorniienko/novaposhta/pkg/build"
	"github.com/vlkorniienko/novaposhta/pkg/common/http/response"
	"github.com/vlkorniienko/novaposhta/pkg/novaposhta"
)

type Config struct {
	APIKey string `config:"api_key"`
}

type Service struct {
	np *novaposhta.Client
}

func NewService(apiKey string) *Service {
	var s = &Service{
		np: novaposhta.NewClient(apiKey),
	}

	return s
}

func (s *Service) RegisterRoutes(r build.HTTPRouter) error {
	r.HandleFunc("/novaposhta/areas", s.GetAreas).
		Methods(http.MethodGet)
	r.HandleFunc("/novaposhta/cities/online", s.GetCitiesOnline).
		Methods(http.MethodGet)
	r.HandleFunc("/novaposhta/warehouses", s.GetWarehouses).
		Methods(http.MethodGet)

	return nil
}

func (s *Service) GetAreas(w http.ResponseWriter, _ *http.Request) {
	resp, err := s.np.GetAreas()
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError,
			fmt.Errorf("can't make get areas request"), "can't show nova poshta areas")

		return
	}

	w.WriteHeader(http.StatusAccepted)

	functionResponse := map[string]interface{}{"data": resp.Data}
	err = json.NewEncoder(w).Encode(functionResponse)
	if err != nil {
		panic(err.Error())
	}
}

type CityRequest struct {
	City  string `json:"city"`
	Limit int    `json:"limit"`
}

func (s *Service) GetCitiesOnline(w http.ResponseWriter, r *http.Request) {
	var cityInfo CityRequest
	err := json.NewDecoder(r.Body).Decode(&cityInfo)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest,
			fmt.Errorf("can't desirialize cities request"), "bad request")

		return
	}

	resp, err := s.np.GetCitiesOnline(cityInfo.City, cityInfo.Limit)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError,
			fmt.Errorf("can't make cities request"), "can't show nova poshta cities")

		return
	}

	w.WriteHeader(http.StatusAccepted)

	functionResponse := map[string]interface{}{"data": resp}
	err = json.NewEncoder(w).Encode(functionResponse)
	if err != nil {
		panic(err.Error())
	}
}

type WarehousesRequest struct {
	CityRef string `json:"city_ref"`
	Limit   int    `json:"limit"`
}

func (s *Service) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	var wareHouseInfo WarehousesRequest
	err := json.NewDecoder(r.Body).Decode(&wareHouseInfo)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest,
			fmt.Errorf("can't desirialize warehouses request"), "bad request")

		return
	}

	resp, err := s.np.GetWarehouses(wareHouseInfo.CityRef, wareHouseInfo.Limit)
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError,
			fmt.Errorf("can't make warehouses request"), "can't show nova poshta warehouses")

		return
	}

	w.WriteHeader(http.StatusAccepted)

	functionResponse := map[string]interface{}{"data": resp}
	err = json.NewEncoder(w).Encode(functionResponse)
	if err != nil {
		panic(err.Error())
	}
}
