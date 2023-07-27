package services

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"go-webapi-template/internal/domain/services"
	"net/http"
)

type DummyJsonProducts struct {
	Products []services.Product `json:"products"`
}

type ApiProductsServiceProvider struct {
	httpClient *resty.Client
}

func NewApiEmployeesServiceProvider(httpClient *resty.Client) services.IProductsServiceProvider {
	return &ApiProductsServiceProvider{
		httpClient: httpClient,
	}
}

func (a *ApiProductsServiceProvider) GetProducts() []services.Product {

	// ref. https://github.com/Ovi/DummyJSON/tree/master
	var err error
	var resp *resty.Response
	var body []byte

	resp, err = a.httpClient.R().Get("https://dummyjson.com/products")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return []services.Product{}
	}

	body = resp.Body()

	if err != nil {
		return []services.Product{}
	}

	products := DummyJsonProducts{}
	err = json.Unmarshal(body, &products)

	if err != nil {
		return []services.Product{}
	}

	return products.Products
}
