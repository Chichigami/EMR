-- +goose Up
CREATE TABLE dashboards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    date_of TIMESTAMP NOT NULL,
    dashboard_state JSONB DEFAULT NULL
);

-- +goose Down
DROP TABLE IF EXISTS dashboards;