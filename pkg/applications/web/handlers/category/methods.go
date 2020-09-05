// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package category

import (
	"errors"
	"github.com/maguerrido/go-rest/pkg/applications/web/handlers"
	"github.com/maguerrido/go-rest/pkg/domain"
	"github.com/maguerrido/go-rest/pkg/domain/category"
	"net/http"
	"net/url"
	"strconv"
)

func checkIDParams(params url.Values) (int64, error) {
	if len(params) != 1 {
		return 0, errors.New(errNumberParams)
	}

	idStr := params.Get("id")
	if idStr == "" {
		return 0, errors.New(category.ErrIDNotFound)
	}

	idInt64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, errors.New(category.ErrIDType)
	}

	return idInt64, nil
}
func checkParamsToCreate(params url.Values) (*category.Creator, error) {
	creator := &category.Creator{}
	if len(params) != creator.NumField() {
		return nil, errors.New(errNumberParams)
	}

	name := params.Get("name")
	if name == "" {
		return nil, errors.New(category.ErrNameNotFound)
	}
	creator.Name = name

	return creator, nil
}
func checkParamsToReadAll(params url.Values) (*domain.Reader, error) {
	reader := &domain.Reader{}
	if len(params) != reader.NumField() {
		return nil, errors.New(errNumberParams)
	}

	elementsStr := params.Get("elements")
	if elementsStr == "" {
		return nil, errors.New(domain.ErrElementsNotFound)
	}
	elements, err := strconv.Atoi(elementsStr)
	if err != nil {
		return nil, errors.New(domain.ErrElementsType)
	}

	pageStr := params.Get("page")
	if pageStr == "" {
		return nil, errors.New(domain.ErrPageNotFound)
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, errors.New(domain.ErrPageType)
	}

	reader.Elements = elements
	reader.Page = page
	return reader, nil
}
func checkParamsToUpdate(params url.Values) (*category.Updater, error) {
	updater := &category.Updater{}
	if len(params) != updater.NumField() {
		return nil, errors.New(errNumberParams)
	}

	idStr := params.Get("id")
	if idStr == "" {
		return nil, errors.New(category.ErrIDNotFound)
	}
	idInt64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errors.New(category.ErrIDType)
	}
	updater.ID = idInt64

	name := params.Get("name")
	if name == "" {
		return nil, errors.New(category.ErrNameNotFound)
	}
	updater.Name = name

	return updater, nil
}

func deleteMethod(service category.Services, params url.Values) (int, *handlers.Response) {
	id, err := checkIDParams(params)
	if err != nil {
		return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
	}

	deleted, err := service.Delete(id)
	if err != nil {
		switch err.Error() {
		case category.ErrIDOutOfRange:
			return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
		case domain.ErrNotFound:
			return http.StatusNotFound, handlers.SetErrResponse(err.Error())
		default: // category.ErrUnexpected
			return http.StatusInternalServerError, handlers.SetErrResponse(err.Error())
		}
	}

	return http.StatusOK, handlers.SetOKResponse("ok", []category.Model{*deleted})
}
func getMethod(service category.Services, params url.Values) (int, *handlers.Response) {
	if id := params.Get("id"); id == "" { // case: ReadAll
		reader, err := checkParamsToReadAll(params)
		if err != nil {
			return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
		}

		categories, err := service.ReadAll(*reader)
		if err != nil {
			switch err.Error() {
			case domain.ErrNoContent:
				return http.StatusOK, handlers.SetOKResponse(err.Error(), categories)
			default: // category.ErrUnexpected
				return http.StatusInternalServerError, handlers.SetErrResponse(err.Error())
			}
		}

		return http.StatusOK, handlers.SetOKResponse("ok", categories)

	} else { // case: Read
		id, err := checkIDParams(params)
		if err != nil {
			return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
		}

		cat, err := service.Read(id)
		if err != nil {
			switch err.Error() {
			case category.ErrIDOutOfRange:
				return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
			case domain.ErrNotFound:
				return http.StatusNotFound, handlers.SetErrResponse(err.Error())
			default: // category.ErrUnexpected
				return http.StatusInternalServerError, handlers.SetErrResponse(err.Error())
			}
		}

		return http.StatusOK, handlers.SetOKResponse("ok", []category.Model{*cat})
	}
}
func postMethod(service category.Services, params url.Values) (int, *handlers.Response) {
	creator, err := checkParamsToCreate(params)
	if err != nil {
		return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
	}

	created, err := service.Create(*creator)
	if err != nil {
		switch err.Error() {
		case category.ErrNameOutOfRange, category.ErrNameInUse:
			return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
		default: // category.ErrUnexpected
			return http.StatusInternalServerError, handlers.SetErrResponse(err.Error())
		}
	}

	return http.StatusOK, handlers.SetOKResponse("ok", []category.Model{*created})
}
func putMethod(service category.Services, params url.Values) (int, *handlers.Response) {
	updater, err := checkParamsToUpdate(params)
	if err != nil {
		return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
	}

	updated, err := service.Update(*updater)
	if err != nil {
		switch err.Error() {
		case category.ErrIDOutOfRange, category.ErrNameOutOfRange, category.ErrNameInUse:
			return http.StatusBadRequest, handlers.SetErrResponse(err.Error())
		case domain.ErrNotFound:
			return http.StatusNotFound, handlers.SetErrResponse(err.Error())
		default: // category.ErrUnexpected
			return http.StatusInternalServerError, handlers.SetErrResponse(err.Error())
		}
	}

	return http.StatusOK, handlers.SetOKResponse("ok", []category.Model{*updated})
}
