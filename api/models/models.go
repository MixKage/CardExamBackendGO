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
	Username   string `json:"Username"`
	Email      string `gorm:"unique" json:"Email"`
	Password   string `json:"Password"`
	University string `json:"University"`
	Course     int    `json:"Course"`
	Role       string `json:"Role"`
}

type Authentication struct {
	Email    string `json:"Email,omitempty"`
	Username string `json:"Username,omitempty"`
	Password string `json:"Password"`
}

type Token struct {
	TokenString string `json:"token"`
}
