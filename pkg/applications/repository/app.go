// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package repository

import (
	"github.com/maguerrido/go-rest/pkg/domain/category"
)

type App struct {
	CategoryServices category.Services
}
