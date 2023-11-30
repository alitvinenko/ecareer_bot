package model

type ClubMember struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Profile  *Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
