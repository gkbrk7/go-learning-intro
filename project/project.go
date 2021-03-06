package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


type product struct {
	ID int `json:"id"`
	ProductName string `json:"name"`
	CategoryID int `json:"categoryId"`
	UnitPrice float64 `json:"unitPrice"`
}

type category struct {
	ID int `json:"id"`
	CategoryName string `json:"name"`
}

func GetAllProducts() ([]product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err!= nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes,_ := ioutil.ReadAll(response.Body)
	var products []product
	json.Unmarshal(bodyBytes, &products)
	return products,nil
}

func AddProduct() (product, error) {
	// _product := product{4, "Phone", 1, 4000}
	_product := product{5, "Snacks", 2, 40}
	jsonProduct,err := json.Marshal(_product)
	if err != nil {
		return _product, err
	}
	
	response,_ := http.Post("http://localhost:3000/products","application/json;charset=utf-8",bytes.NewBuffer(jsonProduct))
	defer response.Body.Close()

	readBytes,_ := ioutil.ReadAll(response.Body)
	var productResponse product
	json.Unmarshal(readBytes,&productResponse)
	return productResponse,nil
}