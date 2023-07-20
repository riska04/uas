package request

import (
	"materi/helper"

	"github.com/go-playground/validator/v10"
)

type LembagaCreate struct {
	NamaLembaga string `json:"nama_lembaga" form:"nama_lembaga" validate:"required"`
}

type LembagaUpdate struct {
	IdLembaga        string            `json:"id_lembaga"`
	NamaLembaga      string            `json:"nama_lembaga" form:"nama_lembaga" validate:"required"`
	SingkatanLembaga helper.NullString `json:"singkatan_lembaga" form:"singkatan_lembaga" validate:"required"`
}

type LembagaProfile struct {
	IdLembaga string `json:"id_lembaga"`
}

type ErrorResponse4 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct4(validasi interface{}) []*ErrorResponse4 {
	var errors []*ErrorResponse4
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse4
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
