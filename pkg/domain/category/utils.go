// Copyright 2020 maguerrido <mauricio.aguerrido@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package category

import "errors"

func ValidateID(id int64) error {
	if !(id >= MinID && id <= MaxID) {
		return errors.New(ErrIDOutOfRange)
	}
	return nil
}

func ModelsToMap(categories []Model) []map[string]interface{} {
	slice := make([]map[string]interface{}, 0)
	for _, category := range categories {
		slice = append(slice, map[string]interface{}{
			"id":   category.ID,
			"name": category.Name,
		})
	}
	return slice
}
