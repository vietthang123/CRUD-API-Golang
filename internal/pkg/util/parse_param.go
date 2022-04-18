package util

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type ParamRequest struct {
	FilterObject map[string]interface{}
	Limit        string
	Page         string
	Order        string
	Filter       string
	Error        error
}

func ParseToParam(values url.Values, r *http.Request) *ParamRequest {
	var paramReq = new(ParamRequest)
	m := make(map[string]interface{})
	var fltd []interface{}
	for key, _ := range values {
		if values.Get(key) == "undefined" {
			return paramReq
		}
		switch key {
		case "limit":
			paramReq.Limit = values.Get(key)
		case "page":
			paramReq.Page = values.Get(key)
		case "order":
			paramReq.Order = values.Get(key)
			objectMap := make(map[string]interface{})
			for order, value := range objectMap {
				if order == "" {
					order = "id.asc"
				} else if order == "sort" {
					m[order] = value
				}
				order = strings.Replace(order, ".", " ", -1)
			}
		case "filter":
			paramReq.Filter = values.Get(key)
			filter := make(map[string]interface{})
			json.Marshal(filter)
			for _, value := range filter {
				fltd = append(fltd, value)
			}
		default:
			m[key] = values.Get(key)
		}
	}
	paramReq.FilterObject = m
	return paramReq
}
