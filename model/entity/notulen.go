package entity

import "materi/helper"

type NotulenEntity struct {
	IdNotulen  string            `json:"id_notulen"`
	IdBerkas   string            `json:"id_berkas"`
	Catatan    string            `json:"catatan"`
	CreateDate helper.NullString `json:"create_date"`
	CreateBy   helper.NullString `json:"create_by"`
	UpdateDate helper.NullString `json:"update_date"`
	UpdateBy   helper.NullString `json:"update_by"`
}
