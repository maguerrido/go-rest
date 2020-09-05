// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package category

import (
	"errors"
	"reflect"
)

const (
	// model constraints
	MaxID   = 32767
	MinID   = 1
	MaxName = 25
	MinName = 1

	// model errors
	ErrIDNotFound   = "id: not found"
	ErrIDType       = "id: must be numeric"
	ErrIDOutOfRange = "id: out of range"
	ErrIDInUse      = "id: in use"

	ErrNameNotFound   = "name: not found"
	ErrNameType       = "name: must be text"
	ErrNameOutOfRange = "name: out of range"
	ErrNameInUse      = "name: in use"
)

type Model struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (m *Model) Validate() error {
	if err := ValidateID(m.ID); err != nil {
		return err
	}
	if !(len(m.Name) >= MinName && len(m.Name) <= MaxName) {
		return errors.New(ErrNameOutOfRange)
	}
	return nil
}
func (m *Model) NumField() int {
	return reflect.TypeOf(Model{}).NumField()
}

type Creator struct {
	Name string
}

func (c *Creator) Validate() error {
	if !(len(c.Name) >= MinName && len(c.Name) <= MaxName) {
		return errors.New(ErrNameOutOfRange)
	}
	return nil
}
func (c *Creator) NumField() int {
	return reflect.TypeOf(Creator{}).NumField()
}

type Updater struct {
	ID   int64
	Name string
}

func (u *Updater) Validate() error {
	if !(u.ID >= MinID && u.ID <= MaxID) {
		return errors.New(ErrIDOutOfRange)
	}
	if !(len(u.Name) >= MinName && len(u.Name) <= MaxName) {
		return errors.New(ErrNameOutOfRange)
	}
	return nil
}
func (u *Updater) NumField() int {
	return reflect.TypeOf(Updater{}).NumField()
}
