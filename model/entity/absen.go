package entity

import (
	"materi/helper"
)

type AbsenEntity struct {
	IdAbsen   string            `json:"id_absen"`
	IdBerkas  string            `json:"id_berkas"`
	IdUnit    string            `json:"id_unit"`
	IdJabatan string            `json:"id_jabatan"`
	IdUser    string            `json:"id_user"`
	CreatedAt helper.NullString `json:"created_at"`
	UpdatedAt helper.NullString `json:"update_at"`
}
