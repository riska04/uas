package request

import "github.com/go-playground/validator/v10"

type AbsenCreate struct {
	IdBerkas  string `json:"id_berkas" form:"id_berkas" validate:"required"`
	IdUnit    string `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan string `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	IdUser    string `json:"id_user" form:"id_user" validate:"required"`
}

type AbsenUpdate struct {
	IdAbsen   string `json:"id_absen"`
	IdBerkas  string `json:"id_berkas" form:"id_berkas" validate:"required"`
	IdUnit    string `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan string `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	IdUser    string `json:"id_user" form:"id_user" validate:"required"`
	//CreateDate string `json:"create_date" form:"create_date" validate:"required"`
	//UpdateDate string `json:"update_date" form:"update_date" validate:"required"`
}

type AbsenProfile struct {
	IdAbsen string `json:"id_absen"`
}

type ErrorResponse3 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct3(validasi interface{}) []*ErrorResponse3 {
	var errors []*ErrorResponse3
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse3
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
