// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package postgresql

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/maguerrido/go-rest/pkg/domain"
	"github.com/maguerrido/go-rest/pkg/domain/category"
)

type CategoryServices struct {
	db *sql.DB
}

func (s *CategoryServices) Create(creator category.Creator) (*category.Model, error) {
	if err := creator.Validate(); err != nil {
		return nil, err
	}

	cat := new(category.Model)
	query := "INSERT INTO categories(name) VALUES ($1) RETURNING id, name;"
	if err := s.db.QueryRow(query, creator.Name).Scan(&cat.ID, &cat.Name); err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				// if err.Column == "name" - commented because there is only one unique column
				return nil, errors.New(category.ErrNameInUse)
			}
		}
		return nil, errors.New(domain.ErrUnexpected)
	}

	return cat, nil
}
func (s *CategoryServices) Delete(id int64) (*category.Model, error) {
	if err := category.ValidateID(id); err != nil {
		return nil, err
	}

	cat := new(category.Model)
	query := "DELETE FROM categories WHERE id = $1 RETURNING id, name;"
	if err := s.db.QueryRow(query, cat.ID).Scan(&cat.ID, &cat.Name); err != nil {
		return nil, errors.New(domain.ErrUnexpected)
	}

	if err := category.ValidateID(cat.ID); err != nil {
		return nil, errors.New(domain.ErrNotFound)
	}

	return cat, nil
}
func (s *CategoryServices) Read(id int64) (*category.Model, error) {
	if err := category.ValidateID(id); err != nil {
		return nil, err
	}

	cat := new(category.Model)
	query := "SELECT id, name FROM categories WHERE id = $1;"
	if err := s.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(domain.ErrNotFound)
		}
		return nil, errors.New(domain.ErrUnexpected)
	}

	return cat, nil
}
func (s *CategoryServices) ReadAll(reader domain.Reader) ([]category.Model, error) {
	reader.Validate()

	query := "SELECT id, name FROM categories ORDER BY name LIMIT $1 OFFSET $2;"
	rows, err := s.db.Query(query, reader.Elements, reader.Elements*(reader.Page-1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]category.Model, 0)
	for rows.Next() {
		cat := new(category.Model)
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, errors.New(domain.ErrUnexpected)
		}
		categories = append(categories, *cat)
	}

	if err = rows.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(domain.ErrNoContent)
		}
		return nil, errors.New(domain.ErrUnexpected)
	}

	return categories, nil
}
func (s *CategoryServices) Update(updater category.Updater) (*category.Model, error) {
	if err := updater.Validate(); err != nil {
		return nil, err
	}

	cat := new(category.Model)
	query := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING id, name;"
	if err := s.db.QueryRow(query, cat.Name, cat.ID).Scan(&cat.ID, &cat.Name); err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			// if err.Column == "name" - commented because there is only one unique column
			return nil, errors.New(category.ErrNameInUse)
		}
		return nil, errors.New(domain.ErrUnexpected)
	}

	if err := category.ValidateID(cat.ID); err != nil {
		return nil, errors.New(domain.ErrNotFound)
	}

	return cat, nil
}
