// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package main

import (
	"github.com/maguerrido/go-rest/pkg/applications/repository"
	"github.com/maguerrido/go-rest/pkg/applications/repository/postgresql"
	"github.com/maguerrido/go-rest/pkg/applications/web"
	"log"
)

func main() {
	var err error
	var repo *repository.App
	var app *web.App

	// TODO: env variables
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "1234"
	dbname := "gorest"
	sslmode := "disable"

	repo, err = postgresql.Repository(host, port, user, password, dbname, sslmode)
	if err != nil {
		log.Fatal(err)
	} else {
		app = web.New(repo)
		app.Run()
	}
}
