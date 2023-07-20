package entity

import "materi/helper"

type UnitEntity struct {
	IdUnit     string            `json:"id_unit"`
	IdLembaga  helper.NullString `json:"id_lembaga"`
	IdFakultas helper.NullString `json:"id_fakultas"`
	IdProdi    helper.NullString `json:"id_prodi"`
	IdStruktur string            `json:"id_struktur"`
	Status     string            `json:"status"`
}
