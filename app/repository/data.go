package repository

import (
	"encoding/json"
	"net/http"
)

type data struct {
	URI string
	Key string
}

func NewData(uri, key string) Data {
	return &data{
		URI: uri,
		Key: key,
	}
}

func (d *data) Get(words string, page int) (interface{}, error) {
	var client = &http.Client{}
	var data interface{}

	path := "/?apikey=" + d.Key + "&s=" + words + "&page="
	request, err := http.NewRequest(http.MethodGet, d.URI+path, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
