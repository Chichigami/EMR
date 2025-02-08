-- name: CreateAppointmentForPatient :one
INSERT INTO appointments (
    patient_id, appointment, reasoning
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
WHERE appointment = $1;

-- name: UpdateAppointment :exec
UPDATE appointments
SET appointment = $2, updated_at = NOW()
WHERE id = $1;

-- name: GetAppointmentBasedOnPatient :many
SELECT *
FROM appointments
WHERE patient_id = $1;