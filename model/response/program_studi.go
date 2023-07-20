package response

import "materi/helper"

type ProgramStudiResponse struct {
	ProdiId      helper.NullString `json:"prodi_id"`
	IdFakultas   helper.NullString `json:"id_fakultas"`
	ProgramStudi helper.NullString `json:"program_studi"`
	Singkatan    helper.NullString `json:"singatan"`
}
