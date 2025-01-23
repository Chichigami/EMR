-- +goose Up
CREATE TABLE charts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    patient_id INT NOT NULL REFERENCES patients(patient_id),
    note JSONB DEFAULT NULL, 
    signed_status BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE IF EXISTS charts;