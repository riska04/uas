package response

import "materi/helper"

type AbsenResponse struct {
	IdAbsen    helper.NullString `json:"id_absen" form:"id_absen"`
	IdBerkas   helper.NullString `json:"id_berkas" form:"id_berkas" validate:"required"`
	IdUnit     helper.NullString `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan  helper.NullString `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	IdUser     helper.NullString `json:"id_user" form:"id_user" validate:"required"`
	CreateDate helper.NullString `json:"create_date" form:"create_date" validate:"required"`
	UpdateDate helper.NullString `json:"update_date" form:"update_date" validate:"required"`
}
