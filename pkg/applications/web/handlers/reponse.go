// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package handlers

import (
	"github.com/maguerrido/go-rest/pkg/domain/category"
)

type Response struct {
	Message string                   `json:"message"`
	Ok      bool                     `json:"ok"`
	Data    []map[string]interface{} `json:"data"`
}

func SetOKResponse(message string, data []category.Model) *Response {
	return &Response{
		Message: message,
		Ok:      true,
		Data:    category.ModelsToMap(data),
	}
}

func SetErrResponse(message string) *Response {
	return &Response{
		Message: message,
		Ok:      false,
		Data:    nil,
	}
}
