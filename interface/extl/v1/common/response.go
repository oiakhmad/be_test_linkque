package common

import (
	"math"
	"net/http"
)

type ResponseMeta struct {
	CurrentPage int `json:"page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	LastPage    int `json:"last_page"`
}

type Response struct {
	Error   interface{}   `json:"error,omitempty"`
	Message string        `json:"message" example:"message"`
	Data    interface{}   `json:"data"`
	Meta    *ResponseMeta `json:"meta,omitempty"`
}

type GeneralResponse struct {
	Status GeneralStatus `json:"status"`
	Data   interface{}   `json:"data,omitempty"`
}

type GeneralStatus struct {
	Code    string      `json:"code"`
	Message string      `json:"message" example:"message"`
	Error   interface{} `json:"error,omitempty"`
}

func NewCreated(data interface{}, message string) *Response {
	return NewResponse(data, message, http.StatusCreated, nil, nil)
}

func NewOK(data interface{}, message string) *Response {
	return NewResponse(data, message, http.StatusOK, nil, nil)
}

func NewOKPaginated(data interface{}, message string, meta *ResponseMeta) *Response {
	return NewResponse(data, message, http.StatusOK, nil, meta)
}

func NewBadRequest(err interface{}, message string) interface{} {
	return NewResponse(nil, message, http.StatusBadRequest, err, nil)
}

func NewInternalServerError(err interface{}, message string) interface{} {
	return NewResponse(nil, message, http.StatusInternalServerError, err, nil)
}

func NewForbidden(err interface{}, message string) interface{} {
	return NewResponse(nil, message, http.StatusForbidden, err, nil)
}

func NewNotFound(err interface{}, message string) interface{} {
	return NewResponse(nil, message, http.StatusNotFound, err, nil)
}

func NewUnauthorized(err interface{}, message string) interface{} {
	return NewResponse(nil, message, http.StatusUnauthorized, err, nil)
}

func NewUnprocessableEntity(err interface{}, message string) interface{} {
	return NewResponse(nil, message, http.StatusUnprocessableEntity, err, nil)
}

func NewResponse(data interface{}, message string, statusCode int, err interface{}, meta *ResponseMeta) *Response {
	response := &Response{
		Error:   err,
		Message: http.StatusText(statusCode),
		Data:    data,
		Meta:    meta,
	}
	if message != "" {
		response.Message = message
	}
	return response
}

func GetPagination(count int, page int, perPage int) *ResponseMeta {
	var to int
	var defaultLimit int
	if page <= 0 {
		page = 1
	}

	allPage := int(math.Ceil(float64(count) / float64(perPage)))

	if page == allPage {
		to = count
	} else {
		to = int(perPage * page)
	}

	from := int((perPage * (page - 1)) + 1)

	if page <= 0 {
		page = 1
		from = 1
		to = count
		perPage = defaultLimit
		allPage = perPage
	}

	meta := new(ResponseMeta)
	meta.CurrentPage = page
	meta.Total = count
	meta.PerPage = perPage
	meta.From = from
	meta.To = to
	meta.LastPage = int(allPage)

	return meta
}

// return ((int) $this->page - 1) * (int) $this->limit;
func GetSkip(page int, perPage int) int {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * perPage
}
