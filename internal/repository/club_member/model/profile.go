package model

type Profile struct {
	ID               int `gorm:"primaryKey"`
	WaitingForAnswer bool
	Data             string
	ClubMemberID     int
}
