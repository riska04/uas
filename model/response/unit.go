package response

import "materi/helper"

type UnitResponse struct {
	IdUnit     helper.NullString `json:"id_unit"`
	IdLembaga  helper.NullString `json:"id_lembaga"`
	IdFakultas helper.NullString `json:"id_fakultas"`
	IdProdi    helper.NullString `json:"id_prodi"`
	IdStruktur helper.NullString `json:"id_struktur"`
	Status     helper.NullString `json:"status"`
}
