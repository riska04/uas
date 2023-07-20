package response

import "materi/helper"

type NotulenResponse struct {
	IdNotulen  helper.NullString `json:"id_notulen"`
	IdBerkas   helper.NullString `json:"id_berkas"`
	Catatan    helper.NullString `json:"catatan"`
	CreateDate helper.NullString `json:"create_date"`
	CreateBy   helper.NullString `json:"create_by"`
	UpdateDate helper.NullString `json:"update_date"`
	UpdateBy   helper.NullString `json:"update_by"`
}
