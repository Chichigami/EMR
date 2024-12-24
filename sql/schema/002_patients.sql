-- +goose Up
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    middle_name TEXT,
    date_of_birth DATE NOT NULL,
    sex TEXT NOT NULL CHECK (sex in ('Male', 'Female', 'Other')),
    social_security_number TEXT NULL UNIQUE,
    pharmacy TEXT NULL,
    email TEXT NULL UNIQUE,
    location_address TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    cell_phone_number TEXT NULL,
    home_phone_number TEXT NULL,
    marital_status TEXT,
    chart_id SERIAL UNIQUE,
    primary_care_doctor TEXT
);

-- +goose Down
DROP TABLE IF EXISTS patients;