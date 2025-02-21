package models

type Appointment struct {
	PatientID   string `json:"patient_id"`
	DateAndTime string `json:"date_and_time"`
	Reason      string `json:"reason"`
}
