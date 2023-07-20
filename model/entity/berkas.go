package entity

import (
	"materi/helper"
)

type BerkasEntity struct {
	IdBerkas      string            `json:"id_berkas"`
	File          string            `json:"file"`
	NamaFile      string            `json:"nama_file"`
	TanggalUpload string            `json:"tanggal_upload"`
	TanggalAcara  string            `json:"tanggal_acara"`
	StartJam      string            `json:"start_jam"`
	SelesaiJam    string            `json:"selesai_jam"`
	IdUnit        string            `json:"id_unit"`
	CreateBy      helper.NullString `json:"create_by"`
	CreateDate    helper.NullString `json:"create_date"`
	UpdateBy      helper.NullString `json:"upate_by"`
	UpdateDate    helper.NullString `json:"update_date"`
}
