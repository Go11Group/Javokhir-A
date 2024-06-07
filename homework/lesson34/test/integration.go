package test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/internal/app/models"
)

func ReadingUsersJson() ([]*models.User, error) {
	var users []*models.User

	data, err := os.Open("./testdata/usersmockdata.json")
	if err != nil {
		return nil, fmt.Errorf("failed while opening json file: %v", err)
	}
	defer data.Close()
	decoder := json.NewDecoder(data)
	if err := decoder.Decode(&users); err != nil {
		return nil, fmt.Errorf("failed while decoding json into slice: %v", err)
	}

	return users, nil
}

func ReadingProductJson() ([]*models.Product, error) {
	var products []*models.Product

	data, err := os.Open("./testdata/productmockdata.json")
	if err != nil {
		return nil, fmt.Errorf("failed while opening json file: %v", err)
	}
	defer data.Close()
	decoder := json.NewDecoder(data)
	if err := decoder.Decode(&products); err != nil {
		return nil, fmt.Errorf("failed while decoding json into slice: %v", err)
	}

	return products, nil
}
