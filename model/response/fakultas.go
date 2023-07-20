package response

import "materi/helper"

type FakultasResponse struct {
	IdFakultas        helper.NullString `json:"id_fakultas"`
	NamaFakultas      helper.NullString `json:"nama_fakultas"`
	SingkatanFakultas helper.NullString `json:"singkatan_fakultas"`
}
