package entities

import (
	"errors"
	"net/http"
	"strconv"
)

type Get struct {
	Search string `json:"search"`
	Page   int    `json:"page"`
}

func NewGet(r *http.Request) (Get, error) {
	s, ok := r.URL.Query()["s"]
	if !ok || len(s[0]) < 1 {
		return Get{}, errors.New("Url Param  is missing")
	}

	p, ok := r.URL.Query()["page"]
	if !ok || len(p[0]) < 1 {
		return Get{}, errors.New("Url Param  is missing")
	}

	pInt, _ := strconv.Atoi(p[0])

	return Get{
		Search: s[0],
		Page:   pInt,
	}, nil
}

type Response struct {
	Response     string      `json:"Response"`
	Search       interface{} `json:"Search"`
	TotalResults string      `json:"totalResults"`
	Error        string      `json:"Error"`
}
