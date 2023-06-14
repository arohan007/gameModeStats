package Models

type User struct {
	Name        string `json:"name" validate:"required"`
	EmailID     string `json:"emailId" validate:"required"`
	AreaCode    string `json:"area_code" validate:"required, len=3"`
	CurrentMode string `json:"currentMode"`
}
