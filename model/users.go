package model

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
	OTP  string `json:"otp"`
}
