-- name: CreateAppointmentForPatient :one
INSERT INTO appointments (
    appointment_date, patient_id, reasoning
) 
VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE patient_id = $1 AND appointment_date = $2;

-- name: GetAppointmentsBasedOnDay :many
SELECT *
FROM appointments
WHERE appointment_date = $1;

-- name: UpdateAppointment :exec
UPDATE appointments
SET appointment_date = $2, appointment_time = $3, updated_at = NOW()
WHERE id = $1;