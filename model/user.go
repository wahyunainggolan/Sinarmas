package model

type User struct {
	Id           uint   `json:"id"`
	UserId       string `json:"user_id"`
	Otp          string `json:"otp"`
	StartDateOtp string `json:"start_date_otp"`
}

type RequestUser struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}

type ResponseUser struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}
