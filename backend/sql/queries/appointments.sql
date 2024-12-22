-- name: CreateAppointmentForPatient :one
INSERT INTO appointments (
    appointment_date, chart_id, reasoning
) 
VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE chart_id = $1 AND appointment_date = $2;
