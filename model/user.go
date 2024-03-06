package model

type User struct {
	UserName    string `json:"HoTen"`
	PhoneNumber string `json:"SoDienThoai"`
	Email       string `json:"Email"`

	EmployeeCode string `json:"MaNhanVien"`
	UnitName     string `json:"TenDonVi"`
	Title        string `json:"ChucDanh"`
}
