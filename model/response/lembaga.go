package response

import "materi/helper"

type LembagaResponse struct {
	IdLembaga        helper.NullString `json:"id_lembaga"`
	NamaLembaga      helper.NullString `json:"nama_lembaga"`
	SingkatanLembaga helper.NullString `json:"singkatan_lembaga"`
}
