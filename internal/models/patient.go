package models

type Patient struct {
	LastName             string `json:"last_name"`
	FirstName            string `json:"first_name"`
	MiddleName           string `json:"middle_name,omitempty"`
	DateOfBirth          string `json:"date_of_birth"`
	Sex                  string `json:"sex"`
	Gender               string `json:"gender"`
	SocialSecurityNumber string `json:"ssn,omitempty"`
	Pharmacy             string `json:"pharmacy"`
	Email                string `json:"email,omitempty"`
	LocationAddress      string `json:"location_address"`
	ZipCode              string `json:"zip_code"`
	CellPhoneNumber      string `json:"cell_phone,omitempty"`
	HomePhoneNumber      string `json:"home_phone,omitempty"`
	MaritalStatus        string `json:"marital_status,omitempty"`
	Insurance            string `json:"insurance,omitempty"`
	PrimaryCareDoctor    string `json:"primary_care_doctor,omitempty"`
	ExtraNotes           string `json:"extra_notes,omitempty"`
}
