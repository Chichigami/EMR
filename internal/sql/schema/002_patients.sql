-- +goose Up
CREATE TABLE patients (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    middle_name TEXT,
    date_of_birth TEXT NOT NULL,
    sex TEXT NOT NULL,
    social_security_number TEXT,
    pharmacy TEXT,
    email TEXT,
    location_address TEXT,
    zip_code TEXT,
    cell_phone_number TEXT,
    home_phone_number TEXT,
    martital_status TEXT,
    primary_care_doctor TEXT
);
-- +goose Down
DROP TABLE patients;