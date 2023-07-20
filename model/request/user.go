package request

import (
	"github.com/go-playground/validator/v10"
)

type UserCreate struct {
	IdUnit      string `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan   string `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	NamaLengkap string `json:"nama_lengkap" form:"nama_lengkap" validate:"required"`
	Alamat      string `json:"alamat" form:"alamat" validate:"required"`
	NomorHp     string `json:"nomor_hp" form:"nomor_hp" validate:"required"`
	UserName    string `json:"user_name" form:"user_name" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
}

type UserUpdate struct {
	IdUser      string `json:"id_user"`
	IdUnit      string `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan   string `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	NamaLengkap string `json:"nama_lengkap" form:"nama_lengkap" validate:"required"`
	NomorHp     string `json:"nomor_hp" form:"nomor_hp" validate:"required"`
	Alamat      string `json:"alamat" form:"alamat" validate:"required"`
	Username    string `json:"user_name" form:"user_name" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	// UdateDate    string `json:"update_date" form:"update_date" validate:"required"`
	// UdateBy      string `json:"update_by" form:"update_by" validate:"required"`
	Status string `json:"status" form:"status" validate:"required"`
}

type UserProfile struct {
	IdUser string `json:"id_user"`
}

type ErrorResponse10 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct10(validasi interface{}) []*ErrorResponse10 {
	var errors []*ErrorResponse10
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse10
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
