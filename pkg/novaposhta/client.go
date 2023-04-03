package novaposhta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiEndpoint       = "https://api.novaposhta.ua/v2.0"
	requestFormatJSON = "json"
)

const (
	maxItemsShowedLimit = 200
	minItemsShowedLimit = 0
	itemsShowedLimit    = 30
)

type Client struct {
	httpClient    *http.Client
	apiKey        string
	apiEndpoint   string
	requestFormat string
}

func NewClient(apikey string) *Client {
	return &Client{
		httpClient:    makeHTTPClient(),
		apiKey:        apikey,
		apiEndpoint:   apiEndpoint,
		requestFormat: requestFormatJSON,
	}
}

func (c *Client) WithAPIEndpoint(apiEndpoint string) {
	c.apiEndpoint = apiEndpoint
}

func (c *Client) WithRequestFormat(requestFormat string) {
	c.requestFormat = requestFormat
}

func (c *Client) GetAreas() (*AreasResponse, error) {
	const calledMethod = "getAreas"

	model := c.generateGetAreasRequestModel(calledMethod)

	var areasResponse AreasResponse

	err := c.postJSON(c.makeGetAreasURL(calledMethod), model, &areasResponse)
	if err != nil {
		return nil, fmt.Errorf("can't make request to nova poshta: %w", err)
	}

	return &areasResponse, nil
}

func (c *Client) generateGetAreasRequestModel(calledMethod string) ModelAddressRequest {
	return ModelAddressRequest{
		APIKey:       c.apiKey,
		ModelName:    AddressModel,
		CalledMethod: calledMethod,
	}
}

func (c *Client) makeGetAreasURL(calledMethod string) string {
	return fmt.Sprintf("%s/%s/%s/%s", c.apiEndpoint, c.requestFormat, AddressModel, calledMethod)
}

func (c *Client) GetCitiesOnline(city string, limit int) (*CitiesResponse, error) {
	model := c.generateCitiesOnlineRequestModel(city, limit)

	var cities CitiesResponse

	err := c.postJSON(c.makeCitiesOnlineURL(), model, &cities)
	if err != nil {
		return nil, fmt.Errorf("can't make request to nova poshta: %w", err)
	}

	return &cities, nil
}

func (c *Client) generateCitiesOnlineRequestModel(city string, limit int) ModelAddressRequest {
	const calledMethod = "searchSettlements"

	return ModelAddressRequest{
		APIKey:       c.apiKey,
		ModelName:    AddressModel,
		CalledMethod: calledMethod,
		MethodProperties: MethodProperties{
			Limit:    getRequestItemsShowedLimit(limit),
			CityName: city,
		},
	}
}

func (c *Client) makeCitiesOnlineURL() string {
	return fmt.Sprintf("%s/%s/%s", c.apiEndpoint, c.requestFormat, "Address/getCities")
}

func getRequestItemsShowedLimit(limit int) int {
	if limit <= minItemsShowedLimit || limit >= maxItemsShowedLimit {
		return itemsShowedLimit
	}

	return limit
}

func (c *Client) GetWarehouses(cityRef string, limit int) (*WareHouseResponse, error) {
	const calledMethod = "getWarehouses"

	model := c.generateWarehousesRequestModel(cityRef, calledMethod, limit)

	var warehouses WareHouseResponse

	err := c.postJSON(c.makeGetWarehousesURL(calledMethod), model, &warehouses)
	if err != nil {
		return nil, fmt.Errorf("can't make request to nova poshta: %w", err)
	}

	return &warehouses, nil
}

func (c *Client) generateWarehousesRequestModel(cityRef, calledMethod string, limit int) ModelAddressRequest {
	return ModelAddressRequest{
		APIKey:       c.apiKey,
		ModelName:    AddressGeneralModel,
		CalledMethod: calledMethod,
		MethodProperties: MethodProperties{
			SettlementRef: cityRef,
			Limit:         getRequestItemsShowedLimit(limit),
		},
	}
}

func (c *Client) makeGetWarehousesURL(calledMethod string) string {
	return fmt.Sprintf("%s/%s/%s/%s", c.apiEndpoint, c.requestFormat, AddressGeneralModel, calledMethod)
}

const jsonMimeType = "application/json"

func (c *Client) postJSON(url string, request, response interface{}) error {
	payload, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("can't serialize request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("can't create new http request: %w", err)
	}

	req.Header.Set("Content-Type", jsonMimeType)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("can't make http request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("can't deserialize response: %w", err)
	}

	return nil
}

func makeHTTPClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return httpClient
}
