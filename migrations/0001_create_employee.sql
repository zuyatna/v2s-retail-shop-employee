-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS employees (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(15) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    access_level VARCHAR(20) NOT NULL CHECK (access_level IN ('ordinary', 'supervisor', 'manager', 'hr', 'intern')),
    position VARCHAR(15) NOT NULL,
    work_location VARCHAR(20),
    personal_id VARCHAR(25) UNIQUE,
    address TEXT,
    zip_code VARCHAR(10),
    province VARCHAR(30),
    city VARCHAR(30),
    district VARCHAR(30),
    phone_number VARCHAR(20),
    photo_url TEXT,
    npwp VARCHAR(25) UNIQUE,
    bank_name VARCHAR(30),
    bank_account VARCHAR(30) UNIQUE,
    salary NUMERIC(15, 2) NOT NULL,
    join_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_employees_email ON employees(created_at DESC);
-- +migrate Down
-- SQL in section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS employees;