package request

import "github.com/go-playground/validator/v10"

type ProgramStudiCreate struct {
	ProgramStudi string `json:"program_studi" form:"program_studi" validate:"required"`
}

type ProgramStudiUpdate struct {
	ProdiId      string `json:"prodi_id"`
	IdFakultas   string `json:"id_fakultas" form:"id_fakultas" validate:"required"`
	ProgramStudi string `json:"program_studi" form:"program_studi" validate:"required"`
	Singkatan    string `json:"singkatan" form:"singkatan" validate:"required"`
}

type ProdisProfile struct {
	IdProdi string `json:"prodi_id"`
}

type ErrorResponse7 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct7(validasi interface{}) []*ErrorResponse7 {
	var errors []*ErrorResponse7
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse7
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
