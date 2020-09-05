// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package category

import "github.com/maguerrido/go-rest/pkg/domain"

type Services interface {
	Create(creator Creator) (*Model, error)
	Read(id int64) (*Model, error)
	ReadAll(reader domain.Reader) ([]Model, error)
	Update(cat Updater) (*Model, error)
	Delete(id int64) (*Model, error)
}
