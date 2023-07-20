package request

import (
	"github.com/go-playground/validator/v10"
)

type UnitCreate struct {
	IdLembaga  string `json:"id_lembaga" form:"id_lembaga" validate:"required"`
	IdStruktur string `json:"id_struktur" form:"id_struktur" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
}

type UnitUpdate struct {
	IdUnit     string `json:"id_unit"`
	IdLembaga  string `json:"id_lembaga" form:"id_lembaga" validate:"required"`
	IdFakultas string `json:"id_fakultas" form:"id_fakultas" validate:"required"`
	IdProdi    string `json:"id_prodi" form:"id_prodi" validate:"required"`
	IdStruktur string `json:"id_struktur" form:"id_struktur" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
}

type UnitProfile struct {
	IdUnit string `json:"id_unit"`
}

type ErrorResponse9 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct9(validasi interface{}) []*ErrorResponse9 {
	var errors []*ErrorResponse9
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse9
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
