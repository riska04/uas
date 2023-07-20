package entity

type LogActivityEntity struct {
	IdLogActivity int64  `json:"id_log_activity"`
	Auth          string `json:"auth"`
	Url           int64  `json:"url"`
	Method        string `json:"method"`
	Ip            string `json:"ip"`
	Agent         int64  `json:"agent"`
	Tanggal       string `json:"tanggal"`
}
