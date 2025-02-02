package handlers_test

import (
	"testing"
	"time"

	"github.com/chichigami/EMR/internal/handlers"
	_ "github.com/lib/pq"
)

func TestConvertStringToDate(t *testing.T) {
	tests := []struct {
		name       string
		stringDate string
		wantDate   time.Time
		wantErr    bool
	}{

		{
			name:       "Valid Date - 1900s",
			stringDate: "19990906",
			wantDate:   time.Date(1999, time.September, 6, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		},
		{
			name:       "Valid Date - 2000s",
			stringDate: "20041109",
			wantDate:   time.Date(2004, time.November, 9, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		},
		{
			name:       "Valid Date - Leap Year",
			stringDate: "20240229",
			wantDate:   time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		},
		{
			name:       "Valid Date - New Year",
			stringDate: "20000101",
			wantDate:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			wantErr:    false,
		},

		{
			name:       "Invalid Date - Incorrect Format",
			stringDate: "2025-11-09",
			wantErr:    true,
		},
		{
			name:       "Invalid Date - Non-numeric",
			stringDate: "abcdefgh",
			wantErr:    true,
		},
		{
			name:       "Invalid Date - Too Short",
			stringDate: "1109",
			wantErr:    true,
		},
		{
			name:       "Invalid Date - Invalid Day",
			stringDate: "32012025",
			wantErr:    true,
		},
		{
			name:       "Invalid Date - Invalid Month",
			stringDate: "13012025",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, err := handlers.ConvertStringToDate(tt.stringDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStringToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !gotDate.Equal(tt.wantDate) {
				t.Errorf("ConvertStringToDate() = %v, want %v", gotDate, tt.wantDate)
			}
		})
	}
}

// func TestMakeNewPatients(t *testing.T) {
// 	json_input := `{
// 		"LastName": "Doe",
// 		"FirstName": "John",
// 		"MiddleName": "A",
// 		"Date_Of_Birth": "20000101",
// 		"Sex": "M",
// 		"Gender": "Male",
// 		"SocialSecurityNumber": "123-45-6789",
// 		"Pharmacy": "PharmaOne",
// 		"Email": "john.doe@example.com",
// 		"LocationAddress": "123 Elm St",
// 		"ZipCode": "12345",
// 		"CellPhoneNumber": "555-1234",
// 		"HomePhoneNumber": "555-5678",
// 		"MaritalStatus": "Single",
// 		"Insurance": "HealthPlus",
// 		"PrimaryCareDoctor": "Dr. Smith",
// 		"ExtraNotes": "No known allergies"
// 	}`
// }
