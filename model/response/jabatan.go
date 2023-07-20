package response

import "materi/helper"

type JabatanResponse struct {
	IdJabatan   helper.NullString `json:"id_jabatan"`
	NamaJabatan helper.NullString `json:"nama_jabatan"`
}
