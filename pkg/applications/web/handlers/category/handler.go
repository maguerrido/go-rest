// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package category

import (
	"encoding/json"
	"github.com/maguerrido/go-rest/pkg/applications/web/handlers"
	"github.com/maguerrido/go-rest/pkg/domain/category"
	"net/http"
)

const (
	errUnsupportedMethod = "unsupported method"
	errParseForm         = "could not read parameters"
	errContentType       = "content-type header must be application/x-www-form-urlencoded"
	errNumberParams      = "wrong number of parameters"
	errURLParams         = "url parameters not allowed"
)

func encodeJSON(code int, response *handlers.Response, writer *http.ResponseWriter) {
	(*writer).Header().Set("Content-Type", "application/json")
	(*writer).WriteHeader(code)
	_ = json.NewEncoder(*writer).Encode(response)
}

func Handler(service category.Services) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request.Body = http.MaxBytesReader(writer, request.Body, 1048576)

		if err := request.ParseForm(); err != nil {
			code, response := http.StatusBadRequest, handlers.SetErrResponse(errParseForm)
			encodeJSON(code, response, &writer)
			return
		}

		if request.Method == http.MethodGet {
			code, response := getMethod(service, request.Form)
			encodeJSON(code, response, &writer)
			return
		}

		if request.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			code, response := http.StatusBadRequest, handlers.SetErrResponse(errContentType)
			encodeJSON(code, response, &writer)
			return
		}

		if len(request.URL.Query()) != 0 {
			code, response := http.StatusBadRequest, handlers.SetErrResponse(errURLParams)
			encodeJSON(code, response, &writer)
			return
		}

		switch request.Method {
		case http.MethodPost:
			code, response := postMethod(service, request.PostForm)
			encodeJSON(code, response, &writer)
			return
		case http.MethodPut:
			code, response := putMethod(service, request.PostForm)
			encodeJSON(code, response, &writer)
			return
		case http.MethodDelete:
			code, response := deleteMethod(service, request.PostForm)
			encodeJSON(code, response, &writer)
			return
		default:
			code, response := http.StatusBadRequest, handlers.SetErrResponse(errUnsupportedMethod)
			encodeJSON(code, response, &writer)
			return
		}
	})
}
