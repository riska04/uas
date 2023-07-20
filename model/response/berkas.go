package response

import "materi/helper"

type BerkasResponse struct {
	IdBerkas      helper.NullString `json:"id_berkas" form:"id_berkas"`
	File          helper.NullString `json:"file" form:"file" validate:"required"`
	NamaFile      helper.NullString `json:"nama_file" form:"nama_file" validate:"required"`
	TanggalUpload helper.NullString `json:"tanggal_upload" form:"tanggal_upload" validate:"required"`
	TanggalAcara  helper.NullString `json:"tanggal_acara" form:"tanggal_acara" validate:"required"`
	StartJam      helper.NullString `json:"start_jam" form:"start_jam" validate:"required"`
	SelesaiJam    helper.NullString `json:"selesai_jam" form:"selesai_jam" validate:"required"`
	IdUnit        helper.NullString `json:"id_unit" form:"id_unit" validate:"required"`
	CreateBy      helper.NullString `json:"create_by" form:"create_by" validate:"required"`
	CreateDate    helper.NullString `json:"create_date" form:"create_date" validate:"required"`
	UpdateBy      helper.NullString `json:"upate_by" form:"update_by" validate:"required"`
	UpdateDate    helper.NullString `json:"update_date" form:"update_date" validate:"required"`
}
