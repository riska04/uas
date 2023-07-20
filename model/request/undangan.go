package request

import "github.com/go-playground/validator/v10"

type UndanganCreate struct {
	IdUnit    string `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan string `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	IdUser    string `json:"id_user" form:"id_user" validate:"required"`
	IdBerkas  string `json:"id_berkas" form:"id_berkas" validate:"required"`
}

type UndanganUpdate struct {
	IdUndangan string `json:"id_undangan"`
	IdUnit     string `json:"id_unit" form:"id_unit" validate:"required"`
	IdJabatan  string `json:"id_jabatan" form:"id_jabatan" validate:"required"`
	IdUser     string `json:"id_user" form:"id_user" validate:"required"`
	IdBerkas   string `json:"id_berkas" form:"id_berkas" validate:"required"`
}

type ErrorResponse8 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct8(validasi interface{}) []*ErrorResponse8 {
	var errors []*ErrorResponse8
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse8
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
