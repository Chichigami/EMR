-- name: CreateDashboard :one
INSERT INTO dashboards (
    date_of, dashboard_state
) 
VALUES (
    $1, $2
)
RETURNING *;