package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Card struct {
	gorm.Model
	Name             string           `json:"Name"`
	University       string           `json:"University"`
	Direction        string           `json:"Direction"`
	Course           string           `json:"Course"`
	Description      string           `json:"Description"`
	CreatorID        int              `json:"CreatorId"`
	Rating           int              `json:"Rating"`
	ViewersID        int              `json:"ViewersId"`
	Comments         []Comment        `json:"Comments"`
	DateExam         string           `json:"DateExam"`
	IsVisible        bool             `json:"IsVisible"`
	TypeCard         int              `json:"TypeCard"`
	QuestionsAnswers []QuestionAnswer `json:"QuestionsAnswers"`
}

type Comment struct {
	gorm.Model
	IDUser int    `json:"IdUser"`
	Text   string `json:"Text"`
	CardID int
}

type QuestionAnswer struct {
	gorm.Model
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
	CardID   int
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	TokenString string `json:"token"`
}
