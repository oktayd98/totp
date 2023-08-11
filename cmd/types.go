package cmd

type OTP struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

type OTPData struct {
	OTPs []OTP `json:"otps"`
}
