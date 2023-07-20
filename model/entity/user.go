package entity

import (
	"materi/helper"
)

type UserEntity struct {
	IdUser      string            `json:"id_user"`
	IdUnit      string            `json:"id_unit"`
	IdJabatan   string            `json:"id_jabatan"`
	NamaLengkap string            `json:"nama_lengkap"`
	Alamat      string            `json:"alamat"`
	NomorHp     string            `json:"nomer_hp"`
	UserName    string            `json:"username"`
	Password    string            `json:"password"`
	CreateDate  helper.NullString `json:"create_date"`
	CreateBy    helper.NullString `json:"create_by"`
	UpdateDate  helper.NullString `json:"update_date"`
	UpdateBy    helper.NullString `json:"update_by"`
	Status      string            `json:"status"`
}
