/*
 * Created on 09/10/22 15.57
 *
 * Copyright (c) 2022 Abdul Ghani Abbasi
 */

package model

import "time"

type Feeder struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Uuid      string    `json:"uuid" gorm:"unique"`
	Barcode   string    `json:"barcode"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
