// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: charts.sql

package database

import (
	"context"

	"github.com/sqlc-dev/pqtype"
)

const createChart = `-- name: CreateChart :one
INSERT INTO charts (
    patient_id
) VALUES (
    $1
)
RETURNING id, created_at, updated_at, patient_id, note, signed_status
`

func (q *Queries) CreateChart(ctx context.Context, patientID int32) (Chart, error) {
	row := q.db.QueryRowContext(ctx, createChart, patientID)
	var i Chart
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PatientID,
		&i.Note,
		&i.SignedStatus,
	)
	return i, err
}

const updateChart = `-- name: UpdateChart :exec
UPDATE charts 
SET note = $2, updated_at = NOW()
WHERE patient_id = $1
`

type UpdateChartParams struct {
	PatientID int32
	Note      pqtype.NullRawMessage
}

func (q *Queries) UpdateChart(ctx context.Context, arg UpdateChartParams) error {
	_, err := q.db.ExecContext(ctx, updateChart, arg.PatientID, arg.Note)
	return err
}
