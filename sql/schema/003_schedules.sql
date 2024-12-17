-- +goose Up
CREATE TABLE dashboard (
    date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    patient_id NOT NULL REFERENCES patients(id);
);

-- +goose Down
DELETE TABLE dashboard;