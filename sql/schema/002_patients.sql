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
    social_security_number TEXT DEFAULT NULL,
    pharmacy TEXT DEFAULT NULL,
    email TEXT DEFAULT NULL,
    location_address TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    cell_phone_number TEXT DEFAULT NULL,
    home_phone_number TEXT DEFAULT NULL,
    martital_status TEXT,
    chart_id INT,
    primary_care_doctor TEXT
);
-- +goose Down
DROP TABLE patients;