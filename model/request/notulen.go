package request

import (
	"github.com/go-playground/validator/v10"
)

type NotulenCreate struct {
	IdBerkas string `json:"id_berkas" form:"id_berkas" validate:"required"`
	Catatan  string `json:"catatan" form:"catatan" validate:"required"`
}

type NotulenUpdate struct {
	IdNotulen string `json:"id_notulen"`
	IdBerkas  string `json:"id_berkas" form:"id_berkas" validate:"required"`
	Catatan   string `json:"catatan" form:"catatan" validate:"required"`
}

type NotulenProfile struct {
	IdNotulen string `json:"id_notulen"`
}

type ErrorResponse5 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct5(validasi interface{}) []*ErrorResponse5 {
	var errors []*ErrorResponse5
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse5
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
