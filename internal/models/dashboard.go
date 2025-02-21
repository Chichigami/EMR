package models

type DashboardCache struct {
	Scheduled   []Patient `json:"scheduled"`
	WaitingRoom []Patient `json:"waiting_room"`
	InProgress  []Patient `json:"in_progress"`
	Finished    []Patient `json:"finished"`
}
