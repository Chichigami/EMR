-- name: CreateChart :one
INSERT INTO charts (
    patient_id
) VALUES (
    $1
)
RETURNING *;


-- name: UpdateChart :exec
UPDATE charts 
SET note = $2, updated_at = NOW()
WHERE patient_id = $1;

-- name: GetPatientCharts :many
SELECT *
FROM charts
WHERE patient_id = $1;