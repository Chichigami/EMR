-- name: CreatePatient :one
INSERT INTO patients (
    last_name, first_name, middle_name, 
    date_of_birth, sex, social_security_number, 
    pharmacy, email, location_address, zip_code, 
    cell_phone_number, home_phone_number, 
    marital_status, primary_care_doctor)
VALUES(
    $1, $2, $3,
    $4, $5, $6,
    $7, $8, $9, $10,
    $11, $12,
    $13, $14
)
RETURNING *;

-- name: DeletePatient :exec
DELETE FROM patients
WHERE chart_id = $1;

-- name: GetPatient :one
SELECT *
FROM patients
WHERE chart_id = $1;
