-- name: CreatePatient :one
INSERT INTO patients (
    last_name, first_name, middle_name, 
    date_of_birth, sex, gender, social_security_number, 
    pharmacy, email, location_address, zip_code, 
    cell_phone_number, home_phone_number, marital_status, 
    insurance, primary_care_doctor, 
    extra_note)
VALUES (
    $1, $2, $3,
    $4, $5, $6, $7, 
    $8, $9, $10, $11,
    $12, $13, $14,
    $15, $16,
    $17
)
RETURNING patient_id;

-- name: DeletePatient :exec
DELETE FROM patients
WHERE patient_id = $1;

-- name: GetPatient :one
SELECT *
FROM patients
WHERE patient_id = $1;

-- name: DeleteAllPatients :exec
DELETE FROM patients;