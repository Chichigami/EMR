-- name: CreateAppointmentForPatient :one
INSERT INTO appointments (
    patient_id, date_of, reasoning
) 
VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE id = $1;

-- name: GetAppointmentsBasedOnDay :many
SELECT *
FROM appointments
WHERE date_of = $1;

-- name: UpdateAppointment :exec
UPDATE appointments
SET date_of = $2, updated_at = NOW()
WHERE id = $1;

-- name: GetAppointmentBasedOnPatient :many
SELECT *
FROM appointments
WHERE patient_id = $1;

-- name: GetAllPatientsOnDate :many
SELECT 
    p.first_name,
    p.middle_name,
    p.last_name,
    p.date_of_birth,
    p.sex,
    p.gender,
    p.social_security_number,
    p.pharmacy,
    p.email,
    p.location_address,
    p.zip_code,
    p.cell_phone_number,
    p.home_phone_number,
    p.marital_status,
    p.insurance,
    p.primary_care_doctor,
    p.extra_note
FROM appointments a
JOIN patients p ON a.patient_id = p.patient_id
WHERE a.date_of::DATE = $1;