CREATE TABLE IF NOT EXISTS assets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    value FLOAT NOT NULL,
    acquisition_date DATE NOT NULL
);
