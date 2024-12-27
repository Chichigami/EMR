-- +goose Up
CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    appointment_date DATE NOT NULL,
    chart_id SERIAL NOT NULL REFERENCES patients(chart_id),
    reasoning TEXT NULL
);

-- +goose Down
DROP TABLE IF EXISTS appointments;