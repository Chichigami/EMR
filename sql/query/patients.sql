-- CreatePatient :one
INSERT INTO patients (id, created_at, updated_at, last_name, first_name, middle_name, date_of_birth, sex, social_security_number, pharmacy, email, location_address, zip_code, cell_phone_number, home_phone_number, martital_status, chart_id, primary_care_doctor)
VALUES(
    gen_random_uuid(),
    NOW(),
    NOW(),
    last_name = $1
    first_name = $2
    middle_name = $3
    date_of_birth = $4
    sex = $5
    social_security_number = $6
    pharmacy = $7
    email = $8
    location_address = $9
    zip_code = $10
    cell_phone_number = $11
    home_phone_number = $12
    martital_status = $13
    chart_id = $14
    primary_care_doctor = $15
)
RETURNING *;

-- UpdatePatient :exec
UPDATE patients
SET $2 = $3, updated_at = NOW()
WHERE chart_id = $1;


-- DeletePatient :exec
DELETE FROM patients
WHERE chart_id = $1;

-- GetPatient :many
SELECT *
FROM patients
WHERE chart_id = $1 OR date_of_birth = $1 OR last_name = $1 AND first_name = $2;
