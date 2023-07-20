package request

import (
	"github.com/go-playground/validator/v10"
)

type BerkasCreate struct {
	File          string `json:"file" form:"file" validate:"required"`
	NamaFile      string `json:"nama_file" form:"nama_file" validate:"required"`
	TanggalUpload string `json:"tanggal_upload" form:"tanggal_upload" validate:"required"`
	TanggalAcara  string `json:"tanggal_acara" form:"tanggal_acara" validate:"required"`
	StartJam      string `json:"start_jam" form:"start_jam" validate:"required"`
	SelesaiJam    string `json:"selesai_jam" form:"selesai_jam" validate:"required"`
	IdUnit        string `json:"id_unit" form:"id_unit" validate:"required"`
	// CreateBy      string `json:"create_by" form:"create_by" binding:"required"`
	// CreateDate    string `json:"create_date" form:"create_date" binding:"required"`
	// UpdateBy      string `json:"upate_by"`
	// UpdateDate    string `json:"update_date"`
}

type BerkasUpdate struct {
	IdBerkas      string `json:"id_berkas"`
	File          string `json:"file" form:"file" validate:"required"`
	NamaFile      string `json:"nama_file" form:"nama_file" validate:"required"`
	TanggalUpload string `json:"tanggal_upload" form:"tanggal_upload" validate:"required"`
	TanggalAcara  string `json:"tanggal_acara" form:"tanggal_acara" validate:"required"`
	StartJam      string `json:"start_jam" form:"start_jam" validate:"required"`
	SelesaiJam    string `json:"selesai_jam" form:"selesai_jam" validate:"required"`
	// CreateBy      string `json:"create_by" form:"create_by" binding:"required"`
	// CreateDate    string `json:"create_date" form:"create_date" binding:"required"`
	// UpdateBy      string `json:"upate_by"`
	// UpdateDate    string `json:"update_date"`
}

type BerkasProfile1 struct {
	IdBerkas string `json:"id_berkas"`
}

type ErrorResponse1 struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct1(validasi interface{}) []*ErrorResponse1 {
	var errors []*ErrorResponse1
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse1
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
