package request

import "github.com/go-playground/validator/v10"

type LogActivityCreate struct {
	Auth    string `json:"auth" form:"auth" validate:"required"`
	Url     int64  `json:"url" form:"url" validate:"required"`
	Method  string `json:"method" form:"method" validate:"required"`
	Ip      string `json:"ip" form:"ip" validate:"required"`
	Agent   int64  `json:"agent" form:"agent" validate:"required"`
	Tanggal string `json:"tanggal" form:"tanggal" validate:"required"`
}

type LogActivityUpdate struct {
	IdLogActivity int64  `json:"id_log_activity"`
	Auth          string `json:"auth" form:"auth" validate:"required"`
	Url           int64  `json:"url" form:"url" validate:"required"`
	Method        string `json:"method" form:"method" validate:"required"`
	Ip            string `json:"ip" form:"ip" validate:"required"`
	Agent         int64  `json:"agent" form:"agent" validate:"required"`
	Tanggal       string `json:"tanggal" form:"tanggal" validate:"required"`
}

type ErrorResponse11 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct11(validasi interface{}) []*ErrorResponse11 {
	var errors []*ErrorResponse11
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse11
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
