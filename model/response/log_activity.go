package response

type LogActivityResponse struct {
	IdLogActivity int64  `json:"id_log_activity"`
	Auth          string `json:"auth" form:"auth" validate:"required"`
	Url           int64  `json:"url" form:"url" validate:"required"`
	Method        string `json:"method" form:"method" validate:"required"`
	Ip            string `json:"ip" form:"ip" validate:"required"`
	Agent         int64  `json:"agent" form:"agent" validate:"required"`
	Tanggal       string `json:"tanggal" form:"tanggal" validate:"required"`
}
