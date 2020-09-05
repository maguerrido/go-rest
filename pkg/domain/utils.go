// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package domain

import "reflect"

const (
	// general errors
	ErrUnexpected = "unexpected error"
	ErrNoContent  = "no content"
	ErrNotFound   = "not found"

	// reader constraints
	MinPage     = 1
	MaxElements = 100
	MinElements = 1

	// reader errors
	ErrElementsNotFound = "elements: not found"
	ErrElementsType     = "elements: must be numeric"

	ErrPageNotFound = "page: not found"
	ErrPageType     = "page: must be numeric"
)

type Reader struct {
	Elements, Page int
}

func (r *Reader) Validate() {
	if r.Elements < MinElements || r.Elements > MaxElements {
		r.Elements = MaxElements
	}
	if r.Page < MinPage {
		r.Page = MinPage
	}
}
func (r *Reader) NumField() int {
	return reflect.TypeOf(Reader{}).NumField()
}
