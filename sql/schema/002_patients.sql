-- +goose Up
CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    middle_name TEXT NULL,
    date_of_birth DATE NOT NULL,
    sex TEXT NOT NULL CHECK (sex in ('Male', 'M', 'Female', 'F', 'Other', 'O')),
    gender TEXT NOT NULL,
    social_security_number TEXT NULL UNIQUE,
    pharmacy TEXT NOT NULL,
    email TEXT NULL UNIQUE,
    location_address TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    cell_phone_number TEXT NULL UNIQUE,
    home_phone_number TEXT NULL UNIQUE,
    marital_status TEXT NULL,
    chart_id SERIAL UNIQUE,
    insurance TEXT DEFAULT 'Self Pay',
    primary_care_doctor TEXT NULL,
    extra_note TEXT NULL
);

-- +goose Down
DROP TABLE IF EXISTS patients;