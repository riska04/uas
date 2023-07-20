package response

import "materi/helper"

type UndanganResponse struct {
	IdUndangan helper.NullString `json:"id_undangan"`
	IdUnit     helper.NullString `json:"id_unit"`
	IdJabatan  helper.NullString `json:"id_jabatan"`
	IdUser     helper.NullString `json:"id_user"`
	IdBerkas   helper.NullString `json:"id_berkas"`
}
