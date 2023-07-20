package response

import "materi/helper"

type UserResponse struct {
	IdUser      int64             `json:"id_user"`
	IdUnit      string            `json:"id_unit"`
	IdJabatan   string            `json:"id_jabatan"`
	NamaLengkap string            `json:"nama_lengkap"`
	Alamat      string            `json:"alamat"`
	NomorHp     string            `json:"Nomer_hp"`
	UserName    string            `json:"Username"`
	Password    string            `json:"password"`
	CreateDate  string            `json:"create_date"`
	CreateBy    helper.NullString `json:"create_by"`
	UpdateDate  string            `json:"update_date"`
	UpdateBy    helper.NullString `json:"update_by"`
	Status      string            `json:"status"`
}
