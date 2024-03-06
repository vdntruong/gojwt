package tjwt

import "github.com/golang-jwt/jwt/v5"

type TClaims struct {
	jwt.RegisteredClaims

	EmployeeCode string `json:"MaNhanVien"`
	UserName     string `json:"HoTen"`
	UnitName     string `json:"TenDonVi"`
	Title        string `json:"ChucDanh"`
	PhoneNumber  string `json:"SoDienThoai"`
	Email        string `json:"Email"`
}
