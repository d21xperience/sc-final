package response

type WsReval struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Status  int     `json:"status"`
	Data    []WsApp `json:"rows"`
}

type WsApp struct {
	SekolahID   string `json:"sekolah_id"`
	AplikasiID  string `json:"aplikasi_id"`
	Nama        string `json:"nama"`
	Token       string `json:"token"`
	Password    string `json:"password"`
	IPAddress   string `json:"ip_address"`
	Port        string `json:"port"`
	MacAddress  string `json:"mac_address"`
	AsalData    string `json:"asal_data"`
	Aktif       string `json:"aktif"`
	ExpiredDate string
	CreateDate  string
	LastUpdate  string
	UpdaterID   string
	LastSync    string
	SoftDelete  string
}
