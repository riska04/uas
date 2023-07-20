package response

import "materi/helper"

type StrukturResponse struct {
	IdStruktur   helper.NullString `json:"id_struktur"`
	NamaStruktur helper.NullString `json:"nama_struktur"`
}
