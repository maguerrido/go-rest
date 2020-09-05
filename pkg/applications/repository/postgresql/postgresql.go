// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/maguerrido/go-rest/pkg/applications/repository"
)

func Repository(host, port, user, password, dbname, sslmode string) (*repository.App, error) {
	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &repository.App{
		CategoryServices: &CategoryServices{db: db},
	}, nil
}
