package model

type User struct {
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Nickname string `json:"nickname" gorm:"not null"`
	Avatar   string `json:"avatar" gorm:"default:static/avatar/default.png"`
	Identity string `json:"identity" gorm:"type:enum('student','teacher');not null"`
}
