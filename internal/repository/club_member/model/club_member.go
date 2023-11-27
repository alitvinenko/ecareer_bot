package model

type ClubMember struct {
	ID       int `gorm:"primaryKey"`
	Username string
}
