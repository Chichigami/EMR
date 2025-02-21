-- +goose Up
CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    date_of TIMESTAMP NOT NULL,
    patient_id INT NOT NULL REFERENCES patients(patient_id) ON DELETE CASCADE,
    reasoning TEXT NULL
);

-- +goose Down
DROP TABLE IF EXISTS appointments;