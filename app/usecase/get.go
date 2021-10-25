package usecase

import (
	"challange-20211024/app/entities"
	"challange-20211024/app/repository"
	"encoding/json"
	"errors"
)

type DataGetter interface {
	GetData(payload entities.Get) (interface{}, error)
}

type get struct{}

func NewDataGetter() DataGetter {
	return &get{}
}

func (g *get) GetData(payload entities.Get) (interface{}, error) {
	var results entities.Response
	page := payload.Page
	if page == 0 {
		page = 1
	}

	response, err := repository.NewData("http://www.omdbapi.com", "faf7e5bb&s").Get(payload.Search, page)
	if err != nil {
		return nil, err
	}

	respDecode, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respDecode, &results)
	if err != nil {
		return nil, err
	}

	if results.Error != "" {
		return nil, errors.New(results.Error)
	}

	return results, nil
}
