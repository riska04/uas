package request

import "github.com/go-playground/validator/v10"

type StrukturCreate struct {
	NamaStruktur string `json:"nama_struktur" form:"nama_struktur" validate:"required"`
}

type StrukturUpdate struct {
	IdStruktur   string `json:"id_struktur"`
	NamaStruktur string `json:"nama_struktur" form:"nama_struktur" validate:"required"`
}

type StrukturProfile struct {
	IdStruktur string `json:"id_struktur"`
}

type ErrorResponse6 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct6(validasi interface{}) []*ErrorResponse6 {
	var errors []*ErrorResponse6
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse6
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
