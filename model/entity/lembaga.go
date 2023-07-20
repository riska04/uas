package entity

import "materi/helper"

type LembagaEntity struct {
	IdLembaga        string            `json:"id_lembaga"`
	NamaLembaga      string            `json:"nama_lembaga"`
	SingkatanLembaga helper.NullString `json:"singkatan_lembaga"`
}
