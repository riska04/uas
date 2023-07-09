package request

import (
	"github.com/go-playground/validator/v10"
)

type JabatanCreate struct {
	NamaJabatan string `json:"nama_jabatan" form:"nama_jabatan" validate:"required"`
}

type JabatanUpdate struct {
	IdJabatan   string `json:"id_jabatan"`
	NamaJabatan string `json:"nama_jabatan" form:"nama_jabatan" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(validasi interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
