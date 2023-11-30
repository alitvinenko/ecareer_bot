package model

type Profile struct {
	ID               int
	WaitingForAnswer bool
	Data             string
	ClubMemberID     int
}

func (p *Profile) Empty() bool {
	return p.Data == ""
}
