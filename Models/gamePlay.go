package Models

type PlayerStats struct {
	Mode      GameMode `json:"mode"`
	AreaCode  string   `json:"area_code"`
	Timestamp int64    `json:"timestamp"`
}

type GameMode struct {
	Name string `json:"name"`
}

var GameModes []GameMode
