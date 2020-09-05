// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package web

import (
	"github.com/maguerrido/go-rest/pkg/applications/repository"
	"log"
	"net/http"
)

func New(repo *repository.App) *App {
	app := &App{
		mux:  http.NewServeMux(),
		repo: repo,
	}
	
	return app
}

type App struct {
	mux  *http.ServeMux
	repo *repository.App
}

func (a *App) Run() {
	log.Println("Running web app...")
	err := http.ListenAndServe(":8080", a.mux)
	log.Fatal(err)
}
