package request

import "github.com/go-playground/validator/v10"

type FakultasCreate struct {
	NamaFakultas      string `json:"nama_fakultas" form:"nama_fakultas" validate:"required"`
	SingkatanFakultas string `json:"singkatan_fakultas" form:"singkatan_fakultas" validate:"required"`
}

type FakultasUpdate struct {
	IdFakultas        string `json:"id_fakultas"`
	NamaFakultas      string `json:"nama_fakultas" form:"nama_fakultas" validate:"required"`
	SingkatanFakultas string `json:"singkatan_fakultas" form:"singkatan_fakultas" validate:"required"`
}

type FakultasProfile struct {
	IdFakultas string `json:"id_fakultas"`
}

type ErrorResponse2 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct2(validasi interface{}) []*ErrorResponse2 {
	var errors []*ErrorResponse2
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse2
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
