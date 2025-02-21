-- name: CreateDashboard :one
INSERT INTO dashboards (
    date_of, dashboard_state
)
VALUES (
    $1, $2
)
ON CONFLICT (
    date_of
)
DO UPDATE SET dashboard_state = $2, updated_at = NOW()
RETURNING dashboard_state;

-- name: GetDashboard :one
SELECT dashboard_state
FROM dashboards
WHERE date_of = $1;

-- name: UpdateDashboard :exec
UPDATE dashboards
SET dashboard_state = $2, updated_at = NOW()
WHERE date_of = $1;